<script>
  import LayoutDefault from '$components/defaultlayout.svelte';
  import LayoutFinal from '$components/finallayout.svelte';
  import Page1 from '$components/page1.svelte';
  import Page2 from '$components/page2.svelte';
  import Page3 from '$components/final.svelte';

  let currentPage = 'page1';

  const handleNextPage = () => {
    if (currentPage === 'page1') {
      currentPage = 'page2';
    } else if (currentPage === 'page2') {
      currentPage = 'page3';
    }
  };
</script>

{#if currentPage === 'page1'}
  <LayoutDefault>
    <Page1 on:nextPage={handleNextPage} />
  </LayoutDefault>
{:else if currentPage === 'page2'}
  <LayoutDefault>
    <Page2 on:nextPage={handleNextPage} />
  </LayoutDefault>
{:else if currentPage === 'page3'}
  <LayoutFinal>
    <Page3 />
  </LayoutFinal>
{/if}
