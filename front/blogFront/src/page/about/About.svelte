<script lang="ts">
import { beforeUpdate, onMount } from "svelte";
import Career from "../../components/about/Career.svelte";
import Introduce from "../../components/about/Introduce.svelte";
import Knowledges from "../../components/about/Knowledges.svelte";
import Project from "../../components/about/Project.svelte";
import Skills from "../../components/about/Skills.svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import aboutStore from "../../stores/about";
import authStore from "../../stores/auth";
import { querystring, link } from "svelte-spa-router";

let isEditMode = false;

const TABLIST = [
  { nick: "YEONG", name: "yeong" },
  { nick: "CSOHB", name: "yujin" },
];
let currentTab = "";

function onClickTab(type: string) {
  currentTab = type;
  isEditMode = false;
}

onMount(() => {
  init();
});

function init() {
  // 처음 tab name 지정
  if ($querystring === "") {
    let rand = Math.floor(Math.random() * TABLIST.length);
    let rValue = TABLIST[rand];
    currentTab = rValue.name;
  } else if ($querystring === "name=yeong") {
    currentTab = "yeong";
  } else {
    currentTab = "yujin";
  }
}

beforeUpdate(() => {
  if (currentTab !== "") {
    aboutStore.setIntroduce(currentTab);
  }
});
</script>

<Header />
<main class="sub-page">
  <section class="sub-about-wrap">
    <div class="inner">
      <div class="sub-about">
        <h1 class="sub-page-title">About Me</h1>
        <ul class="sub-about-tab">
          {#each TABLIST as person}
            <li
              class:foc="{currentTab === person.name}"
              on:click="{() => {
                onClickTab(person.name);
              }}">
              {person.nick}
            </li>
          {/each}
        </ul>
        <div class="sub-about-contents">
          <Introduce />
          <div class="sub-about-career">
            <Career isEditMode="{isEditMode}" />
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
        {#if $authStore.loginNick !== ""}
          <div class="sub-about-action">
            <a href="/about/edit?writer={currentTab}" use:link>등록</a>
            <button
              type="button"
              on:click="{() => {
                isEditMode = !isEditMode;
              }}">수정</button>
          </div>
        {/if}
      </div>
    </div>
  </section>
</main>

<Footer />
