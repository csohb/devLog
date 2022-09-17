<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link, push } from "svelte-spa-router";
import { unlike, like } from "../../icon/Icon";
import authStore from "../../stores/auth";
import { onDestroy, onMount } from "svelte";
import storyStore from "../../stores/story";
import popupStore from "../../stores/popup";
import { fetchStoryDelete } from "../../api/story";

let isLike = false;
let likeCount = 0;

export let params = {
  id: "",
};

onMount(() => {
  if (!params.id) {
    push("/story");
    return;
  }
  getCookie();
  storyStore.setStoryDetail(params.id);
});

function onClickLike() {
  // 좋아요 기능 ( 제거 )
  isLike = !isLike;
  if (isLike) {
    likeCount += 1;
  } else {
    likeCount -= 1;
  }
}

function onClickDelete() {
  if (!params.id) {
    return;
  }
  popupStore.open({
    title: "STORY",
    text: "삭제하시겠습니까?",
    btn: "삭제하기",
    type: "confirm",
    isShow: false,
    action: async () => {
      await fetchStoryDelete(params.id)
        .then(() => {
          push("/story");
        })
        .catch((err) => {
          console.log("삭제 에러:", err);
        });
    },
  });
}

function getCookie() {
  if ($authStore.loginNick === "") {
    let loginNick = authStore.getCookie("user_id");
    if (loginNick != null) {
      authStore.setNick(loginNick);
    }
  }
}

onDestroy(() => {
  storyStore.resetStoryDetail();
});
</script>

<Header />
<main class="sub-page">
  <section class="sub-story-wrap">
    <div class="inner">
      <div class="sub-story">
        <h1 class="sub-page-title">Story</h1>
        <div class="sub-story-detail-title">
          <!-- 선택하는 셀렉트 박스 필요 -->
          <span> Backend / Frontend </span>
          <h2>{$storyStore.storyDetail.title}</h2>
          <p>{$storyStore.storyDetail.description}</p>
          <div class="sub-story-detail-info">
            <span>{$storyStore.storyDetail.created_at}</span>
            <span>|</span>
            <span>{$storyStore.storyDetail.writer}</span>
            <span>|</span>
            <span>{$storyStore.storyDetail.view}</span>
          </div>
          <div class="sub-story-detail-thumb">
            <img
              src="https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg"
              alt="" />
          </div>
        </div>
        <div class="sub-story-detail-contents">
          {@html $storyStore.storyDetail.content}
        </div>
        <!-- 좋아요 기능 제거 -->
        <!-- <div class="sub-story-detail-action">
          <div class="sub-story-liked" on:click="{onClickLike}">
            {#if isLike}
              {@html like}
            {:else}
              {@html unlike}
            {/if}
            {likeCount}
          </div>
        </div> -->
        <div class="sub-story-detail-btn">
          <a href="/story" use:link>목록으로 돌아가기</a>
          {#if $authStore.loginNick !== ""}
            <a href="/story/edit/{params.id}" use:link>수정하기</a>
            <button type="button" on:click="{onClickDelete}">삭제</button>
          {/if}
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
