export function loader() {
    return new Response("OK", {
        status: 200,
        headers: {
            "Content-type": "text/plain"
        }
    })
}