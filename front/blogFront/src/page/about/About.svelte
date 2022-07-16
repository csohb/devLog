<script lang="ts">
import Career from "../../components/about/Career.svelte";
import Introduce from "../../components/about/Introduce.svelte";
import Knowledges from "../../components/about/Knowledges.svelte";
import Project from "../../components/about/Project.svelte";
import Skills from "../../components/about/Skills.svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
</script>

<Header />
<main class="sub-page">
  <section class="sub-about-wrap">
    <div class="inner">
      <div class="sub-about">
        <h1 class="sub-page-title">About Me</h1>
        <ul class="sub-about-tab">
          <li class="foc">YEONG</li>
          <li>CSOHB</li>
        </ul>
        <div class="sub-about-contents">
          <Introduce />
          <div class="sub-about-career">
            <Career />
          </div>
          <div class="sub-about-project">
            <Project />
          </div>
          <div class="sub-about-skills">
            <Skills />
          </div>
          <div class="sub-about-keywords">
            <Knowledges />
          </div>
        </div>
      </div>
    </div>
  </section>
</main>

<Footer />
