/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    return {
        post: {
            id: `${params.ProductID}`
        }
    };
}