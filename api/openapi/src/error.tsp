import "@typespec/http";

namespace Famiphoto;

@error
model ErrorResponse {
  /** HTTPステータスコード */
  statusCode: integer;

  /** APIが定義するエラーコード */
  errorCode: string;

  /** システムエラーメッセージ詳細(本番環境はなし) */
  errorMessage?: string;
}
