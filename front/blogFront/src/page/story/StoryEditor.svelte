<script>
/* 에디터 라이브러리 때문에 타입스크립트를 적용할 수 없음 @ */
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { link, push } from "svelte-spa-router";
import { fetchStoryUpdate, fetchStorySave } from "../../api/story";
import Editor from "cl-editor/src/Editor.svelte";
import storyStore from "../../stores/story";
import authStore from "../../stores/auth";
import popupStore from "../../stores/popup";
import { onMount, tick } from "svelte";
import FilePond, { registerPlugin } from "svelte-filepond";
import FilePondPluginImageExifOrientation from "filepond-plugin-image-exif-orientation";
import FilePondPluginImagePreview from "filepond-plugin-image-preview";
import "filepond/dist/filepond.css";
import "filepond-plugin-image-preview/dist/filepond-plugin-image-preview.css";
import { imgUpload } from "../../api/contact";

registerPlugin(FilePondPluginImageExifOrientation, FilePondPluginImagePreview);

// test
const positions = [
  { val: "B", text: "Backend" },
  { val: "F", text: "Frontend" },
];
let positionSelected = "B";
let writer = "";
let title = "";
let date = "";
let contents = "";
let description = "";

let fileinputEl;
let files;
let fileName = "";
let src = "";

let imageUrlArr = [];

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
  if (localStorage.getItem("s3_img_url")) {
    let storageImg = localStorage.getItem("s3_img_url");

    imageUrlArr = JSON.parse(storageImg);
  }
  console.log("imageUrlArr:", imageUrlArr);
  if (params.id !== "register") {
    // 진입 파라미터가 수정이면 init
    init();
    return;
  }
  // 신규 등록일 경우

  // 작성자 필요
  writer = $authStore.loginNick;
  contents = "<p>내용을 입력해주세요.</p>";
  // 현재 날짜 필요
  setToday();
});

function onChangeFile(file) {
  if (file === undefined) {
    return;
  }
  if (file.length === 0) {
    return;
  }
  if (imageUrlArr.length > 10) {
    return;
  }

  console.log("file", file);
  fileName = file[0].name;

  let fileInfo = {
    pic_name: fileName,
    dir_name: "story",
  };

  let formData = new FormData();
  formData.append("filename", file[0]);
  formData.append(
    "file_info",
    new Blob([JSON.stringify(fileInfo)], { type: "application/json" })
  );

  for (var pair of formData.entries()) {
    console.log(pair[0] + ", " + JSON.stringify(pair[1]));
  }

  imgUpload(formData).then((resp) => {
    console.log("upload:", resp);
    imageUrlArr.push({
      url: resp.file_url,
      name: fileInfo.pic_name,
    });
    localStorage.setItem("s3_img_url", JSON.stringify(imageUrlArr));
    imageUrlArr = imageUrlArr;
    document.getElementById("ex_filename").value = "";
  });
}

$: onChangeFile(files);

function changePosition() {
  // backend, frontend 선택
  console.log("change:::", positionSelected);
}

async function init() {
  // 수정하기 전에 정보값 불러오기
  await storyStore.setStoryDetail(params.id);
  await tick();
  title = $storyStore.storyDetail.title;
  writer = $storyStore.storyDetail.writer;
  date = $storyStore.storyDetail.created_at;
  description = $storyStore.storyDetail.description;
  contents = `${$storyStore.storyDetail.content}`;
  if ($storyStore.storyDetail.images !== null) {
    const temp = $storyStore.storyDetail.images;
    temp.map((val) => {
      imageUrlArr.push({
        url: val,
        name: "등록된 이미지 불러오기",
      });
    });
    console.log(imageUrlArr);
    imageUrlArr = imageUrlArr;
  }
}

function setToday() {
  // 신규 등록일 설정
  const today = new Date();
  date = `${today.getFullYear()}년 ${
    today.getMonth() + 1
  }월 ${today.getDate()}일`;
}

