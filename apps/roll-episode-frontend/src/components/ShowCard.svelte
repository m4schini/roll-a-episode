<script>
    export let value = {};
</script>

<article class="card">
    <div class="side">
         <div class="data">
        <h2>{value.name} <span>{value.year}</span></h2>
        <p>
        {value.overview}
        </p>
    </div>
        <a class="roll-btn" href={"/roll/" + value.tvdb_id}>Random Episode</a>
    </div>

    {#if value.image_url}
        <img src={value?.image_url} alt="{value.name} dancing">
    {/if}
</article>

<style>
	.card {
        background-color: white;
        border-radius: 8px;
        margin: 16px;
        margin-bottom: 32px;
        font-family: helvetica, sans-serif;

        box-shadow: rgba(240, 46, 170, 0.4) 0px 5px, rgba(240, 46, 170, 0.3) 0px 10px, rgba(240, 46, 170, 0.2) 0px 15px, rgba(240, 46, 170, 0.1) 0px 20px, rgba(240, 46, 170, 0.05) 0px 25px;
        text-decoration: none;
        color: black;

        max-height: 300px;
        max-width: 600px;

        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }

    .card:hover {
        box-shadow: none;
        background-color: #f4f4f4;
    }

    .card h2 {
        font-weight: bold;
    }

    .card h2 span {
        font-size: 1rem;
        font-weight: normal;
    }

    .card p {
        overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    line-height: 16px;
    max-height: calc(16 * 3)px;
    -webkit-line-clamp: 3; /* Write the number of
                              lines you want to be
                              displayed */
    -webkit-box-orient: vertical;
    }

    .card img {
        max-height: 300px;
        border-radius: 0px 8px 8px 0px;
        height: 100%;
        width: auto;
    }

    .card .side {
        margin: 16px;
        color: 'white';
        display: flex;
        justify-content: space-between;
        flex-direction: column;
    }

    .card .roll-btn {
        padding: 8px;
        margin: 8px;
        background-color: rgba(240, 46, 170, 0.8);
        border-radius: 5px;
        color: white;
        text-decoration: none;
        font-weight: bold;
    }
</style>
