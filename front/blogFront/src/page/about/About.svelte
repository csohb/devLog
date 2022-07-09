<script lang="ts">
import Career from "../../components/about/Career.svelte";
import Introduce from "../../components/about/Introduce.svelte";
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
          <Introduce />
          <div class="sub-about-career">
            <Career />
          </div>
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
