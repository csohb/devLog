<script lang="ts">
import Router from "svelte-spa-router";
import Popup from "./components/common/Popup.svelte";
import routes from "./router";

// console.log("url:", process.env.NODE_ENV);
// console.log("prod:", process.env.isProd);
// console.log("dd", process.env.BASE_URL);
</script>

<div class="wrap">
  <Router routes="{routes}" />
  <Popup />
</div>
