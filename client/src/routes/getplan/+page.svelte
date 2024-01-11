<script lang="ts">
	import { onMount } from 'svelte';
	import CaptureScreen from '$components/CaptureScreen.svelte';

	let Data : {time : string, plan : string}[];

	onMount (async () =>{
		const resp = await fetch (process.env.database_plan);
		Data = await resp.json();
		console.log(Data.length);
	});
</script>


<div>
	<div class="pt-[50px] text-center font-sans text-[32px]">
		<h1 class="mb-0">{Data.length} Crucial Steps of</h1>
	</div>
	<div class="pb-[30px] text-center font-serif text-[22px]">
		<p class="mt-0">Achieving gold medal IOI</p>
		<!-- <p class="mt-0">{usergoal}</p> -->
	</div>

	
	<div class="plan">
		<div class="flex flex-wrap font-mono text-[20px]">
			{#each Data as {time, plan}}
				<div class="m-auto p-[20px] text-[30px] w-[450px]">
					<p><b>{time}</b></p>
					<p class="font-mono text-[14px]"> {plan} </p>
				</div>
			{/each}
			</div>
		</div>

	<CaptureScreen />

	<a href="/mmb" class="fixed bottom-0 left-0 rounded-full bg-lime-400 h-[55px] w-[55px]">
		<Image src='/src/lib/return.png' />
	</a>
</div>