<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link } from "svelte-spa-router";
import { unlike, like } from "../../icon/Icon";

let isLike = false;
let likeCount = 0;

export let params = {
  id: "",
};

function onClickLike() {
  isLike = !isLike;
  if (isLike) {
    likeCount += 1;
  } else {
    likeCount -= 1;
  }
}

function onClickDelete() {}
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
          <div class="sub-story-detail-info">
            <span>2022-07-24</span>
            <span>|</span>
            <span>yeong</span>
            <span>|</span>
            <span>0</span>
          </div>
          <div class="sub-story-detail-thumb">
            <img
              src="https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg"
              alt="" />
          </div>
        </div>
        <div class="sub-story-detail-contents">내용내용 ㄴ용내용</div>
        <div class="sub-story-detail-action">
          <div class="sub-story-liked" on:click="{onClickLike}">
            {#if isLike}
              {@html like}
            {:else}
              {@html unlike}
            {/if}
            {likeCount}
          </div>
        </div>
        <div class="sub-story-detail-btn">
          <a href="/story" use:link>목록으로 돌아가기</a>
          <a href="/story/edit/{params.id}" use:link>수정하기</a>
          <button type="button" on:click="{onClickDelete}">삭제</button>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
