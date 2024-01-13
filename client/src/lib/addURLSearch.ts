export const addURLSearch = (url: URL, searchParams: [string, string][]) => {
	return new URL(
		`${url.origin}${url.pathname}?${new URLSearchParams([
			...url.searchParams.entries(),
			...searchParams,
		])}`,
	);
};
