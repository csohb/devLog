<script lang="ts">
import Header from "../components/Header.svelte";
import Footer from "../components/Footer.svelte";
import About from "../components/main/About.svelte";
import Branding from "../components/main/Branding.svelte";
import Story from "../components/main/Story.svelte";
import Blog from "../components/main/Blog.svelte";

let y: number = 0;
let headerH: number = 0;
let screenH: number = 0;
let headerColorClass: boolean = false;

function headerHeightHandler(h: CustomEvent) {
  headerH = h.detail;
}

function navColorHandler(y: number, head: number, screen: number) {
  console.log(headerH);
  if (screen - head < y) {
    headerColorClass = true;
  } else {
    headerColorClass = false;
  }
}

$: navColorHandler(y, headerH, screenH);
</script>

<svelte:window bind:scrollY="{y}" bind:innerHeight="{screenH}" />
<Header on:headerH="{headerHeightHandler}" colorClass="{headerColorClass}" />
<main>
  <section class="main-branding">
    <Branding />
  </section>
  <section class="main-about">
    <About />
  </section>
  <section class="main-story">
    <Story />
  </section>
  <section class="main-blog">
    <Blog />
  </section>
</main>
<Footer />
