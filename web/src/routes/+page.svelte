<script>
    import { onMount } from "svelte";

    let langs = [];

    // TODO: Make url ENV var or something
    onMount(async () => {
        fetch("http://192.168.0.4:49157/lang")
            .then(response => response.json())
            .then(data => {
                console.log(data);
                langs = data;
            }).catch(error => {
                console.log(error);
                return [];
            });
    });
</script>

<main>
    <h1 class="f1 f-headline-ns bb b--black-10 lh-solid">Clexicon</h1>

    <p class="lh-copy">
        TODO: Think of something appropriate to put on the home page
    </p>

    <h2 class="f2 f1-ns light-red">List of Languages:</h2>

    <ul class="list">
        {#each langs as lang}
        <a href="/lang/{lang.id}" class="f4 link blue hover-light-blue">{lang.name} - {lang.desc}</a>
        {/each}
    </ul>
</main>
