<script lang="ts">
import Pagination from "../Pagination.svelte";
import { link } from "svelte-spa-router";
import { onMount } from "svelte";
import blogStore from "../../stores/blog";

let page: number = 1;
let count: number = 10;

onMount(() => {
  blogStore.setBlogList(page, count);
});

function setTags(tagList: string[]): any {
  // 태그 리스트에 # 추가하기
  return tagList.reduce((temp: any, item: string) => {
    temp.push(`# ${item}`);
    return temp;
  }, []);
}

function setPagination(e: any) {
  // 페이지 이동 버튼 클릭 시 이벤트
  page = e.detail + 1;
  blogStore.setBlogList(page, count);
}
</script>

<ul class="sub-blog-list">
  {#each $blogStore.blogList as _, index}
    <li class="sub-blog-item">
      <a href="/blog/{_.id}" use:link>
        <span>
          {#each setTags(_.tags) as item}
            {item}&nbsp;&nbsp;&nbsp;&nbsp;
          {/each}
        </span>
        <h2>{_.title}</h2>
        <p>{@html _.description}</p>
        <div class="sub-blog-item-info">
          <span>{_.date}</span>
          <span>|</span>
          <span>{_.writer}</span>
        </div>
      </a>
    </li>
  {/each}
</ul>

<Pagination
  bind:page
  count="{$blogStore.listTotal}"
  bind:pageSize="{count}"
  on:change="{setPagination}" />
