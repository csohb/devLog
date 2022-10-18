<script lang="ts">
import { beforeUpdate, onDestroy, onMount } from "svelte";
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

function init() {
  // 처음 tab name 지정
  if ($querystring === "name=yeong") {
    currentTab = "yeong";
    return;
  }

  if ($querystring === "name=yujin") {
    currentTab = "yujin";
    return;
  }

  let rand = Math.floor(Math.random() * TABLIST.length);
  let rValue = TABLIST[rand];
  currentTab = rValue.name;
}

beforeUpdate(() => {
  if (currentTab !== "") {
    aboutStore.setIntroduce(currentTab);
  }
});

onDestroy(() => {
  aboutStore.resetAbout();
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
          <Introduce bind:isEditMode currentTab="{currentTab}" />
          <div class="sub-about-career">
            <Career isEditMode="{isEditMode}" currentTab="{currentTab}" />
          </div>
          <div class="sub-about-project">
            <Project isEditMode="{isEditMode}" currentTab="{currentTab}" />
          </div>
          <div class="sub-about-skills">
            <Skills />
          </div>
          <div class="sub-about-keywords">
            <Knowledges isEditMode="{isEditMode}" currentTab="{currentTab}" />
          </div>
        </div>
        {#if $authStore.loginNick !== "" && $authStore.loginNick === currentTab}
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
