import {type RouteConfig, index, route, layout} from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),
    route("about", "routes/about.tsx"),

    layout("./layouts/photo-layout.tsx", [
        route("photos/listing", "routes/photos/listing.tsx"),
        route("photos/photo/:photoId", "routes/photos/photo.tsx"),
    ])
] satisfies RouteConfig;
