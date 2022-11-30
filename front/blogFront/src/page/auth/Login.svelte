<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import authStore from "../../stores/auth";
import popupStore from "../../stores/popup";
import { push } from "svelte-spa-router";

let id: string = "";
let pw: string = "";
let inputEnter: boolean = false;

function loginHandler() {
  authStore
    .postLogin(id, pw)
    .then((resp: any) => {
      authStore.setNick(resp.user_id);
      // setCookie(resp.user_id);
      id = "";
      pw = "";
      push("/");
    })
    .catch(() => {
      authStore.setNick("");
      id = "";
      pw = "";
      popupStore.open({
        title: "LOGIN",
        text: "로그인에 실패하였습니다.",
        btn: "확인",
        type: "alert",
        isShow: false,
        action: () => {
          inputEnter = false;
        },
      });
    });
}

function setCookie(value: string) {
  let date = new Date();
  date.setDate(date.getDate() + 1);
  document.cookie =
    encodeURIComponent("user_id") +
    "=" +
    encodeURIComponent(value) +
    (!1 ? "" : "; expires=" + date.toUTCString());
}
</script>

<Header />
<main class="sub-page">
  <section class="sub-login-wrap">
    <div class="inner">
      <div class="login-form">
        <h2>DEVLOG</h2>
        <form action="#">
          <input
            type="text"
            name="id"
            bind:value="{id}"
            placeholder="아이디를 입력해주세요." />
          <input
            type="password"
            name="password"
            bind:value="{pw}"
            on:keypress="{(e) => {
              if (inputEnter) {
                return;
              }
              if (e.code == 'Enter') {
                inputEnter = true;
                loginHandler();
              }
            }}"
            placeholder="비밀번호를 입력해주세요." />
          <button type="button" on:click="{loginHandler}">로그인 하기</button>
        </form>
        <p># 가입은 관리자에게 문의해주세요.</p>
      </div>
    </div>
  </section>
</main>
<Footer />
