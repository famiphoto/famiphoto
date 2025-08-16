#!/bin/sh

if [ -e ".env" ]; then
    echo ".env file already exists."
    echo "Do you renew again? (y/n)"
    read RENEW
    if [[ $RENEW != "y" && $RENEW != "Y" ]]; then
      exit 0
    fi
fi

# ask

echo "Enter path to your Photo directory(like: /home/famiphoto/). The path is mounted by docker and famiphoto imports."
read CONTENTS_DIR

echo "Enter base URL of Famiphoto frontend. (like: https://famiphoto.myhome.com)"
read BASE_URL

# output to .env file
echo "# Environment parameters for famiphoto" > .env
echo "APP_ENV=prod" >> .env
echo "CONTENTS_DIR=$CONTENTS_DIR" >> .env
echo "" >> .env
echo "## Famiphoto API" >> .env
echo "API_PORT=8080" >> .env
echo "ASSET_BASE_URL=${BASE_URL%/}/api/assets" >> .env
echo "" >> .env
echo "#Famihoto frontend env" >> .env
echo "VITE_IS_DEBUG=false" >> .env
echo "VITE_PORT=3000" >> .env
echo "VITE_API_BASE_URL=http://famiphoto_api:8080" >> .env
echo "VITE_SESSION_SECRET=`head -c 50 </dev/urandom | base64`" >> .env
echo "" >> .env
echo "## MySQL" >> .env
echo "MYSQL_HOST_NAME=famiphoto_mysqldb" >> .env
echo "MYSQL_PORT=3306" >> .env
echo "MYSQL_DATABASE=famiphoto_db" >> .env
echo "MYSQL_USER=famiphoto" >> .env
echo "MYSQL_PASSWORD=password" >> .env
echo "MYSQL_ROOT_PASSWORD=password" >> .env
echo "" >> .env
echo "## Redis for auth" >> .env
echo "SESSION_DB_ADDRESSES=famiphoto_session_db:6379" >> .env
echo "" >> .env
echo "## Elasticsearch" >> .env
echo "ELASTICSEARCH_ADDRESSES=http://famiphoto_es:9200/" >> .env
echo "ELASTICSEARCH_PASSWORD=" >> .env
echo "ELASTICSEARCH_FINGER_PRINT=" >> .env
echo "EXIF_TIMEZONE=Asia/Tokyo" >> .env
echo "" >> .env
echo "## Expire time" >> .env
echo "CC_ACCESS_TOKEN_EXPIRE_IN_SEC=600" >> .env
echo "USER_ACCESS_TOKEN_EXPIRE_IN_SEC=604800" >> .env
echo "LOGIN_TOKEN_EXPIRE_SEC=2592000" >> .env
echo "HMAC_KEY=`head -c 16 </dev/urandom | base64`" >> .env
echo "ACCESS_TOKEN_HASHED_PREFIX=`head -c 16 </dev/urandom | base64`" >> .env
echo "UPLOAD_TOKEN_HASHED_PREFIX=`head -c 16 </dev/urandom | base64`" >> .env
echo "LOGIN_TOKEN_HASHED_PREFIX=`head -c 16 </dev/urandom | base64`" >> .env
echo "## Execute user" >> .env
echo "UID=`id -u`" >> .env
echo "GID=`id -g`" >> .env
echo "USERNAME=`id -un`" >> .env
echo "saved .env"
