<script>
    export let data;

    let lang = data.lang;

    function updateLang() {

    }
</script>

<h1 class="f1 f-headline-ns bb b--black-10 lh-solid">
    {#if lang.name=''}
        Create
    {:else}
        Update
    {/if}
    Language
</h1>

<a class="f4 link blue hover-light-blue" href="/">Back home</a>

<div class="mv2">
    <input type="text" class="w80 input-reset" bind:value={lang.name} placeholder="Name">
    <!-- TODO: add toast/popup/whatever for success/failure of update -->
    <span class="f6 link dim ph3 pv2 mb2 dib white bg-mid-gray" on:click={updateLang}>Save Changes</span>
</div>
