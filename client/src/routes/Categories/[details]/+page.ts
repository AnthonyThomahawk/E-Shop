import type {PageLoad} from "../../../../.svelte-kit/types/src/routes/Categories/[details]/$types";
export const load = (({ params }) => {
    return {
        post: {
            id : `${params.details}`
        }
    };
}) satisfies PageLoad;