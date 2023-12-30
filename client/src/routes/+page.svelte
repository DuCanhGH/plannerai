<script>

	import { beforeUpdate, afterUpdate } from 'svelte';
	let div;
	let autoScroll = false;

	const pause = (ms) => new Promise((fulfil) => setTimeout(fulfil, ms));
	let comments = [
		{
			author : 'ChatGPT',
			text : 'What is your goal ?'
		}
	];
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

			let lastquestion = comments[comments.length - 1].text;
			console.log (lastquestion);

			const response = await fetch("UrlOfDatabase",{
				method: "POST",
				headers: {
					"Content-type": "application/json",
				},
				body: JSON.stringify({
					// usergoal : ...,
					question : lastquestion,
					answer : event.target.value
				})
			});


			// await REPLY = GetValue ();
			div.scrollTo(0, div.scrollHeight);
			disable = true;


			
			// cai nay thay bang fetch data, cai nay nghia lam
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
  <div class="button-container">
    <!-- Four buttons inside the container -->
    <button class="action-button">Who are us?</button>
    <button class="action-button">Why are us?</button>
    <button class="action-button">When use this</button>
    <button class="action-button">How to use</button>
  </div>
<img class="logo" src="PLANNER.jpg" alt="Logo" />
<div class="container">
	<div class="chat" bind:this={div}>

		{#each comments as comment}
			<article class={comment.author}>
				<span>{comment.text}</span>
			</article>
		{/each}
	</div>
	<input on:keydown={handleKeyDown} placeholder="Ex : abcxyz, ..."/>
</div>
<p class="posit">Page 1</p>
</body>

<style>
 .button-container {
    position: fixed;
    bottom: 20px;
    right: 20px;
    display: flex;
    flex-direction: column-reverse; /* Align buttons vertically and reverse order */
  }
.action-button {
  width: 200px;
    height: 60px;
   background-color: white;
  border: 2px solid #fff; /* White border */
  color: black;
    margin-bottom: 10px; /* Adjust spacing between buttons */
    cursor: pointer;
    transition: background-color 0.9s;
border: 2px solid black;
  }
.action-button:hover {
  background-color:  black;
color:white;
}	
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
		position: relative;
		display: grid;
		place-items: center;
		width: 1000px;
		height: 530px;
		max-height: 5000px;
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

		/* position: absolute; top: 36.1511px; left: 36.1511px; width: 570.698px; height: 289.209px; */
	}

	.logo {
		position: fixed;
		top: 100px; /* Adjust the top value as needed */
		right: 100px; /* Adjust the right value as needed */
		max-width: 200px; /* Set the maximum width of your logo */
		max-height: 200px; /* Set the maximum height of your logo */
		z-index: 999; /* Set a high z-index to ensure it appears above other elements */
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
		max-width: 60%; 
		background-color: #0074d9; 
		color: var(--fg-1);
		border-radius: 1em 1em 0 1em;
		word-break: break-all;
	}

	.ChatGPT span {
		word-wrap: break-word;
		max-width: 60%; 
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
    width: 80%;
    max-width: 80%;
    height: auto;
	word-wrap: break-word;      /* IE 5.5-7 */
      white-space: -moz-pre-wrap; /* Firefox 1.0-2.0 */
      white-space: pre-wrap;   
    margin-top: 20px;
	}

.posit{
 position: absolute;
    bottom: 10px; 
    left: 50%;
    transform: translateX(-50%);
    color: #ffffff;
}
</style>
