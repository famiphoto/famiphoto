package models

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/famiphoto/famiphoto/api/utils/cast"
)

/**
* elasticsearch 写真のインデックス
 */

// DateTimeOriginalParts represents the parts of the date_time_original field
type DateTimeOriginalParts struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

// OriginalUrl represents an original URL with its metadata
type OriginalUrl struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
	MD5Hash  string `json:"md5_hash"`
}

// ImageUrls represents the URLs for different versions of the image
type ImageUrls struct {
	ThumbnailURL string        `json:"thumbnail_url"`
	PreviewURL   string        `json:"preview_url"`
	OriginalURLs []OriginalUrl `json:"original_urls"`
}

// OriginalImageFile represents an original image file with its metadata
type OriginalImageFile struct {
	Path     string `json:"path"`
	MimeType string `json:"mime_type"`
	MD5Hash  string `json:"md5_hash"`
}

// ExifData represents the EXIF data of the image
type ExifData struct {
	// Camera information
	Make         string `json:"make"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`

	// Date and time information
	DateTimeOriginal   string `json:"date_time_original"`
	DateTimeDigitized  string `json:"date_time_digitized"`
	CreateDate         string `json:"create_date"`
	SubsecTimeOriginal string `json:"subsec_time_original"`
	TimezoneOffset     string `json:"timezone_offset"`

	// Shooting settings
	ExposureTime         string  `json:"exposure_time"`
	FNumber              float64 `json:"f_number"`
	ISO                  int     `json:"iso"`
	FocalLength          float64 `json:"focal_length"`
	FocalLengthIn35mm    float64 `json:"focal_length_in_35mm"`
	ExposureProgram      string  `json:"exposure_program"`
	ExposureCompensation float64 `json:"exposure_compensation"`
	MeteringMode         string  `json:"metering_mode"`
	Flash                string  `json:"flash"`

	// Lens information
	LensMake         string `json:"lens_make"`
	LensModel        string `json:"lens_model"`
	LensSerialNumber string `json:"lens_serial_number"`

	// Image information
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	ColorSpace   string `json:"color_space"`
	WhiteBalance string `json:"white_balance"`
	Orientation  int    `json:"orientation"`

	// GPS information
	GPSLatitude  float64 `json:"gps_latitude"`
	GPSLongitude float64 `json:"gps_longitude"`
	GPSAltitude  float64 `json:"gps_altitude"`

	// Software information
	Software string `json:"software"`
	Firmware string `json:"firmware"`
}

type PhotoIndex struct {
	PhotoID               string                `json:"photo_id"`
	Name                  string                `json:"name"`
	ImportedAt            int64                 `json:"imported_at"`
	DateTimeOriginal      int64                 `json:"date_time_original"`
	DateTimeOriginalParts DateTimeOriginalParts `json:"date_time_original_parts"`
	Orientation           int                   `json:"orientation"`
	Location              interface{}           `json:"location"` // GeoPointProperty
	ImageUrls             ImageUrls             `json:"image_urls"`
	OriginalImageFiles    []OriginalImageFile   `json:"original_image_files"`
	Exif                  ExifData              `json:"exif"`
	DescriptionJa         string                `json:"description_ja"`
	DescriptionEn         string                `json:"description_en"`
}

func (m PhotoIndex) IndexName() string {
	return "photo"
}

func (m PhotoIndex) DocumentID() string {
	return m.PhotoID
}

