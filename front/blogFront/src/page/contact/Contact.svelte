<script lang="ts">
import { onMount } from "svelte";
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";

let receiver = "";
const RECEIVERS = [
  { val: "yeong", text: "김영지" },
  { val: "yujin", text: "권유진" },
];
let fileinputEl: HTMLInputElement;
let files: any;
let fileName = "";

let contents = "";
let textCount = 0;

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

function onChangeFile(file: any) {
  if (file === undefined) {
    return;
  }

  fileName = file[0].name;
}

$: onChangeFile(files);
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
            <input type="text" name="name" maxlength="10" placeholder="이름" />
            <input type="email" name="email" placeholder="이메일" />
            <input type="tel" name="tell" placeholder="전화번호" />
            <input type="text" name="title" placeholder="제목" />
            <textarea
              name="contents"
              cols="30"
              rows="10"
              placeholder="문의내용"
              bind:value="{contents}"
              on:keyup="{onKeyupTextKeyword}"></textarea>
            <span>{textCount}자 / 최대 500자</span>
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
                accept="text/plain,.pdf"
                bind:this="{fileinputEl}"
                bind:files />
            </div>
          </div>
          <div class="sub-contact-action">
            <button type="submit">제출하기</button>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>

<Footer />
