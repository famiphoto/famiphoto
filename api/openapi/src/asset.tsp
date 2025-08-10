import "@typespec/http";
import "./error.tsp";

using Http;
namespace Famiphoto.Assets;

@route("/assets")
@useAuth(LoginAuth) 
interface Assets {
    @route("previews/{photoId}")
    @get
    @summary("プレビュー画像を取得します")
    getPreview(
        @path photoId: string;
    ) : {
        @statusCode statusCode: 200;
        @body data: bytes;
    } | {
        @statusCode statusCode: 404;
        @body data: ErrorResponse;
    };

    @route("thumbnails/{photoId}")
    @get
    @summary("サムネイル画像を取得します")
    getThumbnail(
        @path photoId: string;
    ) : {
        @statusCode statusCode: 200;
        @body data: bytes;
    } | {
        @statusCode statusCode: 404;
        @body data: ErrorResponse;
    };

    @route("original_files/{photoFileId}")
    @get
    @summary("オリジナルの写真ファイルをダウンロードします")
    getOriginalFile(
        @path photoFileId: string;
    ) : {
        @statusCode statusCode: 200;
        @body data: bytes;
    } | {
        @statusCode statusCode: 404;
        @body data: ErrorResponse;
    };
}
