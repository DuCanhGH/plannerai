import { error, redirect } from "@sveltejs/kit";

import type { Actions } from "./$types";

export const actions: Actions = {
	async default({ request }) {
		const formData = await request.formData();
		const usergoal = formData.get("usergoal") as string;
		const time = formData.get("time") as string;
		if (usergoal === null || time === null) {
			error(400, "You must provide an answer!");
		}
		throw redirect(301, `/question?usergoal=${encodeURIComponent(usergoal)}&time=${encodeURIComponent(time)}`);
	},
};
