<script lang="ts">
	import { onMount } from "svelte";

	import CaptureScreen from "$components/CaptureScreen.svelte";
	import Image from "$components/Image.svelte";
	import mmb from "$lib/return.png"
	import { onMount } from "svelte";
	import CaptureScreen from "$components/CaptureScreen.svelte";
	import Image from "$components/Image.svelte";
	import mmb from "$lib/return.png";
	import { PUBLIC_SERVER_URL } from "$env/static/public";

	interface PlanData {
		time: string;
		plan: string;
	}

	let userGoal: string;
	let resultPlan: PlanData[] | undefined = $state();

	onMount(async () => {
		let infomation =
			'{"usergoal": "I want to get IOI gold medal","time": "2 years","Queries": [{"planner": "1. Have you participated in the IOI competition before? If yes, what was your previous performance?","user": "I\'ve got bronze"},{"planner": "2. How much time are you willing to dedicate to preparation for the IOI competition?","user": "12h per day!"}]}';
		if (infomation == "") {
			window.location.href = "/";
			return;
		}
		userGoal = JSON.parse(infomation).usergoal;
		// const resp = await fetch(`${PUBLIC_SERVER_URL}/getplan`, {
		// 	method: "POST",
		// 	headers: {
		// 		"Content-Type": "application/json",
		// 	},
		// 	body: infomation,
		// });
		resultPlan = JSON.parse(
			'[{"time": "Year 1, Month 1", "plan": "Start by familiarizing yourself with the IOI competition format and past problems. Understand the scoring system and the skills required for each task.", "resources": "1) IOI Official Website - https://ioinformatics.org/<br/>2) IOI Syllabus - https://ioinformatics.org/page/syllabus/0<br/>3) IOI Past Problems - https://ioinformatics.org/page/past-contests/0"}, {"time": "Year 1, Month 3", "plan": "Begin practicing with easier programming problems to build a strong foundation. Focus on algorithms and data structures.", "resources": "1) Codeforces - https://codeforces.com/<br/>2) LeetCode - https://leetcode.com/<br/>3) Topcoder - https://www.topcoder.com/"},{"time": "Year 1, Month 6", "plan": "Start solving medium difficulty IOI-style problems to improve your problem-solving skills. Analyze your solutions and learn from your mistakes.", "resources": "1) CodeChef - https://www.codechef.com/<br/>2) AtCoder - https://atcoder.jp/<br/>3) Project Euler - https://projecteuler.net/"},{"time": "Year 1, Month 9", "plan": "Continue practicing with IOI-style problems and participate in online coding competitions to gain experience under time pressure.", "resources": "1) HackerRank - https://www.hackerrank.com/<br/>2) CodeSignal - https://codesignal.com/<br/>3) USACO Training Program - http://www.usaco.org/index.php?page=training"},{"time": "Year 1, Month 12", "plan": "Review your progress and identify areas of improvement. Focus on strengthening weaker topics and solving more challenging problems.", "resources": "1) Competitive Programming 3 by Steven Halim - https://cpbook.net/<br/>2) Introduction to Algorithms by Thomas H. Cormen - https://mitpress.mit.edu/books/introduction-algorithms<br/>3) GeeksforGeeks - https://www.geeksforgeeks.org/"},{"time": "Year 2, Month 3", "plan": "Start participating in local programming contests and aim for a top position. This will help you gain confidence and exposure to competition-like scenarios.", "resources": "1) ACM ICPC Live Archive - https://icpcarchive.ecs.baylor.edu/<br/>2) CodeChef Contests - https://www.codechef.com/contests<br/>3) Topcoder SRMs - https://www.topcoder.com/community/competitive-programming/tutorials/single-round-match-tutorials/"},{"time": "Year 2, Month 6", "plan": "Continue solving challenging IOI-style problems and participate in regional coding competitions to test your skills against stronger competitors.", "resources": "1) Google Code Jam - https://codingcompetitions.withgoogle.com/codejam<br/>2) Facebook Hacker Cup - https://www.facebook.com/codingcompetitions/hacker-cup<br/>3) International Collegiate Programming Contest (ICPC) - https://icpc.global/"},{"time": "Year 2, Month 9", "plan": "Focus on solving past IOI problems and simulating the competition environment. Practice solving problems within the time limit and optimize your solutions.", "resources": "1) IOI Past Problems - https://ioinformatics.org/page/past-contests/0<br/>2) Codeforces Gym - https://codeforces.com/gyms<br/>3) AtCoder Beginner Contest - https://atcoder.jp/contests/abc"},{"time": "Year 2, Month 12", "plan": "Revise all the important concepts and algorithms. Solve a variety of challenging problems to strengthen your problem-solving skills.", "resources": "1) Competitive Programmer\'s Handbook by Antti Laaksonen - https://cses.fi/book/index.html<br/>2) LeetCode Explore - https://leetcode.com/explore/<br/>3) CodeChef Long Challenges - https://www.codechef.com/LTIME"},{"time": "Year 2, Month 15", "plan": "Focus on practicing under timed conditions and solving IOI-style problems. Analyze your performance and make necessary adjustments.", "resources": "1) Codeforces Educational Rounds - https://codeforces.com/contests/educational<br/>2) HackerRank Contests - https://www.hackerrank.com/contests<br/>3) USACO Training Program - http://www.usaco.org/index.php?page=training"}]',
		);
	});
</script>

{#if resultPlan}
	<div class="w-full">
		<div class="my-7">
			<p class="text-center font-sans text-4xl font-semibold">
				{resultPlan.length} Crucial {resultPlan.length > 1 ? "Steps" : "Step"} of
			</p>
			<p class="text-center font-sans text-2xl">Achieving your goal</p>
		</div>

		<div id="plan" class="flex flex-wrap text-[20px] border-2 items-start justify-evenly max-sm:mx-10 mb-6 mx-20 px-7">
			{#each resultPlan as { time, plan }}
				<div class="my-7 max-w-[450px] space-y-1">
					<p class="text-3xl font-medium">{time}</p>
					<p class="text-xl">{plan}</p>
				</div>
			{/each}
		</div>

		<CaptureScreen />

		<a
			href="/"
			class="fixed bottom-3 left-2 rounded-full bg-lime-500 h-[55px] w-[55px] hover:bg-lime-700"
			title="Generate again"
		>
			<Image src={mmb} />
		</a>
	</div>
{:else}
	<p>Loading</p>
{/if}
