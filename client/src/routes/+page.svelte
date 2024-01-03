<script lang="ts">
	import type { HTMLAttributes } from "svelte/elements";

	import Input from "$components/Input.svelte";
	import { clsx } from "$lib/clsx";

	let chatUl = $state<HTMLUListElement | null>(null);

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
		chatUl?.scrollTo({
			top: chatUl.scrollHeight,
			behavior: "smooth",
		});
	});

	const submitChat: HTMLAttributes<HTMLFormElement>["on:submit"] = async (event) => {
		const formData = new FormData(event.currentTarget);
		const value = formData.get("value");
		if (!value || typeof value !== "string" || disable || !chatUl) {
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
		chatUl.scrollTo(0, chatUl.scrollHeight);
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

<div class="w-full flex flex-col px-2 py-4 max-h-dvh gap-3">
	<ul class="chat grow overflow-y-auto" bind:this={chatUl}>
		{#each comments as comment}
			<li class={clsx("chat-bubble", comment.author === "user" ? "self" : "others")}>
				<div class="px-4 py-2">{comment.text}</div>
			</li>
		{/each}
	</ul>
	<form class="" on:submit|preventDefault={submitChat}>
		<Input
			type="text"
			name="value"
			label="Type something..."
			id="chat-input"
			autocomplete="off"
		/>
		<button disabled={disable} type="submit" style="display: none;" />
	</form>
</div>
