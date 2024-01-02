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
		comments;
		div?.scrollTo(0, div.scrollHeight);
	});

	const submitChat: HTMLAttributes<HTMLFormElement>["on:submit"] = async (event) => {
		const formData = new FormData(event.currentTarget);
		const value = formData.get("value");
		if (!value || typeof value !== "string" || disable || !div) {
			return;
		}
<<<<<<< Updated upstream
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
=======
	}
let isPopupVisible1 = false;

  function togglePopup1() {
    isPopupVisible1 = !isPopupVisible1;
  }
</script>

<body>
    <nav class="navbar">
      <img class="logo" src="PLANNER.jpg" alt="Logo" />
      <div class="button-container">
        <!-- Four buttons inside the container -->
        <button class="action-button1" on:click={togglePopup1}>Who are we?</button>
        <button class="action-button2">Why are we?</button>
        <button class="action-button3">When use this?</button>
        <button class="action-button4">How to use?</button>
      </div>
    </nav>
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

<style>
	.button-container {
		position: fixed;
		bottom: 20px;
		right: 20px;
		display: flex;
		flex-direction: column-reverse; /* Align buttons vertically and reverse order */
=======
</body>
<style>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px; /* Adjust padding for a smaller navbar */
  background-color:  #000000;
  color: #ffffff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
height:10px
}
.vertical-navbar img {
  max-width: 150px;
  max-height: 50px;
}
.button-container {
    display: flex;
    gap: 0px; /* Adjust gap for spacing between buttons */
  }
 .action-button1,.action-button2,.action-button3,.action-button4 {
    padding: 8px 12px; 
    background-color: #000000;
    border: none;
    color: #ffffff;
    cursor: pointer;
    transition: background-color 0.3s;
  }
.action-button1:hover {
  background-color: #0056b3;
}
.action-button2:hover {
  background-color: #0056b3;
}
.action-button3:hover {
  background-color: #0056b3;
}
.action-button4:hover {
  background-color: #0056b3;
}


body{
		align-items: center;
		/* background: linear-gradient(#eb9292, #f8ff78, #6363c9); */
		background: #333333;
		/* display: flex; */
		/* font-family: 'Dosis', 'san-serif'; */
		font-display: swap;
		height: inherit;
		justify-content: center;
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
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
=======
	
	
		.container {
  position: relative;
  display: grid;
  place-items: center;
  width: 100%;
  max-width: 1000px;
  margin: 0 auto;
}

		
>>>>>>> Stashed changes

	.logo {
		position: fixed;
		top: 100px; /* Adjust the top value as needed */
		left:0px; /* Adjust the right value as needed */
		max-width: 1000px; /* Set the maximum width of your logo */
		max-height: 1000px; /* Set the maximum height of your logo */
		width:200px;
                height:200px
	}
	.chat {
<<<<<<< Updated upstream
		height: 1em;
		flex: 1 1 auto;
		padding: 20px 1em 0;
		overflow-y: auto;
		scroll-behavior: smooth;
		width: 1000px;
		height: 500px;
		max-height: 500px;
	}
=======
  height: 500px; /* Increase the height of the chat area */
  width: 100%;
  overflow-y: auto;
  background-color: #f4f4f4;
  border-radius: 8px;
  margin-top: 20px;
  padding: 15px;
}

	
>>>>>>> Stashed changes

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
<<<<<<< Updated upstream
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
=======
word-wrap: break-word;
		max-width: 60%;
word-break: break-all;
   background-color: #0074d9;
   color: #ffffff;
   border-radius: 15px;
   padding: 10px;
   margin: 5px;
   display: inline-block;
   box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

	.ChatGPT span {
   background-color: #778899;
   color: #ffffff;
   border-radius: 15px;
   padding: 10px;
   margin: 5px;
   display: inline-block;
   box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
word-wrap: break-word;
		max-width: 60%;
word-break: break-all;

}

input {
  width: 100%;
  padding: 10px;
  margin-top: 10px;
  border: 1px solid #333333;
  border-radius: 5px;
}
.posit{
 position: absolute;
    bottom: 10px; 
    left: 50%;
    transform: translateX(-50%);
    color: #ffffff;
}
.chat::-webkit-scrollbar {
   width: 8px;
}	
.chat::-webkit-scrollbar-track {
   background-color: #ffffff; /* color of the track */
}
.chat::-webkit-scrollbar-thumb {
   background-color: #000000; /* color of the handle */
   border-radius: 4px; /* rounded corners on the handle */
}
</style>
>>>>>>> Stashed changes
