<script lang="ts">
	import type { HTMLAttributes } from "svelte/elements";

	let div = $state<HTMLDivElement | null>(null);

	const pause = (ms: number) => new Promise((fulfil) => setTimeout(fulfil, ms));

	let comments = $state([
		{
			author: "ChatGPT",
			text: "What is your goal ?",
		},
	]);
	let disable = $state(false);

	const typing = { author: "ChatGPT", text: "..." };

	$effect(() => {
		if (div && div.offsetHeight + div.scrollTop > div.scrollHeight - 20) {
			div.scrollTo(0, div.scrollHeight);
		}
	});

	const submitChat: HTMLAttributes<HTMLFormElement>["on:submit"] = async (event) => {
		const formData = new FormData(event.currentTarget);
		const value = formData.get("value");
		if (!value || typeof value !== "string" || disable || !div) {
			return;
		}
		const comment = {
			author: "user",
			text: value,
		};

		let lastQuestion = comments[comments.length - 1].text;
		console.log(lastQuestion);

		// const response = await fetch("UrlOfDatabase", {
		// 	method: "POST",
		// 	headers: {
		// 		"Content-type": "application/json",
		// 	},
		// 	body: JSON.stringify({
		// 		// usergoal : ...,
		// 		question: lastquestion,
		// 		answer: event.target.value,
		// 	}),
		// });

		// await REPLY = GetValue ();
		div.scrollTo(0, div.scrollHeight);
		disable = true;

		// cai nay thay bang fetch data, cai nay nghia lam
		const reply = {
			author: "ChatGPT",
			// text : GetValue
			text: "Me May Beo",
		};

		comments.push(comment);

		event.currentTarget.reset();

		await pause(Math.floor(Math.random() * 500) + 1);
		comments.push(typing);
		await pause(Math.floor(Math.random() * 500) + 1);
		comments.push(reply);
		comments = comments.filter((comment) => comment.text !== "...");
		disable = false;
	};
</script>

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
	<form on:submit|preventDefault={submitChat}>
		<input type="text" name="value" aria-label="Chat input" placeholder="Example: abcxyz,..." />
		<button disabled={disable} type="submit" style="display: none;" />
	</form>
	<p class="posit">Page 1</p>
</div>

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
		background-color: black;
		color: white;
	}

	/* h1 {
    float: left;
    display: block;
    padding: 0.5px;
    text-decoration: none;
    font-size: larger;
    color: #f0f3ed;
  } */

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

	.user span {
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
		background-color: #778899;
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
		word-wrap: break-word; /* IE 5.5-7 */
		white-space: -moz-pre-wrap; /* Firefox 1.0-2.0 */
		white-space: pre-wrap;
		margin-top: 20px;
	}

	.posit {
		position: absolute;
		bottom: 10px;
		left: 50%;
		transform: translateX(-50%);
		color: #ffffff;
	}
</style>
