import express from "express";
import {createRequestHandler} from "@react-router/express";
import {createProxyMiddleware} from "http-proxy-middleware";
export const app = express();

app.get("/server_health", (req, res) => {
    res.json({"health": "ok"})
})

app.use("/api", createProxyMiddleware({
    target: import.meta.env.VITE_API_BASE_URL,
    changeOrigin: true,
    pathRewrite: {"^/api": ""}
}))
console.log('Created proxy middleware target: ', import.meta.env.VITE_API_BASE_URL)

app.use(
    createRequestHandler({
        build: () =>
            import("virtual:react-router/server-build"),
    }),
);
