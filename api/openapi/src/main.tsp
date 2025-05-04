import "@typespec/http";
import "./auth.tsp";
import "./photo.tsp";

using Http;
@service(#{ title: "Famiphoto API" })
@server("http://localhost:8080", "local")
namespace Famiphoto;

/**
 * ユーザのログイン認証
 */
alias LoginAuth = ApiKeyAuth<ApiKeyLocation.cookie, "session">;

@route("health")
interface Health {
  /** APIサーバーの状態を確認する */
  @get
  @summary("ヘルスチェック")
  @useAuth(NoAuth)
  getHealth(): {
    @statusCode statusCode: 200;
    @body res: HealthResponse;
  };
}

model HealthResponse {
  status: string = "ok";
}
