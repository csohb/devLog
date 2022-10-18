<script lang="ts">
import { imgUpload } from "../../api/contact";
import aboutStore from "../../stores/about";
import popupStore from "../../stores/popup";

export let isEditMode: boolean = false;
export let currentTab = "";

// 기본정보
let intro = "";
let address = "";
let mail = "";

function editModeHandler(edit: boolean) {
  if (edit) {
    intro = $aboutStore.description;
    address = $aboutStore.info.addr;
    mail = $aboutStore.info.email;
    if ($aboutStore.info.img !== "") {
      introImageUrl = $aboutStore.info.img;
    }
    if (localStorage.getItem("s3_img_url")) {
      let storageImg = localStorage.getItem("s3_img_url");

      introImageUrl = JSON.parse(storageImg);
    }
  }
}

$: editModeHandler(isEditMode);

function editHandler() {
  if (validation() && currentTab.trim() !== "") {
    aboutStore
      .updateIntroduce(currentTab, intro, address, mail, introImageUrl)
      .then(() => {
        popupStore.open({
          title: "프로필",
          text: "수정되었습니다.",
          type: "alert",
          isShow: false,
          action: () => {
            isEditMode = false;
            localStorage.removeItem("s3_img_url");
          },
        });
      });
  }
}

function validation(): boolean {
  if (intro.trim() === "") {
    return false;
  }
  if (address.trim() === "") {
    return false;
  }
  if (mail.trim() === "") {
    return false;
  }
  if (introImageUrl.trim() === "") {
    return false;
  }
  return true;
}
let fileName = "";
let fileinputEl: HTMLInputElement;
let files: FileList;
let introImageUrl = "";

function onChangeFile(file: FileList) {
  if (file === undefined) {
    return;
  }
  if (file.length === 0) {
    return;
  }

  // 이미지 한개만 등록가능..
  if (fileName.trim() !== "") {
    console.log("이미지 한개만 등록 가능..");
    return;
  }

  fileName = file[0].name;

  let fileInfo = {
    pic_name: fileName,
    dir_name: "introduce",
  };

  let formData = new FormData();
  formData.append("filename", file[0]);
  formData.append(
    "file_info",
    new Blob([JSON.stringify(fileInfo)], { type: "application/json" })
  );

  imgUpload(formData).then((resp: { file_url: string }) => {
    popupStore.open({
      title: "프로필",
      text: "프로필 이미지가 업로드 되었습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        introImageUrl = resp.file_url;
        localStorage.setItem("s3_img_url", JSON.stringify(introImageUrl));
        introImageUrl = introImageUrl;
        const imgEl = document.getElementById("introImg") as HTMLInputElement;
        imgEl.value = "";
        fileName = "";
      },
    });
  });
}

$: onChangeFile(files);
</script>

{#if isEditMode}
  <div class="sub-about-introduce">
    <span>소개말 :</span>
    <textarea bind:value="{intro}"></textarea>
    <span>프로필 이미지 :</span>
    <div class="filebox">
      <input
        class="upload-name"
        value="{fileName}"
        placeholder="이미지만 첨부 가능"
        disabled />

      <label for="introImg">첨부파일</label>
      <input type="file" id="introImg" bind:files bind:this="{fileinputEl}" />
      <span>등록할 이미지 : {introImageUrl}</span>
    </div>
    <button type="button" on:click="{editHandler}">프로필 내용 수정버튼</button>
  </div>
{:else}
  <div class="sub-about-introduce">
    {#if $aboutStore.info.img !== ""}
      <img src="{$aboutStore.info.img}" alt="" />
    {/if}
    <pre>{#if $aboutStore.description === ""}description 등록 중 입니다.{:else}
        {$aboutStore.description}{/if}</pre>
  </div>
{/if}

<div class="sub-about-info">
  <h2>Info</h2>
  <ul>
    <li>
      <strong>Name</strong>{$aboutStore.info.name} ( {$aboutStore.info
        .nick_name}
      )
    </li>
    <li><strong>Date of birth</strong>{$aboutStore.info.birthday}</li>
    {#if isEditMode}
      <li>
        <strong>주소 :</strong>
        <input type="text" bind:value="{address}" />
      </li>
      <li>
        <strong>이메일 :</strong>
        <input type="text" bind:value="{mail}" />
      </li>
    {:else}
      <li><strong>Address</strong>{$aboutStore.info.addr}</li>
      <li><strong>E-mail</strong>{$aboutStore.info.email}</li>
    {/if}
  </ul>
</div>
