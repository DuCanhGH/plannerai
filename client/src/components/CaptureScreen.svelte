<script lang="ts">
	import html2canvas from "html2canvas";

	async function captureScreen() {
		const canvas = await html2canvas(document.body);
		const dataURL = canvas.toDataURL();
		const link = document.createElement("a");
		link.href = dataURL;
		link.download = "GetPlan.png";
		document.body.appendChild(link);
		link.click();
		document.body.removeChild(link);
	}
</script>

<button
	on:click={captureScreen}
	class="fixed bottom-3 right-2 rounded-full bg-lime-500 h-[55px] w-[55px] hover:bg-lime-700"
	data-html2canvas-ignore="true"
	title="Capture the plan"
>
	<enhanced:img src="$images/camera.png" alt="Capture" class="w-[100px] h-auto" />
</button>
