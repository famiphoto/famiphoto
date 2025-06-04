import "@typespec/http";

using Http;

@tag("admin")
namespace Famiphoto.Admin;

@route("/admin/users")
@useAuth(AdminAuth)
interface AdminUserManagement {
    @post
    @summary("ユーザアカウントを作成します")
    createUser(
        @body req: CreateUserRequest;
    ) : {
        @statusCode statusCode: 200;
        @body res: CreatedResponse;
    } | {
        /** ユーザIDが使用済みの場合 */
        @statusCode statusCode: 409;
        @body res: ErrorResponse;
    };

    @route("/{userId}")
    @delete    
    @summary("ユーザアカウントを無効化します")
    deleteUser(
        /** 削除対象のユーザID */
        @path userId: string;
    ) : {
        @statusCode statusCode: 204;
    } | {
        /** ユーザIDが見つからない場合 */
        @statusCode statusCode: 404;
        @body res: ErrorResponse;
    };
}

model CreateUserRequest {
  /** 取得したいUserIDを指定します */
  userId: string;

  /** 設定したいパスワード */
  password: string;

  /** 管理者権限を付与するか */
  isAdmin?: boolean = false;
}

model CreateUserResponse {
  userId: string;
  isAdmin: boolean;
}
