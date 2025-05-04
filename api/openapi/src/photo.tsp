import "@typespec/http";
import "./error.tsp";

using Http;
namespace Famiphoto.Photos;

@route("/photos")
@useAuth(LoginAuth)
interface Photos {
  @get
  @summary("写真一覧を取得します。")
  getPhotoList(@query limit?: int64 = 30, @query offset?: int64 = 0): {
    @statusCode statusCode: 200;
    @body res: PhotoListResponse;
  };

  @route("/{photoId}")
  @get
  @summary("写真情報を取得します。")
  getPhoto(@path photoId: string): {
    @statusCode statusCode: 200;
    @body res: Photo;
  } | {
    @statusCode statusCode: 404;
    @body error: ErrorResponse;
  };
}

model PhotoItem {
  photoId: string;
  name: string;
  importedAt: string;
  dateTimeOriginal: offsetDateTime;
  previewUrl: url;
  thumbnailUrl: url;
}

model PhotoListResponse {
  items: PhotoItem[];
  total: int64;
  offset: int64;
}

model Photo {
  photoId: string;
  name: string;
  importedAt: string;
  dateTimeOriginal: offsetDateTime;
  previewUrl: url;
  thumbnailUrl: url;
  exifData: Exif[];
  fileTypes: string[];
  files: File[];
}

model Exif {
  id: string;
  tagId: string;
  tagType: string;
  valueString: string;
}

model File {
  fileId: string;
  photoId: string;
  fileType: string;
  fileHash: string;
  fileName: string;
}
