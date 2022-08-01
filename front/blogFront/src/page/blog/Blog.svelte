<script lang="ts">
import List from "../../components/blog/List.svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link } from "svelte-spa-router";
import blogStore from "../../stores/blog";
import authStore from "../../stores/auth";

let page: number = 1;
let count: number = 10;

function onClickTag(tag: string) {
  page = 1;
  count = 10;
  blogStore.setTagList(tag, page, count);
}
</script>

<Header />
<main class="sub-page">
  <section class="sub-blog-wrap">
    <div class="inner">
      <div class="sub-blog">
        <h1 class="sub-page-title">Blog</h1>
        <ul class="sub-blog-tag">
          {#each $blogStore.tagList as tag, index}
            <li>
              <a
                href="#none"
                on:click|preventDefault="{() => {
                  onClickTag(tag);
                }}"># {tag}</a>
            </li>
          {/each}
        </ul>
        <div class="sub-blog-contents">
          <List />
        </div>
        {#if $authStore.loginNick !== ""}
          <a href="/blog/edit/register" use:link>등록하기</a>
        {/if}
      </div>
    </div>
  </section>
</main>
<Footer />
