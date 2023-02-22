<script> import Word from './Word.svelte';

    export let data;

    let query = '';
    let results = [];

    function search() {
        const url = new URL(data.baseURL + "/lang/" + data.lang.id + "/search");
        url.searchParams.append("q", query);
        fetch(url.href)
            .then(response => response.json())
            .then(data => {
                results = data;
            }).catch(error => {
                console.log(error);
                return [];
            });
    }
</script>

<h1 class="f1 f-headline-ns bb b--black-10 lh-solid">{data.lang.name}</h1>

<a class="f4 link blue hover-light-blue" href="/">Back home</a>

<p class="lh-copy">
    {data.lang.desc}
</p>

<h2 class="f2 f1-ns light-red">Dictionary</h2>

<div class="mv2">
    <input type="text" class="w80 input-reset" bind:value={query}>
    <span class="f6 link dim ph3 pv2 mb2 dib white bg-mid-gray" on:click={search}>Search!</span>
</div>

<div class="mv2 bb bt">
    {#if results.length > 0}
        {#each results as word}
            <Word
                word={word}
                baseURL={data.baseURL}
                langID={data.lang.id}
            />
        {/each}
    {:else}
        <div class="w100 tc pa4">
            No Results
        </div>
    {/if}
</div>

<div class="pa4">
    TODO: Footer...
</div>
