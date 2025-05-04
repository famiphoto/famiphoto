import "@typespec/http";
import "./error.tsp";

using Http;
@tag("auth")
namespace Famiphoto.Auth;

@route("auth")
interface Auth {
  @route("sign_up")
  @post
  @summary("アカウント作成")
  @useAuth(NoAuth)
  postSignUp(@body req: SignUpRequest): {
    @statusCode statusCode: 200;
    @body res: SignUpResponse;
  } | {
    @statusCode statusCode: 409;
    @body error: ErrorResponse;
  };

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

model SignUpRequest {
  /** 取得したいMyIDを指定します */
  myId: string;

  /** 設定したいパスワード */
  password: string;

  /** 管理者権限を付与するか */
  isAdmin?: boolean = false;
}

model SignUpResponse {
  myId: string;
  isAdmin: boolean;
}

model SignInRequest {
  /** 取得したいMyIDを指定します */
  myId: string;

  /** 設定したいパスワード */
  password: string;
}

model SignInResponse {
  myId: string;
  isAdmin: boolean;
}

model MeResponse {
  myId: string;
  isAdmin: boolean;
}
