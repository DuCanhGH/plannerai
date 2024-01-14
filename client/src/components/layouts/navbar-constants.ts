import ChatGPT from "$components/icons/ChatGPT.svelte";

import type { SidebarHrefInfo } from "./navbar-types";

export const SIDEBAR_LINKS = [
	{ image: ChatGPT, href: "/", label: "New", lastText: "Plan to your goal" },
] satisfies SidebarHrefInfo[];
