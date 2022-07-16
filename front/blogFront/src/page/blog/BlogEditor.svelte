<script>
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link } from "svelte-spa-router";
import Editor from "cl-editor/src/Editor.svelte";

let html = "<p>내용을 입력해주세요.</p>";

let title = "";
let tag = "";

$: console.log(title);

export let params = {
  id: "",
};
</script>

<Header />
<main class="sub-page">
  <section class="sub-blog-wrap">
    <div class="inner">
      <div class="sub-blog">
        <h1 class="sub-page-title">Blog</h1>
        <div class="sub-blog-detail-title">
          <span
            ># <input
              type="text"
              placeholder="태그"
              bind:value="{tag}" /></span>
          <h2><input type="text" placeholder="제목" bind:value="{title}" /></h2>
          <p>
            <textarea
              name=""
              placeholder="블로그에 대한 간단한 내용을 적어주세요."
              style="resize: none;"></textarea>
          </p>
          <div class="sub-blog-detail-info">
            <span>2022년 01월 22일</span>
            <span>|</span>
            <span>yeong</span>
          </div>
        </div>
        <div class="sub-blog-detail-contents">
          <Editor
            on:change="{({ detail }) => (html = detail)}"
            html="{html}"
            height="{'200px'}" />
        </div>
        <div class="sub-blog-detail-btn">
          {#if params.id === "register"}
            <a href="/blog" use:link>등록</a>
          {:else}
            <a href="/blog" use:link>수정</a>
          {/if}

          <a href="/blog" use:link>취소</a>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
