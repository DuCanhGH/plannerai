<script lang="ts">
	import type { HTMLAttributes } from "svelte/elements";

    type ErrorTextProps = Omit<HTMLAttributes<HTMLParagraphElement>, "class">;

    const { ...props } = $props<ErrorTextProps>();
</script>

<p class="text-red-500 dark:text-red-400" {...props} />