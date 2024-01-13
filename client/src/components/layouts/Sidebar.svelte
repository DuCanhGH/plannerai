<script lang="ts">
	import { quintOut } from "svelte/easing";
	import { slide } from "svelte/transition";

	import { page } from "$app/stores";
	import ChatGPT from "$components/icons/ChatGPT.svelte";
	import { clsx } from "$lib/clsx";
	import { isLinkActive } from "$lib/isLinkActive";

	import { SIDEBAR_LINKS } from "./navbar-constants";
	import SidebarLink from "./SidebarLink.svelte";

	const links = $derived(
		SIDEBAR_LINKS.map(({ href, ...rest }) => ({
			href,
			...rest,
			isActive: isLinkActive(href, $page.url.pathname),
		})),
	);

	let isNavMobileMenuOpened = $state(false);
</script>

<nav
	class={clsx(
		"print:hidden w-full max-h-screen md:w-80 md:shrink-0 z-[50]",
		"transform-gpu transition-all duration-150 ease-out sticky top-0",
		"p-2 bg-white dark:bg-slate-800 md:bg-transparent dark:md:bg-transparent",
		"border-neutral-300 border-b-[0.25px] md:border-b-0 dark:border-gray-700",
	)}
>
	<div class="relative flex justify-between flex-row md:flex-col">
		<div class="absolute inset-y-0 left-0 flex items-center gap-2 md:hidden">
			<input
				type="checkbox"
				id="navbar-mobile-menu-toggle"
				class="hidden"
				aria-labelledby="navbar-mobile-menu-toglab"
				aria-expanded={isNavMobileMenuOpened}
				aria-controls="navbar-mobile-menu"
				bind:checked={isNavMobileMenuOpened}
			/>
			<label
				id="navbar-mobile-menu-toglab"
				for="navbar-mobile-menu-toggle"
				class={clsx(
					"flex h-[2rem] w-[2rem] cursor-pointer flex-col justify-center gap-[0.5rem] md:hidden",
					"[&>span]:bg-black [&>span]:transition-all [&>span]:dark:bg-white",
					"[&>span]:h-[0.2rem] [&>span]:w-full [&>span]:rounded-md",
					isNavMobileMenuOpened && [
						"[&>:nth-child(1)]:rotate-45",
						"[&>:nth-child(1)]:translate-y-[0.7rem]",
						"[&>:nth-child(2)]:opacity-0",
						"[&>:nth-child(3)]:-translate-y-[0.7rem]",
						"[&>:nth-child(3)]:-rotate-45",
					],
				)}
				aria-label="Toggle navbar menu"
			>
				<span class="origin-center duration-300" />
				<span class="duration-200 ease-out" />
				<span class="origin-center duration-300" />
			</label>
		</div>
		<div
			class="flex flex-1 items-center justify-center md:flex-none md:items-stretch md:justify-start [&>*]:px-0 md:mb-[5px]"
		>
			<span>
				<SidebarLink href="/" aria-label="Go to home">
					{#snippet icon()}
						<ChatGPT
							width={20}
							height={20}
							class="rounded-full max-w-[20px] min-w-[20px] max-h-[20px] min-h-[20px]"
							aria-hidden="true"
						/>
					{/snippet}
					Planner
				</SidebarLink>
			</span>
		</div>
		<div
			class="absolute inset-y-0 right-0 flex h-full w-fit gap-[5px] md:static md:w-full flex-col"
		>
			<div class="hidden h-full grow overflow-x-hidden pr-2 md:flex md:pr-0">
				<div
					class="overflow-x-overlay hidden h-full grow flex-row gap-[5px] overflow-x-auto md:flex"
				>
					<ul class="w-full flex max-h-full gap-[inherit] flex-col">
						{#each links as { label, image: Image, href, isActive, lastText }}
							<li class="w-full">
								<SidebarLink {href} textCenter={false} {isActive}>
									{#snippet icon()}
										<Image
											width={40}
											height={40}
											class="rounded-full max-w-[40px] min-w-[40px] max-h-[40px] min-h-[40px]"
										/>
									{/snippet}
									<span class="flex flex-col gap-[2px]">
										<span class="text-lg">
											{label}
										</span>
										{lastText}
									</span>
								</SidebarLink>
							</li>
						{/each}
					</ul>
				</div>
			</div>
			<div class="flex flex-row gap-[5px] items-center"></div>
		</div>
	</div>
	{#if isNavMobileMenuOpened}
		<ul
			class="space-y-1 px-2 pb-3 pt-2 md:hidden"
			id="navbar-mobile-menu"
			transition:slide={{ duration: 200, easing: quintOut, axis: "y" }}
		>
			{#each links as { label, lastText, image: Image, href, isActive }}
				<li>
					<SidebarLink {href} textCenter={false} {isActive}>
						{#snippet icon()}
							<Image
								width={30}
								height={30}
								class="rounded-full max-w-[30px] min-w-[30px] max-h-[30px] min-h-[30px]"
							/>
						{/snippet}
						<span class="flex flex-col gap-[2px]">
							<span class="text-base">
								{label}
							</span>
							{lastText}
						</span>
					</SidebarLink>
				</li>
			{/each}
		</ul>
	{/if}
</nav>
