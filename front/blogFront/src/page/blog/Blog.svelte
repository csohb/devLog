<script lang="ts">
import List from "../../components/blog/List.svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link } from "svelte-spa-router";
</script>

<Header />
<main class="sub-page">
  <section class="sub-blog-wrap">
    <div class="inner">
      <div class="sub-blog">
        <h1 class="sub-page-title">Blog</h1>
        <ul class="sub-blog-tag">
          {#each Array(5) as _, index}
            <li>
              <a href="#none"># svelte</a>
            </li>
          {/each}
        </ul>
        <div class="sub-blog-contents">
          <List />
        </div>
        <a href="/blog/edit/register" use:link>등록하기</a>
      </div>
    </div>
  </section>
</main>
<Footer />
