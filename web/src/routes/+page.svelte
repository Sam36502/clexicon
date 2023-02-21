<script>
    import { onMount } from "svelte";

    let langs = [];

    // TODO: Make url ENV var or something
    onMount(async () => {
        fetch("http://192.168.0.4:49156/lang")
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

    <table class="collapse ma3 ba br2 b--black-10 pv2 ph3">
        <thead>
            <tr class="striped--light-gray">
                <th class="tl f6 ttu fw6 pv2 ph3">Name</th>
                <th class="tl f6 ttu fw6 pv2 ph3">Description</th>
            </tr>
        </thead>
        <tbody>

            {#each langs as lang}
            <a class="link" href="/dict/{lang.id}">
                <tr class="striped--light-gray">
                    <td class="pv2 ph3">{lang.name}</td>
                    <td class="pv2 ph3">{lang.desc}</td>
                </tr>
            </a>
            {/each}

        </tbody>
    </table>
</main>
