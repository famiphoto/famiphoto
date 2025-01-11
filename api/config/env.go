package config

import "github.com/kelseyhightower/envconfig"

type env struct {
	AppEnv                     string   `envconfig:"APP_ENV"`
	Port                       int64    `envconfig:"API_PORT"`
	MySQLUser                  string   `envconfig:"MYSQL_USER"`
	MySQLPassword              string   `envconfig:"MYSQL_PASSWORD"`
	MySQLHostName              string   `envconfig:"MYSQL_HOST_NAME" required:"1"`
	MySQLPort                  string   `envconfig:"MYSQL_PORT"`
	MySQLDatabase              string   `envconfig:"MYSQL_DATABASE"`
	OauthRedisHostName         string   `envconfig:"OAUTH_REDIS_HOST_NAME"`
	OauthRedisDatabase         int64    `envconfig:"OAUTH_REDIS_DATABASE"`
	HMacKey                    string   `envconfig:"HMAC_KEY"`
	CCAccessTokenExpireInSec   int64    `envconfig:"CC_ACCESS_TOKEN_EXPIRE_IN_SEC"`
	UserAccessTokenExpireInSec int64    `envconfig:"USER_ACCESS_TOKEN_EXPIRE_IN_SEC"`
	AccessTokenHashedPrefix    string   `envconfig:"ACCESS_TOKEN_HASHED_PREFIX"`
	LoginTokenHashedPrefix     string   `envconfig:"LOGIN_TOKEN_HASHED_PREFIX"`
	UploadTokenHashedPrefix    string   `envconfig:"UPLOAD_TOKEN_HASHED_PREFIX"`
	LoginTokenExpireSec        int64    `envconfig:"LOGIN_TOKEN_EXPIRE_SEC"`
	ElasticsearchAddresses     []string `envconfig:"ELASTICSEARCH_ADDRESSES"`
	ElasticsearchUserName      string   `envconfig:"ELASTICSEARCH_USERNAME" default:"elastic"`
	ElasticsearchPassword      string   `envconfig:"ELASTICSEARCH_PASSWORD"`
	ElasticsearchFingerPrint   string   `envconfig:"ELASTICSEARCH_FINGER_PRINT"`
	ExifTimezone               string   `envconfig:"EXIF_TIMEZONE"`
	AssetBaseURL               string   `envconfig:"ASSET_BASE_URL"`
	PhotoUploadBaseURL         string   `envconfig:"PHOTO_UPLOAD_BASE_URL"`
	StorageRootPath            string   `envconfig:"STORAGE_ROOT_PATH" default:"/mnt/famiphoto"`
	TempLocalRootPath          string   `envconfig:"TEMP_LOCAL_ROOT_PATH" default:"/tmp"`
	AssetRootPath              string   `envconfig:"ASSET_ROOT_PATH" default:"/var/www/famiphoto"`
}

var Env env

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
