import { error, redirect } from "@sveltejs/kit";

import type { Actions } from "./$types";

export const actions: Actions = {
	async default({ request }) {
		const formData = await request.formData();
		const formUserGoal = formData.get("usergoal");
		const formTime = formData.get("time") as string;
		if (typeof formUserGoal !== "string" || typeof formTime !== "string") {
			error(400, "You must provide an answer!");
		}
		const userGoal = encodeURIComponent(formUserGoal);
		const time = encodeURIComponent(formTime);
		throw redirect(301, `/question?usergoal=${userGoal}&time=${time}`);
	},
};
