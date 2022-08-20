<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import List from "../../components/story/List.svelte";
import { link } from "svelte-spa-router";
import authStore from "../../stores/auth";
import { onMount } from "svelte";

onMount(() => {
  getCookie();
});

function getCookie() {
  if ($authStore.loginNick === "") {
    let loginNick = authStore.getCookie("user_id");
    if (loginNick != null) {
      authStore.setNick(loginNick);
    }
  }
}
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
        {#if $authStore.loginNick !== ""}
          <a href="/story/edit/register" use:link>등록하기</a>
        {/if}
      </div>
    </div>
  </section>
</main>
<Footer />
