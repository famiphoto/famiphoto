package sessions

import (
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/utils/random"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

func NewStore(sessionAdapter adapters.SessionAdapter) sessions.Store {
	return &store{
		codecs: []securecookie.Codec{securecookie.New([]byte(config.Env.SessionSecretKey), nil)},
		opts: func() *sessions.Options {
			return &sessions.Options{
				Path:        "/",
				Domain:      "",
				MaxAge:      config.Env.SessionExpireSec,
				Secure:      true,
				HttpOnly:    false,
				Partitioned: false,
				SameSite:    http.SameSiteLaxMode,
			}
		},
		db: sessionAdapter,
		sessionIDGenerator: func() string {
			return random.GenerateRandomString(32)
		},
	}
}

type store struct {
	codecs             []securecookie.Codec
	opts               func() *sessions.Options
	db                 adapters.SessionAdapter
	sessionIDGenerator func() string
}

func (s *store) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(s, name)
}

func (s *store) New(r *http.Request, name string) (*sessions.Session, error) {
	sess := sessions.NewSession(s, name)
	sess.Options = s.opts()
	sess.IsNew = true
	if cookie, errCookie := r.Cookie(name); errCookie == nil {
		err := securecookie.DecodeMulti(name, cookie.Value, &sess.ID, s.codecs...)
		if err == nil {
			if values, err2 := s.db.LoadSession(r.Context(), sess.ID); err2 == nil {
				sess.Values = values
				sess.IsNew = false
			}
		}
	}
	return sess, nil
}

func (s *store) Save(r *http.Request, w http.ResponseWriter, sess *sessions.Session) error {
	if sess.Options.MaxAge <= 0 {
		if err := s.db.DeleteSession(r.Context(), sess.ID, s.getUserID(sess)); err != nil {
			return err
		}
		http.SetCookie(w, sessions.NewCookie(sess.Name(), "", sess.Options))
		return nil
	}

	if sess.ID == "" || sess.IsNew {
		sess.ID = s.sessionIDGenerator()
	}
	if err := s.db.SaveSession(r.Context(), sess.ID, s.getUserID(sess), sess.Values, sess.Options.MaxAge); err != nil {
		return err
	}

	encoded, err := securecookie.EncodeMulti(sess.Name(), sess.ID, s.codecs...)
	if err != nil {
		return err
	}
	http.SetCookie(w, sessions.NewCookie(sess.Name(), encoded, sess.Options))

	return nil
}

func (s *store) getUserID(sess *sessions.Session) int64 {
	return getInt64(sess.Values, userId, 0)
}
