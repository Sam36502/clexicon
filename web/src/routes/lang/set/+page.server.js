import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import { onMount } from "svelte";

export async function load({ url, fetch }) {
    
    let all = [];
    const apiGetLang = new URL(PUBLIC_API_URL + '/lang');
    await fetch(apiGetLang.href)
        .then(response => response.json())
        .then(data => {
            data.forEach(l => {
                all.push({
                    value: l.code,
                    label: l.name,
                });
            });
        }).catch(err => {
            console.log(err);
        });
    
    let ancestors = [];
    let lang = {
        name: '',
    };
    if (url.searchParams.has('id')) {
        const apiurl = new URL(PUBLIC_API_URL + '/lang');
        const id = url.searchParams.get('id');
        apiurl.searchParams.append('id', id);

        await fetch(apiurl.href)
            .then(response => response.json())
            .then(data => {
                lang = data;
                var i = all.findIndex( x => x.value === lang.code );
                if (i > -1) all.splice(i, 1);
                if (lang.ancestors) {
                    lang.ancestors.forEach(c => {
                        var l = all.find(x => x.value === c);
                        if (l) { ancestors.push(l); }
                    });
                }
            }).catch(err => {
                console.log(err);
                throw error(404);
            });
    }

    return {
        lang,
        all,
        ancestors,
        baseURL: PUBLIC_API_URL,
    };
}
