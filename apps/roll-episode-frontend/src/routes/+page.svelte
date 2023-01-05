<script>
    import ShowCard from '../components/ShowCard.svelte';

    const API_URL = 'http://localhost:8080/search?q=';
    let name = '';
    let shows = [];

    $: name !== '' && fetch(API_URL + encodeURIComponent(name))
    .then(r => {
    //console.log(r)
    return r.json()
    })
    .then(json => {
    //console.log(json)
    shows = json
    })
</script>


<div class="results">
    <div>
        <input class="search" bind:value={name} placeholder="enter your name">
        {#each shows as show}
            <ShowCard value={show} />
        {/each}
    </div>
</div>

<style>
    .search {
        width: 100%;
        height: 32px;
    }
    .results {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .results {
        max-width: 600px;
        width: 100%;
    }
</style>
