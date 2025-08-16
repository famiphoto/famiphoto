import express from "express";
import {createRequestHandler} from "@react-router/express";
export const app = express();

app.get("/server_health", (req, res) => {
    res.json({"health": "ok"})
})

app.use(
    createRequestHandler({
        build: () =>
            import("virtual:react-router/server-build"),
    }),
);
