<script lang="ts">
import { onMount } from "svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";

let receiver = "";
const RECEIVERS = [
  { val: "yeong", text: "김영지" },
  { val: "yujin", text: "권유진" },
];

onMount(() => {
  init();
});

function init() {
  // 처음 select box name 지정
  let rand = Math.floor(Math.random() * RECEIVERS.length);
  let rValue = RECEIVERS[rand];
  receiver = rValue.val;
}

function changeReceiver() {}
</script>

<Header />

<main class="sub-page">
  <section class="sub-contact-wrap">
    <div class="inner">
      <div class="sub-contact">
        <h1 class="sub-page-title">Contact</h1>
        <select
          class="sub-contact-select"
          bind:value="{receiver}"
          on:change="{changeReceiver}">
          {#each RECEIVERS as item}
            <option value="{item.val}">
              {item.text}
            </option>
          {/each}
        </select>
        <i class="sub-contact-select-arrow down" role="img"></i>
      </div>
    </div>
  </section>
</main>

<Footer />
