import ChatGPT from "$components/icons/ChatGPT.svelte";

import type { SidebarHrefInfo } from "./navbar-types";

export const SIDEBAR_LINKS = [
	{ image: ChatGPT, href: "/", label: "ChatGPT", lastText: "me may beo" },
] satisfies SidebarHrefInfo[];
