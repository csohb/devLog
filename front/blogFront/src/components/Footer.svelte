<script lang="ts">
import { onMount } from "svelte";
import { link } from "svelte-spa-router";
import authStore from "../stores/auth";

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

function onClickLogout() {
  authStore.setNick("");
  deleteCookie();
}

function deleteCookie() {
  document.cookie = "user_id=;expires=" + new Date(0).toUTCString();
}
</script>

<footer class="footer">
  <div class="inner">
    <div class="footer-wrap">
      <h2>DEVLOG</h2>
      {#if $authStore.loginNick === ""}
        <a href="/login" use:link>LOGIN</a>
      {:else}
        <a href="#none" on:click|preventDefault="{onClickLogout}">LOGOUT</a>
      {/if}
    </div>
  </div>
</footer>
