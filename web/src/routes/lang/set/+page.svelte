<script>
    import Select from 'svelte-select';
    export let data;

    let lang = data.lang;
    let all = data.all;
    let ancestors = data.ancestors;

    async function updateLang() {
        const url = new URL(data.baseURL + "/lang");
        const res = await fetch(url.href, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(lang),
        });
    }

    function ancestorsChanged() {
        lang.ancestors = [];
        ancestors.forEach(i => {
            lang.ancestors.push(i.value);
        });
        console.log(lang);
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
{#if lang.name != ''}
<br><br><a class="f4 link blue hover-light-blue" href="/lang/{lang.id}">Back to {lang.name} page</a>
{/if}

<div class="mv2">
    <form on:submit={updateLang} onsubmit="return false;">
        <div class="cf pv2">
            <div class="fl w-third tr f4 pa2">
                <label class="f4">Language Name</label>
            </div>
            <div class="fl w-two-thirds">
                <input type="text" class="f4 w-100 ba br2 pv2 ph3 input-reset" bind:value={lang.name} on:click={updateLang}>
            </div>
        </div>
        <div class="cf pv2">
            <div class="fl w-third tr f4 pa2">
                <label class="tr f4 pa2">Language Code</label>
            </div>
            <div class="fl w-two-thirds">
                <input type="text" class="f4 w-100 ba br2 pv2 ph3 input-reset" bind:value={lang.code} on:click={updateLang}>
            </div>
        </div>
        <div class="cf pv2">
            <div class="fl w-third tr f4 pa2">
                <label class="tr f4 pa2">Language Description</label>
            </div>
            <div class="fl w-two-thirds" style="height:150px">
                <textarea
                    class="f4 w-100 h-100 ba br2 pv2 ph3 input-reset"
                    style="resize:none;"
                    bind:value={lang.desc}
                    on:click={updateLang}
                ></textarea>
            </div>
        </div>
        <div class="cf pv2">
            <div class="fl w-third tr f4 pa2">
                <label class="tr f4 pa2">Ancestors</label>
            </div>
            <div class="fl w-two-thirds" style="height:150px">
                <Select items={all} bind:value={ancestors} on:change={ancestorsChanged} multiple></Select>
            </div>
        </div>
        <div class="cf pv2">
            <div class="fl w-third tr f4 pa2"></div>
            <div class="fl w-two-thirds">
                <input 
                    type="button"
                    class="f4 ba br2 b--mid-gray pv2 ph4 link dim white bg-mid-gray pointer"
                    value="Save Changes"
                    on:click={updateLang}
                >
            </div>
        </div>
    </form>
</div>
