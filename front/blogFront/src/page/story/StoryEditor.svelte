<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link } from "svelte-spa-router";
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
        <div class="sub-story-detail-contents">수정</div>
        <div class="sub-story-detail-btn">
          <a href="/story" use:link>목록으로 돌아가기</a>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
