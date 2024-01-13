<script>
	import Input from "./Input.svelte";
</script>

<h1 class="text-2xl font-semibold mb-4">Planner wants to know:</h1>
<form class="text-center" id="planner-question-form">
	<div class="mt-4">
		<Input label="Type something..." type="text" id="planner-question-input" required />
	</div>
	<button
		type="submit"
		id="planner-question-submit"
		class="mt-4 bg-gray-800 text-white p-2 rounded-md w-full hover:bg-black"
	>
		Continue
	</button>
</form>
