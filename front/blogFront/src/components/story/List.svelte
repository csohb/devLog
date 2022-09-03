<script lang="ts">
import Pagination from "../Pagination.svelte";
import { link } from "svelte-spa-router";
import { onMount } from "svelte";
import storyStore from "../../stores/story";

let page: number = 1;
let count: number = 10;

const TOTAl: number = 100;

function setPagination(e: any) {
  // 페이지 이동 버튼 클릭 시 이벤트
  page = e.detail + 1;
}
onMount(() => {
  storyStore.setStoryList(1, 10);
});
</script>

<ul class="sub-story-list">
  {#each Array(5) as _, idx}
    <li class="sub-story-item">
      <a href="/story/{idx}" use:link>
        <div class="sub-story-content-img">
          <img
            src="https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg"
            alt="sample" />
        </div>
        <div class="sub-story-content-text">
          <span>Backend</span>
          <h2>제목</h2>
          <p>devblog를 만들게 된 계기~</p>
        </div>
      </a>
    </li>
  {/each}
</ul>

<Pagination
  bind:page
  count="{TOTAl}"
  bind:pageSize="{count}"
  on:change="{setPagination}" />
