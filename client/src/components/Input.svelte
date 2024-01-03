<script lang="ts">
	import type { HTMLInputAttributes } from "svelte/elements";

	import { clsx } from "$lib/clsx";

	import ErrorText from "./ErrorText.svelte";

	interface InputProps extends Omit<HTMLInputAttributes, "placeholder"> {
		label: string;
		id: string;
		errorTextId?: string;
		errorText?: string;
		sameLine?: boolean;
		value?: any;
	}

	const {
		label,
		id,
		errorTextId,
		errorText,
		value,
		...rest
	} = $props<InputProps>();
</script>

<div class="relative">
	<style>
		.input:placeholder-shown:focus + .label {
			@apply top-0.5 translate-y-0 text-xs text-neutral-700 dark:text-gray-300;
		}
		.input:placeholder-shown:not(:focus) + .label {
			@apply top-1/2 -translate-y-1/2 text-black dark:text-white text-sm;
		}
		.input:not(:placeholder-shown) +.label {
			@apply top-0.5 text-xs text-neutral-700 dark:text-gray-300;
		}
	</style>
	<input
		{id}
		class={clsx(
			"input block h-[44px] rounded-lg text-sm shadow-md transition-opacity disabled:opacity-50 w-full px-2.5 pt-2.5",
			"focus:border-accent-light dark:focus:border-accent-dark border border-neutral-400 focus:outline-none dark:border-neutral-700",
			"bg-slate-100 text-black dark:bg-slate-700 dark:text-white",
		)}
		aria-invalid={!!errorText}
		aria-describedby={errorTextId}
		placeholder=" "
		{...rest}
	/>
	<label
		class="label absolute left-2.5 block font-medium transition-all duration-100 ease-in"
		for={id}
	>
		{label}
	</label>
	{#if !!errorText && errorTextId}
		<ErrorText id={errorTextId}>{errorText}</ErrorText>
	{/if}
</div>
