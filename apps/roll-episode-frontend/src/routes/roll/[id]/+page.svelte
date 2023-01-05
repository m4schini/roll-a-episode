<script>
  const API_URL = 'http://141.95.104.175:8080';

  export let episode = {};
  export let episodeL = false;
  export let show = {};
  export let showL = false;

  let loaded = false;
  $: loaded = episodeL && showL

  import { page } from '$app/stores';
  console.log($page.params)

  export let seasons = [];

  fetch(API_URL + "/roll?id=" + $page.params.id)
  .then(r => r.json())
  .then(r => {
    episode = r
    //console.log(episode)
    episodeL = true;
  })

  fetch(API_URL + "/show?id=" + $page.params.id)
  .then(r => r.json())
  .then(r => {
    show = r
    seasons = r.seasons
    //console.log(show)
    showL = true;
  })

</script>

{#if loaded}
<main>
<div class="show">
    <img src={API_URL + "/img?href=" + show.image} alt={"show poster"} />
    <div>
        <h1>{show.name}</h1>
        <button class="roll-btn" on:click={() => location.reload()}>Random Episode</button>
    </div>
</div>
<div class="episode">
    <img src={API_URL + "/img?href=" + episode.image} />
    <div>
        <h1>{episode.name}</h1>
        <h2>Season: {episode.seasonNumber}</h2>
        <h2>Episode: {episode.number}</h2>
    </div>
</div>
</main>

{:else}

<main class="loading">
<div>
loading...
</div>
</main>

{/if}




<style>
    @media screen and (min-width: 1100px) {
        .episode img {
            box-shadow: rgba(240, 46, 170, 0.4) 5px 5px, rgba(240, 46, 170, 0.3) 10px 10px, rgba(240, 46, 170, 0.2) 15px 15px, rgba(240, 46, 170, 0.1) 20px 20px, rgba(240, 46, 170, 0.05) 25px 25px;
        }
        .episode {
            display: flex;
        }

        .episode div {
            margin-top: 32px;
            margin-left: 48px;
        }

        main {
            display: flex;
            justify-content: center;
            align-items: center;
        }
    }

   :global(body) {
		margin: 0;
	}
    .loading {
        display: flex;
        justify-content: center;
        align-items: center;
    }
    main {
            height: 100vh;
            width: 100vw;
            font-family: helvetica, sans-serif;
    }
    .show {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        max-height: 200px;
    }

    .show h1 {
        margin-left: 16px;
        margin-bottom: 0;
        margin-top: 0;
        font-size: 1.5rem;
    }

    .show img {
        height: 200px;
        width: auto;
    }

    .episode img {
        width: 100vw;
        max-width: 800px;
        margin: 0;
        padding: 0;
    }

   .roll-btn {
        padding: 8px;
        margin: 8px;
        background-color: rgba(240, 46, 170, 0.8);
        border-radius: 5px;
        color: white;
        text-decoration: none;
        font-weight: bold;
    }

</style>