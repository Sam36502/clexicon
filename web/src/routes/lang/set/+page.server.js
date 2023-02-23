import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import { onMount } from "svelte";

export async function load({ url, fetch }) {
    
    let lang = {
        name: '',
    };
    if (url.searchParams.has('id')) {
        const id = url.searchParams.get('id');
        const apiurl = new URL(PUBLIC_API_URL + '/lang');
        apiurl.searchParams.append('id', id);

        await fetch(apiurl.href)
            .then(response => response.json())
            .then(data => {
                lang = data;
            }).catch(err => {
                console.log(err);
                throw error(404);
            });
    }
    
    return {
        lang,
        baseURL: PUBLIC_API_URL,
    };
}
