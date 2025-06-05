import "@typespec/http";
import "./error.tsp";

using Http;
@tag("auth")
namespace Famiphoto.Auth;

@route("auth")
interface Auth {
  @route("sign_in")
  @post
  @summary("サインイン")
  @useAuth(NoAuth)
  postSignIn(@body req: SignInRequest): {
    @statusCode statusCode: 200;
    @body res: SignInResponse;
  };

  @route("sign_out")
  @post
  @summary("サインアウトします")
  @useAuth(LoginAuth)
  postSignOut(): {
    @statusCode statusCode: 204;
  };

  @route("me")
  @get
  @summary("ログイン中ユーザの情報を取得します。")
  @doc("ログイン中かどうかの判定にも使います。")
  @useAuth(LoginAuth)
  getMe(): {
    @statusCode statusCode: 200;
    @body res: MeResponse;
  } | {
    /** 未ログインの場合 */
    @statusCode statusCode: 401;
    @body res: ErrorResponse;

  }
}

model SignInRequest {
  /** 取得したいUserIDを指定します */
  userId: string;

  /** 設定したいパスワード */
  password: string;
}

model SignInResponse {
  userId: string;
  isAdmin: boolean;
}

model MeResponse {
  userId: string;
  isAdmin: boolean;
}
