<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link } from "svelte-spa-router";

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
          <span># svelte</span>
          <h2>제목</h2>
          <p>
            간략한 내용에 대한 소개를 하는 부분입니다..간략한 내용에 대한 소개를
            하는 부분입니다..간략한 내용에 대한 소개를 하는 부분입니다..간략한
            내용에 대한 소개를 하는 부분입니다..간략한 내용에 대한 소개를 하는
            부분입니다..간략한 내용에 대한 소개를 하는 부분입니다..간략한 내용에
            대한 소개를 하는 부분입니다..간략한 내용에 대한 소개를 하는
            부분입니다..간략한 내용에 대한 소개를 하는 부분입니다..간략한 내용에
            대한 소개를 하는 부분입니다..간략한 내용에 대한 소개를 하는
            부분입니다..간략한 내용에 대한 소개를 하는 부분입니다..간략한 내용에
            대한 소개를 하는 부분입니다..
          </p>
          <div class="sub-blog-detail-info">
            <span>2022년 01월 22일</span>
            <span>|</span>
            <span>yeong</span>
          </div>
        </div>
        <div class="sub-blog-detail-contents">내용내용내용{params.id}</div>
        <div class="sub-blog-detail-btn">
          <a href="/blog" use:link>목록으로 돌아가기</a>
          <a href="/blog/edit/{params.id}" use:link>수정하기</a>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
