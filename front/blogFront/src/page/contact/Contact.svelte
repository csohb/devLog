<script lang="ts">
import { onMount } from "svelte";
import { fetchEmailSend } from "../../api/contact";
import type { ContactRequestBody } from "../../api/types/contact";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import popupStore from "../../stores/popup";

let receiver = "";
const RECEIVERS = [
  { val: "yeong", text: "yeongJi" },
  { val: "yujin", text: "yujin" },
];

let contents = "";
let textCount = 0;

let sendName = "";
let errorName = false;
let sendEmail = "";
let errorEmail = false;
let sendTell = "";
let errorTell = false;
let sendTitle = "";

onMount(() => {
  init();
});

function init() {
  // 처음 select box name 지정
  let rand = Math.floor(Math.random() * RECEIVERS.length);
  let rValue = RECEIVERS[rand];
  receiver = rValue.val;
}

function changeReceiver() {}

function onKeyupTextKeyword() {
  textCount = contents.length;
}

function validateCheckName(name: string) {
  if (name.trim() === "") {
    errorName = false;
    return;
  }
  if (name) {
    errorName = !name.match(/^[가-힣a-zA-Z]{2,}$/);
  }
}

function validateCheckEmail(email: string) {
  // 이메일 정규식 체크
  if (email.trim() === "") {
    errorEmail = false;
    return;
  }
  if (email) {
    errorEmail = !email.match(
      /^[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*@[0-9a-zA-Z]([-_.]?[0-9a-zA-Z])*.[a-zA-Z]{2,3}$/i
    );
  }
}

function validateCheckTell(tell: string) {
  // 전화번호 정규식
  if (tell.trim() === "") {
    errorTell = false;
    return;
  }
  if (tell) {
    errorTell = !tell.match(/^\d{3}-\d{3,4}-\d{4}$/);
  }
}

let nameEl: HTMLInputElement;
let emailEl: HTMLInputElement;
let tellEl: HTMLInputElement;
let titleEl: HTMLInputElement;
let contentsEl: HTMLTextAreaElement;

async function onClickSend() {
  if (errorName || sendName.trim() === "") {
    popupStore.open({
      title: "CONTACT",
      text: "내용이 올바르지 않습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        nameEl.focus();
      },
    });
    return;
  }
  if (errorEmail || sendEmail.trim() === "") {
    popupStore.open({
      title: "CONTACT",
      text: "내용이 올바르지 않습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        emailEl.focus();
      },
    });
    return;
  }
  if (errorTell || sendTell.trim() === "") {
    popupStore.open({
      title: "CONTACT",
      text: "내용이 올바르지 않습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        tellEl.focus();
      },
    });
    return;
  }
  if (sendTitle.trim() === "") {
    popupStore.open({
      title: "CONTACT",
      text: "제목이 없습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        titleEl.focus();
      },
    });
    return;
  }
  if (contents.trim() === "" || textCount > 500) {
    popupStore.open({
      title: "CONTACT",
      text: "내용이 올바르지 않습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        contentsEl.focus();
      },
    });
    return;
  }
  let request: ContactRequestBody = {
    receiver,
    name: sendName,
    email: sendEmail,
    number: sendTell,
    title: sendTitle,
    content: contents,
  };

  console.log("request:", request);
  await fetchEmailSend(request).then((resp) => {
    console.log("email resp:", resp);
  });
}

$: validateCheckName(sendName);
$: validateCheckEmail(sendEmail);
$: validateCheckTell(sendTell);
</script>

<Header />

<main class="sub-page">
  <section class="sub-contact-wrap">
    <div class="inner">
      <div class="sub-contact">
        <h1 class="sub-page-title">Contact</h1>
        <div>
          <select
            class="sub-contact-select"
            bind:value="{receiver}"
            on:change="{changeReceiver}">
            {#each RECEIVERS as item}
              <option value="{item.val}">
                {item.text}
              </option>
            {/each}
          </select>
          <i class="sub-contact-select-arrow down" role="img"></i>
        </div>
        <div class="sub-contact-contents">
          <div class="sub-contact-title">
            <p>ㅁㅁㅁ 문의</p>
            <p>문의사항에 대한 답변은 일주일 안에 회신드리겠습니다.</p>
          </div>
          <div class="sub-contact-form">
            <input
              type="text"
              name="name"
              maxlength="10"
              placeholder="이름"
              bind:this="{nameEl}"
              bind:value="{sendName}"
              class="{errorName ? 'warning' : ''}" />
            {#if errorName}
              <p class="warning">한글과 영문만 입력 가능합니다.</p>
            {/if}
            <input
              type="email"
              name="email"
              placeholder="이메일"
              bind:this="{emailEl}"
              bind:value="{sendEmail}"
              class="{errorEmail ? 'warning' : ''}" />
            {#if errorEmail}
              <p class="warning">이메일 형식이 올바르지 않습니다.</p>
            {/if}
            <input
              type="tel"
              name="tell"
              placeholder="전화번호( 010-0000-0000 )"
              bind:this="{tellEl}"
              bind:value="{sendTell}"
              class="{errorTell ? 'warning' : ''}" />
            {#if errorTell}
              <p class="warning">전화번호 형식이 올바르지 않습니다.</p>
            {/if}
            <input
              type="text"
              name="title"
              placeholder="제목"
              bind:this="{titleEl}"
              bind:value="{sendTitle}" />
            <textarea
              name="contents"
              cols="30"
              rows="10"
              placeholder="문의내용"
              bind:this="{contentsEl}"
              bind:value="{contents}"
              on:keyup="{onKeyupTextKeyword}"></textarea>
            <span>{textCount}자 / 최대 500자</span>
          </div>
          <div class="sub-contact-action">
            <button type="submit" on:click="{onClickSend}">제출하기</button>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>

<Footer />
