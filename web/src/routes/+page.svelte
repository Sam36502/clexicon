<script>
    import { PUBLIC_API_URL } from '$env/static/public';
    import { onMount } from "svelte";

    let langs = [];
    let loaded = false;

    const url = new URL(PUBLIC_API_URL + "/lang");

    onMount(async () => {
        fetch(url.href)
            .then(response => response.json())
            .then(data => {
                langs = data;
                loaded = true;
            }).catch(error => {
                console.log(error);
                loaded = true;
                return [];
            });
    });
</script>

<main>
    <h1 class="f1 f-headline-ns bb b--black-10 lh-solid">Clexicon</h1>

    <p class="lh-copy">
        TODO: Think of something appropriate to put on the home page
    </p>

    <h2 class="f2 f1-ns light-red">Languages</h2>

    {#if loaded}
        <ul class="list">
            {#if langs.length > 0}
                {#each langs as lang}
                <li><a href="/lang/{lang.id}" class="f4 link blue hover-light-blue">{lang.name} - {lang.desc}</a></li>
                {/each}
            {:else}
                <li><span class="f5 lh-copy">None yet...</span></li>
                <li><a href="/lang/set" class="f4 link blue hover-light-blue">Why not make one?</a></li>
            {/if}
        </ul>
    {:else}
        <img src="/img/loading.gif" alt="Loading...">
    {/if}
</main>
