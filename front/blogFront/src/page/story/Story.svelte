<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import List from "../../components/story/List.svelte";
import { link } from "svelte-spa-router";
</script>

<Header />
<main class="sub-page">
  <section class="sub-story-wrap">
    <div class="inner">
      <div class="sub-story">
        <h1 class="sub-page-title">Story</h1>
        <div class="sub-story-contents">
          <List />
        </div>
        <a href="/story/edit/register" use:link>등록하기</a>
      </div>
    </div>
  </section>
</main>
<Footer />
