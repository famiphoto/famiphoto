import { reactRouter } from "@react-router/dev/vite";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";
import path from 'path'

export default defineConfig(({mode, isSsrBuild}) => ({
  build: {
    rollupOptions: isSsrBuild ? { input: "./server/app.ts" } : undefined
  },
  plugins: [tailwindcss(), reactRouter(), tsconfigPaths()],

  envPrefix: ['VITE_'],

  // 開発実行時のみプロジェクトルートの.env(famiphoto/.env)を環境変数ファイルとして参照する。
  ...(mode === "development" ? { envDir: path.resolve(__dirname, "..")} : {})
}));
