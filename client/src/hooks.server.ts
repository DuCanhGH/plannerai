import { type Handle, redirect } from "@sveltejs/kit";

const permanentRedirects: Record<string, string> = {
	"/": "/introduction",
};

export const handle: Handle = async ({ event, resolve }) => {
	if (permanentRedirects[event.url.pathname]) {
		throw redirect(302, permanentRedirects[event.url.pathname]);
	}
	return await resolve(event);
};
