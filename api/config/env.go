package config

import "github.com/kelseyhightower/envconfig"

type env struct {
	/* Application Settings */
	AppEnv       string `envconfig:"APP_ENV"`
	Port         int64  `envconfig:"API_PORT"`
	ExifTimezone string `envconfig:"EXIF_TIMEZONE"`

	/* Path & URL */
	AssetBaseURL       string `envconfig:"ASSET_BASE_URL"`
	PhotoUploadBaseURL string `envconfig:"PHOTO_UPLOAD_BASE_URL"`
	StorageRootPath    string `envconfig:"STORAGE_ROOT_PATH" default:"/mnt/famiphoto"`
	TempLocalRootPath  string `envconfig:"TEMP_LOCAL_ROOT_PATH" default:"/tmp"`
	AssetRootPath      string `envconfig:"ASSET_ROOT_PATH" default:"/var/www/famiphoto"`

	/* Session */
	SessionSecretKey string `envconfig:"SESSION_SECRET_KEY"`
	SessionExpireSec int    `envconfig:"SESSION_EXPIRE_SEC"`

	/* MySQL */
	MySQLUser     string `envconfig:"MYSQL_USER"`
	MySQLPassword string `envconfig:"MYSQL_PASSWORD"` // SECRET
	MySQLHostName string `envconfig:"MYSQL_HOST_NAME"`
	MySQLPort     string `envconfig:"MYSQL_PORT"`
	MySQLDatabase string `envconfig:"MYSQL_DATABASE"`

	/* SESSION DB */
	SessionDBAddresses []string `envconfig:"SESSION_DB_ADDRESSES"`

	/* Elasticsearch */
	ElasticsearchAddresses   []string `envconfig:"ELASTICSEARCH_ADDRESSES"`
	ElasticsearchUserName    string   `envconfig:"ELASTICSEARCH_USERNAME" default:"elastic"`
	ElasticsearchPassword    string   `envconfig:"ELASTICSEARCH_PASSWORD"`     // SECRET
	ElasticsearchFingerPrint string   `envconfig:"ELASTICSEARCH_FINGER_PRINT"` // SECRET
}

var Env env

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
