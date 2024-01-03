<script lang="ts">
	import type { Snippet } from "svelte";
	import type { HTMLAnchorAttributes } from "svelte/elements";

	import { clsx } from "$lib/clsx";

	interface NavLinkProps extends HTMLAnchorAttributes {
		href: string;
		icon?: Snippet<void>;
		isActive?: boolean;
		wideText?: boolean;
		textCenter?: boolean;
	}

	const {
		href,
		icon,
		isActive = false,
		wideText = false,
		textCenter = true,
		children,
		...props
	} = $props<NavLinkProps>();
</script>

<span
	class={clsx(
		"w-full transition-colors-opacity duration-100 rounded-md flex flex-row justify-between cursor-pointer group",
		isActive ? "bg-slate-300 dark:bg-slate-700" : "hover:bg-slate-300 dark:hover:bg-slate-700",
	)}
>
	<a
		{href}
		class={clsx(
			"w-full h-full font-medium gap-2 px-3 py-2 text-black dark:text-white flex items-center flex-row",
			textCenter && "text-center",
			wideText ? "shrink-0 text-base uppercase tracking-widest" : "text-sm",
		)}
		aria-current={isActive}
		{...props}
	>
    {#if icon}
      {@render icon()}
    {/if}
		{#if children}
			{@render children()}
		{/if}
	</a>
</span>
