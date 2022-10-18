<script lang="ts">
import Pagination from "../Pagination.svelte";
import { link } from "svelte-spa-router";
import { onMount } from "svelte";
import storyStore from "../../stores/story";

let page: number = 1;
let count: number = 9;

onMount(() => {
  storyStore.setStoryList(page, count);
});

function setPagination(e: any) {
  // 페이지 이동 버튼 클릭 시 이벤트
  page = e.detail + 1;
  storyStore.setStoryList(page, count);
}

function onClickViewCount(id: string) {
  // 글을 읽은 사람 체크
  storyStore.viewCount(id);
}
</script>

<ul class="sub-story-list">
  {#if $storyStore.storyList.length !== 0}
    {#each $storyStore.storyList as item, idx}
      <li class="sub-story-item">
        <a
          href="/story/{item.id}"
          use:link
          on:click="{() => {
            onClickViewCount(item.id);
          }}">
          <div class="sub-story-content-img">
            <img
              src="{item.image === ''
                ? 'https://picsum.photos/seed/picsum/536/536'
                : item.image}"
              alt="sample" />
          </div>
          <div class="sub-story-content-text">
            <!-- TODO: 셀렉트박스로 선택할 수 있도록 하기 -->
            <span>{item.type ? item.type : "Etc"}</span>
            <h2>{item.title}</h2>
            <p>{@html item.description}</p>
          </div>
        </a>
      </li>
    {/each}
  {:else}
    <li class="sub-story-item">등록된 내용이 없습니다.</li>
  {/if}
</ul>

{#if $storyStore.storyList.length !== 0}
  <Pagination
    bind:page
    count="{$storyStore.listTotal}"
    bind:pageSize="{count}"
    on:change="{setPagination}" />
{/if}
