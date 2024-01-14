<script>
	import "../app.css";

	import { page } from "$app/stores";

	const title = $derived($page.data.title ? `${$page.data.title} - Planner` : "Planner");
</script>

<a
	class="absolute -top-full z-[100] text-black underline focus:top-0 dark:text-white"
	href="#main-content"
>
	Skip to main content
</a>
<svelte:head>
	<title>{title}</title>
	<meta property="og:title" content={title} />
	<meta name="twitter:title" content={title} />
	<link rel="canonical" href={$page.url.href} />
</svelte:head>

<slot />
