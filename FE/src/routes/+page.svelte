<script>
	let div;

	const pause = (ms) => new Promise((fulfil) => setTimeout(fulfil, ms));
	let comments = [];
	const typing = { author: 'ChatGPT', text: '...' };
	let REPLY;
	let disable = false;

	function GetValue (){
		// async fetch ();
	}


	async function handleKeyDown (event){
		if (event.key === 'Enter' && event.target.value){
			const comment = {
				author : 'user',
				text : event.target.value
			}
			// doi den khi nao fetch xong

			// await REPLY = GetValue ();
			disable = true;

			const reply = {
				author : 'ChatGPT',
				// text : GetValue
				text : 'Me May Beo'
			}

			event.target.value = '';
			comments = [...comments, comment];
			
			await pause(Math.floor(Math.random() * 500) + 1);
			comments = [...comments, typing];
			
			await pause(Math.floor(Math.random() * 500) + 1);
			comments = [...comments, reply].filter(comment => comment !== typing);

			disable = false;
		}
	}
</script>

<nav>
	<h1> Gooool!</h1>
</nav>


<div class="container">
	<div class="chat" bind:this={div}>
		<article class="ChatGPT">
			<span>What is your goal ?</span>
		</article>

		{#each comments as comment}
			<article class={comment.author}>
				<span>{comment.text}</span>
			</article>
		{/each}
	</div>
	<input on:keydown={handleKeyDown} placeholder="Ex : abcxyz, ..." disabled = {disable}/>
</div>

<style>
	nav{
		margin: 0;
		padding: 0;
		overflow: hidden;
		background-color: #333;
	}

	h1 {
		float: left;
		display: block;
		padding: 0.5px;
		text-decoration: none;
		font-size:larger;
		color: #f0f3ed;
	}

	.container {
		display: grid;
		place-items: center;
		width: 1400px;
		height: 500px;
		max-height: 500px;
	}
	
	.chat {
		height: 1em;
		flex: 1 1 auto;
		padding: 0 1em;
		overflow-y: auto;
		scroll-behavior: smooth;
		width: 1000px;
		height: 500px;
		max-height: 500px;
	}

	article {
		margin: 1px 0 1px;
	}

	.user {
		text-align: right;
	}

	.ChatGPT {
		text-align: left;
	}

	span {
		padding: 0.5em 1em;
		display: inline-block;
	}

	.user span {
		background-color: #0074d9;
		color: var(--fg-1);
		border-radius: 1em 1em 0 1em;
		word-break: break-all;
	}
	
	.ChatGPT span {
		background-color: #5F9EA0;
		border-radius: 1em 1em 1em 0;
		color: var(--fg-1);
	}


	input {
		place-items: right;
		border-radius: 0.5em 0.5em 0.5em 0.5em;
		overflow-y: auto;
		scroll-behavior: smooth;
		width: 1000px;
		height: 30px;
		max-width: 500px;
	}
</style>