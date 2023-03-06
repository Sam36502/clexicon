<script>
    export let data;

    let lang = data.lang;

    async function updateLang() {
        const url = new URL(data.baseURL + "/lang");
        const res = await fetch(url.href, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(lang),
        });
        console.log(await res.json());
        console.log(JSON.stringify(lang));
    }
</script>

<h1 class="f1 f-headline-ns bb b--black-10 lh-solid">
    {#if lang.name==''}
        Create
    {:else}
        Update
    {/if}
    Language
</h1>

<a class="f4 link blue hover-light-blue" href="/">Back to home</a>

<div class="mv2">
    <form on:submit={updateLang} onsubmit="return false;">
        <label>Language Name</label>
        <input type="text" class="f4 ba br2 pv2 ph3 input-reset" bind:value={lang.name} on:click={updateLang}>
        <br>
        <label>Language Code</label>
        <input type="text" class="f4 ba br2 pv2 ph3 input-reset" bind:value={lang.code} on:click={updateLang}>
        <br>
        <label>Language Description</label>
        <input type="text" class="f4 ba br2 pv2 ph3 input-reset" bind:value={lang.desc} on:click={updateLang}>
        <br>
        <input type="button" class="f4 ba br2 b--mid-gray pv2 ph4 link dim white bg-mid-gray pointer" value="Save Changes" on:click={updateLang}>
    </form>
</div>
