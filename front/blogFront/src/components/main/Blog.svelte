<script lang="ts">
import { onDestroy, onMount } from "svelte";
import blogStore from "../../stores/blog";
import { link } from "svelte-spa-router";

const PAGE = 1;
const COUNT = 5;

onMount(() => {
  blogStore.setBlogList(PAGE, COUNT);
});

function onClickViewCount(id: string) {
  blogStore.viewCount(id);
}

onDestroy(() => {
  blogStore.resetBlogList();
});
</script>

<div class="inner">
  <div class="main-blog-wrap">
    <h1>BLOG</h1>
    {#if $blogStore.blogList.length === 0}
      등록된 내용이 없습니다.
    {:else}
      <ul class="main-blog-list">
        {#each $blogStore.blogList as item}
          <li class="main-blog-list-item">
            <a
              href="/blog/{item.id}"
              use:link
              on:click="{() => {
                onClickViewCount(item.id);
              }}">
              <h2 class="main-blog-list-title">
                {item.title}
              </h2>
              <span class="main-blog-list-text"
                ><strong>{item.writer}</strong> {item.date}</span>
            </a>
          </li>
        {/each}
        <!-- <li class="main-blog-list-item">
        <a href="#none">
          <h2 class="main-blog-list-title">제목</h2>
          <span class="main-blog-list-text"
            ><strong>작성자,</strong> 2022. 06. 04</span>
        </a>
      </li>
      <li class="main-blog-list-item">
        <a href="#none">
          <h2 class="main-blog-list-title">제목</h2>
          <span class="main-blog-list-text"
            ><strong>작성자,</strong> 2022. 06. 04</span>
        </a>
      </li> -->
      </ul>
    {/if}
  </div>
</div>
