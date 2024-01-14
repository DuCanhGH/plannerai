<script>
	import "../../app.css";

	import Sidebar from "$components/layouts/Sidebar.svelte";
</script>

<a
	class="absolute -top-full z-[100] text-black underline focus:top-0 dark:text-white"
	href="#main-content"
>
	Skip to main content
</a>
<div class="flex flex-grow w-full flex-col md:flex-row">
	<Sidebar />
	<main id="main-content" class="w-full min-w-0">
		<slot />
	</main>
</div>
