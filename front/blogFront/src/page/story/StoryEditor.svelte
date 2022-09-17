<script>
/* 에디터 라이브러리 때문에 타입스크립트를 적용할 수 없음 @ */
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link, push } from "svelte-spa-router";
import Editor from "cl-editor/src/Editor.svelte";
import storyStore from "../../stores/story";
import authStore from "../../stores/auth";
import { onMount } from "svelte";

let writer = "";
let title = "";
let contents = "<p>내용을 입력해주세요.</p>";

export let params = {
  id: "",
};

onMount(() => {
  if ($authStore.loginNick === "") {
    // 만약 닉네임이 없으면
    let loginNick = authStore.getCookie("user_id");
    if (loginNick == null) {
      // 쿠키에 없으면 로그인 페이지로 이동
      push("/login");
      return;
    }
    authStore.setNick(loginNick);
  }
  // 진입 파라미터가 수정이면 init
  if (params.id !== "register") {
    init();
    return;
  }
  // 신규 등록일 경우

  // 작성자 필요
  writer = $authStore.loginNick;
  // 현재 날짜 필요
  setToday();
});

const positions = [
  { val: "Backend", text: "Backend" },
  { val: "Frontend", text: "Frontend" },
];
let positionSelected = "Backend";

function changePosition() {
  // backend, frontend 선택
  console.log("change:::", positionSelected);
}

function init() {
  // 수정하기 전에 정보값 불러오기
  storyStore.setStoryDetail(params.id);
}

function setToday() {
  // 신규 등록일 설정
}

function onClickSave() {}

function onClickUpdate() {}

function onClickCancel() {
  if (params.id === "register") {
    push("/story");
  } else {
    push(`/story/${params.id}`);
  }
}
</script>

<Header />
<main class="sub-page">
  <section class="sub-story-wrap">
    <div class="inner">
      <div class="sub-story">
        <h1 class="sub-page-title">Story</h1>
        <div class="sub-story-detail-title">
          <span>
            <!-- backend / frontend 선택 -->
            <select
              bind:value="{positionSelected}"
              on:change="{changePosition}">
              {#each positions as item}
                <option value="{item.val}">
                  {item.text}
                </option>
              {/each}
            </select>
          </span>
          <h2><input type="text" placeholder="제목" bind:value="{title}" /></h2>
          <p>
            <Editor
              on:change="{({ detail }) => (contents = detail)}"
              html="{contents}"
              height="{'200px'}" />
          </p>
          <div class="sub-story-detail-thumb">
            <!-- 이미지 업로드 -->
            <img
              src="https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg"
              alt="" />
          </div>
          <div class="sub-story-detail-info">
            <span>2022-07-24</span>
            <span>|</span>
            <span>yeong</span>
          </div>
        </div>
        <div class="sub-story-detail-contents">
          {#if params.id === "register"}등록{:else}수정{/if}
        </div>
        <div class="sub-story-detail-btn">
          {#if params.id === "register"}
            <a href="#none" on:click|preventDefault="{onClickSave}">등록</a>
          {:else}
            <a href="#none" on:click|preventDefault="{onClickUpdate}">수정</a>
          {/if}

          <a href="#none" on:click|preventDefault="{onClickCancel}">취소</a>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
