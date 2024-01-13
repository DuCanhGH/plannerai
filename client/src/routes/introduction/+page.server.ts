import { error, redirect } from "@sveltejs/kit";

import type { Actions } from "./$types";

export const actions: Actions = {
	async default({ request }) {
		const formData = await request.formData();
		const userAnswer = formData.get("userAnswer");
		if (userAnswer === null) {
			error(400, "You must provide an answer!");
		}
		throw redirect(301, "/chat");
	},
};
