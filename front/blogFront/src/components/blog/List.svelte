<script lang="ts">
import Pagination from "../Pagination.svelte";
import { link } from "svelte-spa-router";

function onClickDetailPage(id: number): void {
  console.log(id);
}
</script>

<ul class="sub-blog-list">
  {#each Array(10) as _, index}
    <li class="sub-blog-item">
      <a href="/blog/{index}" use:link>
        <span># svelte</span>
        <h2>제목</h2>
        <p>말줄임 필요한</p>
        <div class="sub-blog-item-info">
          <span>2022년 01월 22일</span>
          <span>|</span>
          <span>yeong</span>
        </div>
      </a>
    </li>
  {/each}
</ul>

<Pagination />
