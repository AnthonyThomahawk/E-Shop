// @ts-ignore
import type {PageLoad} from './$types';
export const load = (({ params }) => {
    return {
        post: {
            id : `${params.details}`
        }
    };
}) satisfies PageLoad;