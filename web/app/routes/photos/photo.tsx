import type { Route } from "./+types/photo";


export default function Photo({ params }: Route.ComponentProps) {
    return <>Photo: {params.photoId}</>
}