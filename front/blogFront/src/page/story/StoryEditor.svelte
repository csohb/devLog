<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link, push } from "svelte-spa-router";

export let params = {
  id: "",
};

function onClickSave() {}

function onClickUpdate() {}

function onClickCancel() {
  if (params.id === "register") {
    push("/story");
  } else {
    push(`/story/${params.id}`);
  }
}
</script>

<Header />
<main class="sub-page">
  <section class="sub-story-wrap">
    <div class="inner">
      <div class="sub-story">
        <h1 class="sub-page-title">Story</h1>
        <div class="sub-story-detail-title">
          <span> Backend </span>
          <h2>제목</h2>
          <p>devblog를 시작한 계기</p>
          <div class="sub-story-detail-thumb">
            <img
              src="https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg"
              alt="" />
          </div>
          <div class="sub-story-detail-info">
            <span>2022-07-24</span>
            <span>|</span>
            <span>yeong</span>
          </div>
        </div>
        <div class="sub-story-detail-contents">
          {#if params.id === "register"}등록{:else}수정{/if}
        </div>
        <div class="sub-story-detail-btn">
          {#if params.id === "register"}
            <a href="#none" on:click|preventDefault="{onClickSave}">등록</a>
          {:else}
            <a href="#none" on:click|preventDefault="{onClickUpdate}">수정</a>
          {/if}

          <a href="#none" on:click|preventDefault="{onClickCancel}">취소</a>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
