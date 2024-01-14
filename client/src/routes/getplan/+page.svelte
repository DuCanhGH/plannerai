<script lang="ts">
	import { onMount } from "svelte";

	import CaptureScreen from "$components/CaptureScreen.svelte";
	import Image from "$components/Image.svelte";
	import mmb from "$lib/return.png"
	import { onMount } from "svelte";
	import CaptureScreen from "$components/CaptureScreen.svelte";
	import Image from "$components/Image.svelte";
	import mmb from "$lib/return.png";
	import { PUBLIC_SERVER_URL } from "$env/static/public";
	import Loading from "$components/icons/Loading.svelte"

	interface PlanData {
		time: string;
		plan: string;
	}

	let resultPlan: PlanData[] | undefined = $state();

	onMount(async () => {
		let infomation =
			'{"usergoal": "I want to get IOI gold medal","time": "2 years","Queries": [{"planner": "1. Have you participated in the IOI competition before? If yes, what was your previous performance?","user": "I\'ve got bronze"},{"planner": "2. How much time are you willing to dedicate to preparation for the IOI competition?","user": "12h per day!"}]}';
		if (infomation == "") {
			window.location.href = "/";
			return;
		}
		const resp = await fetch(`${PUBLIC_SERVER_URL}/getplan`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: infomation,
		});

		resultPlan = await resp.json();
	});
</script>

{#if resultPlan}
	<div class="w-full">
		<div class="mt-7 mb-8">
			<p class="text-center font-sans text-4xl font-semibold">
				{resultPlan.length} Crucial {resultPlan.length > 1 ? "Steps" : "Step"} of
			</p>
			<p class="text-center font-sans text-2xl">Achieving your goal</p>
		</div>

		<div
			id="plan"
			class="flex flex-wrap text-[20px] border-2 items-start justify-evenly max-sm:mx-10 mb-6 mx-20 px-7"
		>
			{#each resultPlan as { time, plan }}
				<div class="my-7 max-w-[450px] space-y-1">
					<p class="text-3xl font-medium">{time}</p>
					<p class="text-xl">{plan}</p>
				</div>
			{/each}
		</div>

		<CaptureScreen />

		<a
			href="/"
			class="fixed bottom-3 left-2 rounded-full bg-lime-500 h-[55px] w-[55px] hover:bg-lime-700"
			data-html2canvas-ignore="true"
			title="Generate again"
		>
			<Image src={mmb} />
		</a>
	</div>
{:else}
	<div class="w-full h-lvh flex justify-center items-center">
		<p class="text-5xl">
			Generating the plan
			<Loading width={60} height={60} class="inline"/>
		</p>
	</div>
{/if}
