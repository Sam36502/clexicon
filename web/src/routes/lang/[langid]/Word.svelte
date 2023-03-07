<script>
    export let word;
    export let baseURL;
    export let langID;

    let meanstring = word.meanings.join(", ");

    // Get root
    let rootword = {};
    if (word.root != 0) {
        const url = new URL(baseURL + "/lang/" + langID + "/word");
        url.searchParams.append("id", word.root);
        fetch(url.href)
            .then(response => response.json())
            .then(data => {
                rootword = data;
            }).catch(error => {
                console.log(error);
                return [];
            });
    }

</script>

<a href={'/lang/' + langID + '/word/' + word.id} class="cf pv2 link db black hoverbg">

    <div class="fl w-20 ph2">
        {word.orthography}
        &nbsp;
    </div>

    <div class="fl w-20 pa2 br b--silver tr">
        <strong class="f3">{word.romanisation}</strong><br>
        {#each word.tags as tag}
        <i class="link underline-hover" title="{tag.name} - {tag.desc}">{tag.tag}.</i>
        {/each}
        <p class="f5 mv1">/{word.ipa}/</p>
    </div>

    <div class="fl w-40 ph2 f5 lh-copy">
        {meanstring}
    </div>

    <div class="fl w-20 ph2">
        {#if word.etymology != ""}
            <strong class="f5">Etymology</strong><br>
            <p class="f6 mv0 lh-copy">{word.etymology}</p><br>
        {/if}
        {#if word.root != 0}
            <strong class="f5">Derived from</strong><br>
            <a href={'/lang/' + langID + '/word/' + rootword.id} class="f5 link blue hover-light-blue">{rootword.romanisation}</a>
        {/if}
    </div>

</a>

<style>
    .hoverbg:hover {
        background-color: #EEEEEE;
    }
</style>