// PhotoElasticSearchMapping elasticsearchの写真のマッピングを定義します。
func PhotoElasticSearchMapping() *types.TypeMapping {
	return &types.TypeMapping{
		Dynamic: &dynamicmapping.False,
		Properties: map[string]types.Property{
			// 写真ID
			"photo_id": types.KeywordProperty{},

			// 写真の名前(デフォルトはファイル名の拡張子なし）
			"name": types.TextProperty{},

			// 日本語での説明
			"description_ja": types.TextProperty{
				Analyzer:       cast.Ptr("kuromoji"),
				SearchAnalyzer: cast.Ptr("kuromoji_search"),
				Fields: map[string]types.Property{
					"keyword": types.KeywordProperty{},
				},
			},

			// 英語での説明
			"description_en": types.TextProperty{
				Analyzer: cast.Ptr("standard"),
			},

			// 取り込み日時
			"imported_at": types.DateProperty{
				Format: cast.Ptr("epoch_second"),
			},

			// 撮影日時(Exifと同じ、但しこちらはUNIX TIME)
			"date_time_original": types.DateProperty{
				Format: cast.Ptr("epoch_second"),
			},
			"date_time_original_parts": types.ObjectProperty{
				Properties: map[string]types.Property{
					"year":   types.IntegerNumberProperty{},
					"month":  types.IntegerNumberProperty{},
					"day":    types.IntegerNumberProperty{},
					"hour":   types.IntegerNumberProperty{},
					"minute": types.IntegerNumberProperty{},
				},
			},

			// 画像の向き(Exifと同じ)
			"orientation": types.IntegerNumberProperty{},

			// 撮影場所
			"location": types.GeoPointProperty{},

			"image_urls": types.ObjectProperty{
				Enabled: cast.Ptr(false),
				Properties: map[string]types.Property{
					"thumbnail_url": types.TextProperty{
						Index: cast.Ptr(false),
					},
					"preview_url": types.TextProperty{
						Index: cast.Ptr(false),
					},
					"original_urls": types.NestedProperty{
						Properties: map[string]types.Property{
							"url":       types.TextProperty{Index: cast.Ptr(false)},
							"mime_type": types.KeywordProperty{},
							"md5_hash":  types.KeywordProperty{},
						},
					},
				},
			},

			// 元ファイルへのパス
			"original_image_files": types.NestedProperty{
				Properties: map[string]types.Property{
					"path":      types.TextProperty{Index: cast.Ptr(false)},
					"mime_type": types.KeywordProperty{},
					"md5_hash":  types.KeywordProperty{},
				},
			},

			// EXIF情報
			"exif": types.ObjectProperty{
				Properties: map[string]types.Property{
					/** カメラ情報 */
					// メーカー名
					"make": types.KeywordProperty{},
					// モデル名
					"model": types.KeywordProperty{},
					// シリアル番号
					"serial_number": types.KeywordProperty{},

					/** 日時関連情報 */
					// 撮影日時
					"date_time_original": types.DateProperty{
						Format: cast.Ptr("strict_date_time"),
					},
					// デジタル化日時
					"date_time_digitized": types.DateProperty{
						Format: cast.Ptr("strict_date_time"),
					},
					// 作成日時
					"create_date": types.DateProperty{
						Format: cast.Ptr("strict_date_time"),
					},
					// ミリ秒以下の精度
					"subsec_time_original": types.KeywordProperty{},
					// タイムゾーンオフセット
					"timezone_offset": types.KeywordProperty{},

					/** 撮影設定 */
					// 露出時間
					"exposure_time": types.KeywordProperty{},
					// F値
					"f_number": types.FloatNumberProperty{},
					// ISO感度
					"iso": types.IntegerNumberProperty{},
					// 焦点距離
					"focal_length": types.FloatNumberProperty{},
					// 35mm換算焦点距離
					"focal_length_in_35mm": types.FloatNumberProperty{},
					// 露出プログラム
					"exposure_program": types.KeywordProperty{},
					// 露出補正値
					"exposure_compensation": types.FloatNumberProperty{},
					// 測光モード
					"metering_mode": types.KeywordProperty{},
					// フラッシュ設定
					"flash": types.KeywordProperty{},

					/** レンズ情報 */
					// レンズメーカー
					"lens_make": types.KeywordProperty{},
					// レンズモデル
					"lens_model": types.KeywordProperty{},
					// レンズシリアル番号
					"lens_serial_number": types.KeywordProperty{},

					/** 画像情報 */
					// 画像幅
					"width": types.IntegerNumberProperty{},
					// 画像高さ
					"height": types.IntegerNumberProperty{},
					// 色空間
					"color_space": types.KeywordProperty{},
					// ホワイトバランス
					"white_balance": types.KeywordProperty{},
					// 画像の向き
					"orientation": types.IntegerNumberProperty{},

					/** 位置情報 */
					// GPS緯度
					"gps_latitude": types.FloatNumberProperty{},
					// GPS経度
					"gps_longitude": types.FloatNumberProperty{},
					// GPS高度
					"gps_altitude": types.FloatNumberProperty{},

					/** ソフトウェア情報 */
					// 使用ソフトウェア
					"software": types.KeywordProperty{},
					// ファームウェアバージョン
					"firmware": types.KeywordProperty{},
				},
			},
		},
	}
}
