<script>
	import ShowCard from '../components/ShowCard.svelte';
	import { env } from '$env/dynamic/public';


	let name = '';
	let shows = [];

	const search = () => {
		fetch(env.PUBLIC_API_BASE_URL + '/search?q=' + encodeURIComponent(name))
			.then((r) => {
				//console.log(r)
				return r.json();
			})
			.then((json) => {
				//console.log(json)
				shows = json;
			})
			.catch((e) => {
				console.log(e);
			});
	};

	function handleKeyDown(event) {
		if (event.key === 'Enter') {
			search();
		}
	}
</script>

<div class="results">
	<div>
		<input class="search" bind:value={name} placeholder="enter your name" on:keydown={handleKeyDown}/>
		{#each shows as show}
			<ShowCard value={show} />
		{/each}
		<input type="submit" value="Search" on:click={search} />
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
