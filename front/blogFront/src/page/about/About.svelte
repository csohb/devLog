<script lang="ts">
import Project from "../../components/about/Project.svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
</script>

<Header />
<main class="sub-page">
  <section class="sub-about-wrap">
    <div class="inner">
      <div class="sub-about">
        <ul class="sub-about-tab">
          <li class="foc">YEONG</li>
          <li>CSOHB</li>
        </ul>
        <h1>About Me</h1>
        <div class="sub-about-contents">
          <div class="sub-about-introduce">
            <pre>안녕하세요, 코딩하는 것이 즐거운 행복한 개발자입니다.
저는 현재 백엔드 개발을 하고 있으며, 주로 Python 기반 백엔드 프레임워크를 다루고 있습니다.
이곳은 저를 소개하는 동시에, 저의 개발 성장 일지를 기록하는 공간입니다. 어떻게 보면 일종의 회고록으로도 볼 수 있겠네요. 😅
저에게 문의할 내용이 있으면 언제든지 연락주셔도 좋습니다.
감사합니다 !</pre>
          </div>
          <div class="sub-about-info">
            <h2>Info</h2>
            <ul>
              <li><strong>Name</strong>김영지 ( yeong )</li>
              <li><strong>Date of birth</strong>Feb / 26 / 1993</li>
              <li><strong>Address</strong>경기도 수원시</li>
              <li><strong>E-mail</strong>kyj210@gmail.com</li>
            </ul>
          </div>
          <div class="sub-about-career"></div>
          <div class="sub-about-project">
            <Project />
          </div>
          <div class="sub-about-skills"></div>
          <div class="sub-about-keywords"></div>
        </div>
      </div>
    </div>
  </section>
</main>

<Footer />
