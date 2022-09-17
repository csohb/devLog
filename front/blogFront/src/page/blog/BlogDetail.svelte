<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link, push } from "svelte-spa-router";
import { onDestroy, onMount } from "svelte";
import blogStore from "../../stores/blog";
import popupStore from "../../stores/popup";
import authStore from "../../stores/auth";
import { fetchBlogDelete } from "../../api/blog";
import { unlike, like } from "../../icon/Icon";

export let params = {
  id: "",
};

onMount(() => {
  if (!params.id) {
    push("/blog");
    return;
  }
  blogStore.setBlogDetail(params.id);
  getCookie();
});

function setTags(tagList: string[]): any {
  // 태그 리스트에 # 추가하기
  return tagList.reduce((temp: any, item: string) => {
    temp.push(`# ${item}`);
    return temp;
  }, []);
}

function onClickDelete() {
  // 삭제 여부를 묻는 알림 팝업 필요
  if (!params.id) {
    return;
  }
  popupStore.open({
    title: "BLOG",
    text: "삭제하시겠습니까?",
    btn: "삭제하기",
    type: "confirm",
    isShow: false,
    action: async () => {
      await fetchBlogDelete(params.id)
        .then((resp) => {
          push("/blog");
        })
        .catch((err) => {
          console.log("blog delete err:", err);
        });
    },
  });
}

let heartClick = false;
function onClickHeart(isAdd: boolean) {
  if (!params.id) {
    return;
  }
  if (heartClick) {
    return;
  }
  heartClick = true;
  blogStore.heartCount(params.id, isAdd).finally(() => {
    heartClick = false;
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
  blogStore.resetBlogDetail();
});
</script>

<Header />
<main class="sub-page">
  <section class="sub-blog-wrap">
    <div class="inner">
      <div class="sub-blog">
        <h1 class="sub-page-title">Blog</h1>
        {#if !!$blogStore.blogDetail.id}
          <div class="sub-blog-detail-title">
            <span>
              {#each setTags($blogStore.blogDetail.tags) as item}
                {item}&nbsp;&nbsp;&nbsp;&nbsp;
              {/each}
            </span>
            <h2>{$blogStore.blogDetail.title}</h2>
            <p>
              {$blogStore.blogDetail.description}
            </p>
            <div class="sub-blog-detail-info">
              <span>{$blogStore.blogDetail.date}</span>
              <span>|</span>
              <span>{$blogStore.blogDetail.writer}</span>
              <span>|</span>
              <span>{$blogStore.blogDetail.view}</span>
              <span>|</span>
              <span class="sub-blog-detail-heart">
                <div
                  on:click="{() => {
                    onClickHeart(true);
                  }}">
                  {@html like}+
                </div>
                <div
                  on:click="{() => {
                    onClickHeart(false);
                  }}">
                  {@html unlike}-
                </div>
                {$blogStore.blogDetail.heart}</span>
            </div>
          </div>
          <div class="sub-blog-detail-contents">
            {@html $blogStore.blogDetail.content}
          </div>
        {/if}
        <div class="sub-blog-detail-btn">
          <a href="/blog" use:link>목록으로 돌아가기</a>
          {#if $authStore.loginNick !== "" && $blogStore.blogDetail.writer === $authStore.loginNick}
            <a href="/blog/edit/{params.id}" use:link>수정하기</a>
            <button type="button" on:click="{onClickDelete}">삭제</button>
          {/if}
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