async function onClickSave() {
  // 신규 등록
  if (params.id !== "register") {
    return;
  }
  if (title.trim() === "") {
    alert("제목을 등록해주세요.");
    return;
  }

  if (contents.trim() === "" || contents === "<p><br></p>") {
    alert("내용이 없습니다.");
    return;
  }

  let images = imageUrlArr.map((val) => {
    return val.url;
  });

  console.log("images:", images);

  let request = {
    type: positionSelected,
    title,
    content: contents,
    description,
    images: images,
  };

  await fetchStorySave(request).then((resp) => {
    console.log("등록 완료:::", resp);
    localStorage.removeItem("s3_img_url");
    push("/story");
  });
}

async function onClickUpdate() {
  // 수정
  if (params.id === "register") {
    return;
  }
  let request = {
    id: params.id,
    title: title,
    content: contents,
    description,
  };

  await fetchStoryUpdate(request)
    .then((resp) => {
      resetValue();
      init();
      popupStore.open({
        title: "STORY",
        text: "수정이 완료되었습니다<br />목록으로 이동하시겠습니까?",
        btn: "목록으로 이동하기",
        type: "confirm",
        isShow: false,
        action: () => {
          push(`/story/${params.id}`);
        },
      });
    })
    .catch(() => {
      push(`/story/${params.id}`);
    });
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
    title: "STORY",
    text: "취소하시겠습니까?<br />취소하면 입력한 데이터가 사라집니다.",
    btn,
    type: "confirm",
    isShow: false,
    action: () => {
      resetValue();
      if (params.id === "register") {
        // 신규 등록
        localStorage.removeItem("s3_img_url");
        push("/story");
        return;
      }
      push(`/story/${params.id}`);
    },
  });
}

function resetValue() {
  // 변수값 초기화
  contents = "";
  title = "";
  date = "";
  writer = "";
  positionSelected = "Backend";
}

/* 이미지 업로드 */
function handleInit() {
  // FilePond 인스턴스가 생성되어 준비되었습니다.
  console.log("FilePond has initialised");
}

function handleAddFile(err, fileItem) {
  // 오류가 없으면 파일이 성공적으로 로드된 것입니다.
  console.log("A file has been added", fileItem);
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
            홈 화면에서만 보일 간단한 글
            <textarea
              type="text"
              placeholder="description"
              bind:value="{description}"></textarea>
          </p>
          <p>
            {#if contents !== ""}
              <Editor
                on:change="{({ detail }) => (contents = detail)}"
                html="{contents}"
                height="{'200px'}" />
            {/if}
          </p>
          <div class="sub-story-detail-thumb">
            <!-- 이미지 업로드   server="/api" -->
            <p>이미지 등록 [ 첫번째가 메인 이미지 ]</p>
            <!-- 
            <FilePond
              name="메인이미지"
              server="/api/v1/upload"
              instantUpload="{false}"
              allowMultiple="{true}"
              oninit="{handleInit}"
              onaddfile="{handleAddFile}" />
            -->
            <!-- <img
              src="https://ridicorp.com/wp-content/uploads/2022/06/sf_1_thumb@0.5x.jpg"
              alt="" /> -->
            {#if params.id === "register"}
              <div class="filebox">
                <input
                  class="upload-name"
                  value="{fileName}"
                  placeholder="txt, pdf 첨부 가능"
                  disabled />

                <label for="ex_filename">첨부파일</label>
                <input
                  type="file"
                  id="ex_filename"
                  bind:this="{fileinputEl}"
                  bind:files />
                <!--    accept="text/plain,.pdf" -->
              </div>
            {/if}
            <div class="upload-images">
              <ul>
                {#each imageUrlArr as img, index}
                  <li>
                    <span>
                      <strong style="color:#551a8b;"
                        >{index + 1} {img.name} :
                      </strong>
                      {img.url}
                    </span>
                    <img class="thumb" src="{img.url}" alt="{img.name}" />
                  </li>
                  <br />
                {/each}
              </ul>
            </div>
          </div>
          <div class="sub-story-detail-info">
            <span>{date}</span>
            <span>|</span>
            <span>{writer}</span>
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
