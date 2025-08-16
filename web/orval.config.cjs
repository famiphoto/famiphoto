/**
 * OpenAPIのスキーマ定義からAPIクライアントを自動生成するためのorvalの設定ファイル
 * */

module.exports = {
    "famiphoto-api": {
        input: "../api/openapi/openapi.yaml",
        output: {
            target: "shared/api/schema.ts",
            baseUrl: "/api"
        },
    }
}