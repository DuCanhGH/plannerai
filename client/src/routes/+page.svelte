<script>

	import { beforeUpdate, afterUpdate } from 'svelte';
	let div;
	let autoScroll = false;

	const pause = (ms) => new Promise((fulfil) => setTimeout(fulfil, ms));
	let comments = [];
	const typing = { author: 'ChatGPT', text: '...' };
	let REPLY;
	let disable = false;

	function GetValue (){
		// async fetch ();
	}

	beforeUpdate(() => {
		autoScroll = div && (div.offsetHeight + div.scrollTop) > (div.scrollHeight - 20);
	});

	afterUpdate(() => {
		if (autoScroll) div.scrollTo(0, div.scrollHeight);
	});


	async function handleKeyDown (event){
		if (event.key === 'Enter' && event.target.value && !disable){
			const comment = {
				author : 'user',
				text : event.target.value
			}
			// doi den khi nao fetch xong

			// await REPLY = GetValue ();
			div.scrollTo(0, div.scrollHeight);
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

<body>
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
	<input on:keydown={handleKeyDown} placeholder="Ex : abcxyz, ..."/>
</div>
<p class="posit">Page 3</p>
</body>

<style>
	body{
		align-items: center;
		/* background: linear-gradient(#eb9292, #f8ff78, #6363c9); */
		background: #1E90FF;
		/* display: flex; */
		/* font-family: 'Dosis', 'san-serif'; */
		font-display: swap;
		height: inherit;
		justify-content: center;
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
		width: 1000px;
		height: 500px;
		max-height: 500px;
		justify-content: center;

		backdrop-filter: blur(5.5px);
		-webkit-backdrop-filter: blur(5.5px);
		border: 1px solid rgba(255, 255, 255, 0.1);
		border-radius: 16px;
		box-shadow: 0 4px 30px rgba(35, 35, 35, 0.1);
		color: #ffffff;
		backdrop-filter: blur(4px);
		-webkit-backdrop-filter: blur(4px);
		background: #333333;
		border: 1px solid rgba(255, 255, 255, 0.34);
		flex-basis: 400px;		
	}
	
	.chat {
    height: 1em;
    flex: 1 1 auto;
    padding: 20px 1em 0; 
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

	span {
		padding: 0.5em 1em;
		display: inline-block;
	}

.user span{
word-wrap: break-word;
    max-width: 80%; 
    background-color: #0074d9; 
    color: var(--fg-1);
    border-radius: 1em 1em 0 1em;
    word-break: break-all;
}
.ChatGPT span {
    word-wrap: break-word;
    max-width: 80%; 
    background-color:   #778899;
    color: var(--fg-1);
    border-radius: 1em 1em 0 1em;
    word-break: break-all;
}


	input {
    place-items: right;
    border-radius: 0.5em 0.5em 0.5em 0.5em;
    overflow-y: auto;
    scroll-behavior: smooth;
    width: 80%; /* Adjust the percentage as needed */
    height: 30px;
    max-width: 500px;
}
.posit{
 position: absolute;
    bottom: 10px; 
    left: 50%;
    transform: translateX(-50%);
    color: #ffffff;
}
</style>
