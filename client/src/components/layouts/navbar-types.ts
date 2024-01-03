import type { ComponentType, SvelteComponent } from "svelte";
import type { SVGAttributes } from "svelte/elements";

export interface SidebarHrefInfo {
	href: `/${string}`;
	image: ComponentType<SvelteComponent<SVGAttributes<SVGElement>>>;
	label: string;
  lastText: string;
}
