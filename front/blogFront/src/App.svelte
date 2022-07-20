<script lang="ts">
import Router from "svelte-spa-router";
import Popup from "./components/common/Popup.svelte";
import routes from "./router";
</script>

<div class="wrap">
  <Router routes="{routes}" />
  <Popup />
</div>
