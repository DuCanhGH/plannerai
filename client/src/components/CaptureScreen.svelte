<script>
	import html2canvas from 'html2canvas';
	import Image from './Image.svelte';

	async function CaptureScreen() {
		const canvas = await html2canvas(document.body);
		const dataURL = canvas.toDataURL();
		const link = document.createElement('a');
		link.href = dataURL;
		link.download = 'GetPlan.png';
		document.body.appendChild(link);
		link.click();
		document.body.removeChild(link);
	}
</script>

<button on:click={CaptureScreen} class="fixed bottom-0 right-0 rounded-full bg-lime-400 h-[55px] w-[55px]">
	<Image src="src/lib/camera.png" />
</button>