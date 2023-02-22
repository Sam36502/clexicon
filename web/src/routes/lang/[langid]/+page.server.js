import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import { onMount } from "svelte";

export async function load({ fetch, params }) {

    const url = new URL(PUBLIC_API_URL + "/lang");
    url.searchParams.append("id", params.langid);

    let lang = {};

    await fetch(url.href)
        .then(response => response.json())
        .then(data => {
            lang = data;
        }).catch(err => {
            console.log(err);
            throw error(404);
        });
    
    return {
        lang,
        baseURL: url.origin,
    }
}
