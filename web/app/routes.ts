import {type RouteConfig, index, route} from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),
    route("about", "routes/about.tsx"),

    route("photos/listing", "routes/photos/listing.tsx"),
    route("photos/photo/:photoId", "routes/photos/photo.tsx"),
] satisfies RouteConfig;
