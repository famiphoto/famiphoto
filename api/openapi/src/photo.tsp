import "@typespec/http";
import "@typespec/openapi";
import "./error.tsp";

using Http;
using OpenAPI;
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
  exifData: ExifData;
  fileTypes: string[];
  files: File[];
}


model ExifData {
  // Camera information
  make: string;
  "model": string;
  serialNumber: string;

  // Date and time information
  dateTimeOriginal: int64;
  createDate: int64;
  subsecTimeOriginal: string;
  timezoneOffset: string;

  // Shooting settings
  exposureTime: float64;
  fNumber: float64;
  iso: int64;
  focalLength: float64;
  focalLengthIn35mm: float64;
  exposureProgram: int64;
  exposureCompensation: float64;
  meteringMode: int64;
  flash: int64;

  // Lens information
  lensMake: string;
  lensModel: string;
  lensSerialNumber: string;

  // Image information
  width: int64;
  height: int64;
  colorSpace: int64;
  whiteBalance: int64;
  orientation: PhotoOrientation;

  // Software information
  software: string;
  firmware: string;
}

model File {
  fileId: string;
  photoId: string;
  fileType: string;
  fileHash: string;
  fileName: string;
}

@extension("x-enum-varnames", #[
    "ExifOrientationNone",
    "ExifOrientationHorizontal",
    "ExifOrientationRotate180",
    "ExifOrientationVertical",
    "ExifOrientationHorizontalRotate270",
    "ExifOrientationRotate90",
    "ExifOrientationHorizontalRotate90",
    "ExifOrientationRotate270",
 ])
@extension("x-go-type", "int64")
enum PhotoOrientation {
     // 不要
    ExifOrientationNone: 1;
    // 水平方向に反転
    ExifOrientationHorizontal          : 2;
    // 時計回りに180度回転
    ExifOrientationRotate180           : 3;
    // 垂直方向に反転
    ExifOrientationVertical            : 4;
    // 水平方向に反転 + 時計回りに270度回転
    ExifOrientationHorizontalRotate270 : 5;
    // 時計回りに90度回転
    ExifOrientationRotate90            : 6;
    // 水平方向に反転 + 時計回りに90度回転
    ExifOrientationHorizontalRotate90  : 7;
    // 時計回りに270度回転
    ExifOrientationRotate270           : 8;
 }