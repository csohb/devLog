<script lang="ts">
import aboutStore from "../../stores/about";
</script>

<h2>Knowledges</h2>
<ul class="keywords-wrap">
  {#if $aboutStore.keywords.length !== 0}
    {#each $aboutStore.keywords as keyword}
      <li>{keyword.name}</li>
    {/each}
  {:else}
    <li>등록 중 입니다.</li>
  {/if}
</ul>
