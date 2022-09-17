<script>
/* 에디터 라이브러리 때문에 타입스크립트를 적용할 수 없음 @ */
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { push } from "svelte-spa-router";
import Editor from "cl-editor/src/Editor.svelte";
import { onDestroy, onMount, tick } from "svelte";
import { fetchBlogSave, fetchBlogUpdate } from "../../api/blog";
import blogStore from "../../stores/blog";
import popupStore from "../../stores/popup";
import authStore from "../../stores/auth";

let title = "";
let tag = "";
let description = "";
let contents = "<p>내용을 입력해주세요.</p>";
let isSave = false;
let tagList = [];
let writer = "";
let date = "";

export let params = {
  id: "",
};

onMount(() => {
  // 만약 수정이면 내용 불러오기
  if (params.id !== "register") {
    init();
    return;
  }
  if ($authStore.loginNick === "") {
    let loginNick = authStore.getCookie("user_id");
    if (loginNick == null) {
      push("/login");
    } else {
      authStore.setNick(loginNick);
    }
    return;
  }
  // writer 설정 필요
  writer = $authStore.loginNick;
  setToday();
});

async function init() {
  // 수정하기 전 정보값 가져오기
  await blogStore.setBlogDetail(params.id);
  await tick();

  // tagList = $blogStore.blogDetail.tags;
  title = $blogStore.blogDetail.title;
  description = $blogStore.blogDetail.description;
  writer = $blogStore.blogDetail.writer;
  date = $blogStore.blogDetail.date;
}

function setToday() {
  // 오늘 날짜 계산
  const today = new Date();
  date = `${today.getFullYear()}년 ${
    today.getMonth() + 1
  }월 ${today.getDate()}일`;
}

function onClickTagAdd() {
  // 태그 리스트 추가하기
  if (tag.trim() === "") {
    return;
  }
  const idx = tagList.indexOf(tag);
  if (idx === -1) {
    tagList.push(tag);
    tag = "";
  } else {
    alert("등록한 태그 입니다.");
  }
  tagList = tagList;
}

async function onClickUpdate() {
  // 블로그 내용 수정 저장
  if (params.id === "register") {
    return;
  }
  let req = {
    id: params.id,
  };
  if (title.trim() !== "" && title !== $blogStore.blogDetail.title) {
    Object.assign(req, {
      title: title,
    });
  }
  if (
    contents !== "<p>내용을 입력해주세요.</p>" &&
    contents !== $blogStore.blogDetail.content
  ) {
    Object.assign(req, {
      content: contents,
    });
  }
  if (tagList.length !== 0) {
    Object.assign(req, {
      tags: tagList,
    });
  }
  if (description !== $blogStore.blogDetail.description) {
    Object.assign(req, {
      description,
    });
  }

  if (req.title == null) {
    Object.assign(req, {
      title: $blogStore.blogDetail.title,
    });
  }
  if (req.content == null) {
    Object.assign(req, {
      content: $blogStore.blogDetail.content,
    });
  }
  if (req.tags == null) {
    Object.assign(req, {
      tags: $blogStore.blogDetail.tags,
    });
  }
  if (req.description == null) {
    Object.assign(req, {
      description: $blogStore.blogDetail.description,
    });
  }
  await fetchBlogUpdate(req).then(() => {
    push(`/blog/${params.id}`);
  });
}

async function onClickSave() {
  // 초기 저장
  if (isSave) {
    alert("이미 저장 되었습니다.");
    return;
  }

  if (tagList.length === 0) {
    alert("태그를 등록해주세요.");
    return;
  }

  if (title.trim() === "") {
    alert("제목을 등록해주세요.");
    return;
  }

  if (contents.trim() === "") {
    alert("내용이 없습니다.");
    return;
  }

  // writer 를 세션에서 가져와야 할 것 같은데,,
  await fetchBlogSave({
    title,
    content: contents,
    writer,
    description,
    tags: tagList,
  })
    .then(() => {
      alert("저장 완료!");
      isSave = true;
    })
    .catch((err) => {
      console.log("blog save err :", err);
    });
}

function resetValue() {
  // 변수값 초기화
  tagList = [];
  contents = "<p>내용을 입력해주세요.</p>";
  title = "";
  tag = "";
  description = "";
  isSave = false;
  date = "";
}

function onClickCancel() {
  let btn = "";
  if (params.id === "register") {
    btn = "목록으로 이동하기";
  } else {
    btn = "페이지로 돌아가기";
  }
  // 수정 취소
  popupStore.open({
    title: "BLOG",
    text: "취소하시겠습니까?<br />취소하면 입력한 데이터가 사라집니다.",
    btn,
    type: "confirm",
    isShow: false,
    action: () => {
      resetValue();
      if (params.id === "register") {
        push("/blog");
      } else {
        push(`/blog/${params.id}`);
      }
    },
  });
}

onDestroy(() => {
  resetValue();
});
</script>

<Header />
<main class="sub-page">
  <section class="sub-blog-wrap">
    <div class="inner">
      <div class="sub-blog">
        <h1 class="sub-page-title">Blog</h1>
        <div class="sub-blog-detail-title">
          <span
            >#{tagList}
            {#if $blogStore.blogDetail.tags.length !== 0}
              {`[수정 전: ${$blogStore.blogDetail.tags} ]`}
            {/if}
            <div>
              <input
                type="text"
                placeholder="태그"
                maxlength="15"
                bind:value="{tag}" />
              <button type="button" on:click="{onClickTagAdd}">추가</button>
            </div>
          </span>
          <h2><input type="text" placeholder="제목" bind:value="{title}" /></h2>
          <p>
            <textarea
              name=""
              bind:value="{description}"
              placeholder="블로그에 대한 간단한 내용을 적어주세요."
              style="resize: none;"></textarea>
          </p>
          <div class="sub-blog-detail-info">
            <span>{date}</span>
            <span>|</span>
            <span>{writer}</span>
          </div>
        </div>
        <div class="sub-blog-detail-contents">
          <Editor
            on:change="{({ detail }) => (contents = detail)}"
            html="{contents}"
            height="{'200px'}" />
        </div>
        <div class="sub-blog-detail-btn">
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
