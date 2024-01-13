<script>
  import { onMount } from 'svelte';
  import { createEventDispatcher } from 'svelte';

  let inputArray = [];
  let inputValue = '';
  const dispatch = createEventDispatcher();

  const handleFormSubmission = () => {
    inputArray.push(inputValue);
    dispatch('nextPage');
    console.log('Input Array:', inputArray);
    inputValue = '';
  };

  onMount(() => {
    const inputForm = document.getElementById('inputForm');
    const inputField = document.getElementById('inputField');
    const continueButton = document.getElementById('storeInput');

    inputForm.addEventListener('submit', (event) => {
      event.preventDefault();
      handleFormSubmission();
    });

    continueButton.addEventListener('click', handleFormSubmission);

    inputField.addEventListener('keypress', (event) => {
      if (event.keyCode === 13) {
        handleFormSubmission();
      }
    });
  });
</script>

<div class="bg-gray-800 text-white">
  <div class="container mx-auto flex items-center justify-center h-screen">
    <div class="bg-gray-700 p-8 rounded-lg shadow-lg w-96">
      <h1 class="text-2xl font-semibold mb-4">PLANNER want to know: </h1>
      <form on:submit|preventDefault={handleFormSubmission} class="text-center" id="inputForm">
        <div class="mt-4">
          <label for="inputField" class="block text-gray-200">How long can you spend? </label>
          <input bind:value={inputValue} type="text" id="inputField" class="mt-1 p-2 border rounded-md w-full text-gray-800" placeholder="Type something..." required />
        </div>
        <button type="submit" id="storeInput" class="mt-4 bg-gray-800 text-white p-2 rounded-md w-full hover:bg-black">Continue</button>
      </form>
    </div>
  </div>
</div>
