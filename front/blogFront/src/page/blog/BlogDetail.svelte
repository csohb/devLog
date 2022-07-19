<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link, push } from "svelte-spa-router";
import { onDestroy, onMount } from "svelte";
import blogStore from "../../stores/blog";

export let params = {
  id: "",
};

onMount(() => {
  if (!params.id) {
    push("/blog");
    return;
  }
  blogStore.setBlogDetail(params.id);
});

function setTags(tagList: string[]): any {
  // 태그 리스트에 # 추가하기
  return tagList.reduce((temp: any, item: string) => {
    temp.push(`# ${item}`);
    return temp;
  }, []);
}

onDestroy(() => {
  blogStore.resetBlogDetail();
});
</script>

<Header />
<main class="sub-page">
  <section class="sub-blog-wrap">
    <div class="inner">
      <div class="sub-blog">
        <h1 class="sub-page-title">Blog</h1>
        {#if !!$blogStore.blogDetail.id}
          <div class="sub-blog-detail-title">
            <span>
              {#each setTags($blogStore.blogDetail.tags) as item}
                {item}&nbsp;&nbsp;&nbsp;&nbsp;
              {/each}
            </span>
            <h2>{$blogStore.blogDetail.title}</h2>
            <p>
              {$blogStore.blogDetail.description}
            </p>
            <div class="sub-blog-detail-info">
              <span>{$blogStore.blogDetail.date}</span>
              <span>|</span>
              <span>{$blogStore.blogDetail.writer}</span>
            </div>
          </div>
          <div class="sub-blog-detail-contents">
            {@html $blogStore.blogDetail.content}
          </div>
        {/if}
        <div class="sub-blog-detail-btn">
          <a href="/blog" use:link>목록으로 돌아가기</a>
          <a href="/blog/edit/{params.id}" use:link>수정하기</a>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
