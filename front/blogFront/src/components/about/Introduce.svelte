<script lang="ts">
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
  }
}

$: editModeHandler(isEditMode);

function editHandler() {
  if (validation() && currentTab.trim() !== "") {
    aboutStore.updateIntroduce(currentTab, intro, address, mail).then(() => {
      popupStore.open({
        title: "프로필",
        text: "수정되었습니다.",
        type: "alert",
        isShow: false,
        action: () => {
          isEditMode = false;
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
  return true;
}
</script>

{#if isEditMode}
  <div class="sub-about-introduce">
    <span>소개말 :</span>
    <textarea bind:value="{intro}"></textarea>
    <button type="button" on:click="{editHandler}">프로필 내용 수정버튼</button>
  </div>
{:else}
  <div class="sub-about-introduce">
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
