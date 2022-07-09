<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import authStore from "../../stores/auth"
import {push } from "svelte-spa-router"

let id:string = ""
let pw:string = ""

function loginHandler(){
  authStore.postLogin(id, pw).then(()=>{
    id = ""
    pw = ""
    push('/')
  }).catch(()=>{
    id = ""
    pw = ""
    alert('로그인 실패')
  })
}

</script>

<Header />
<main class="sub-page">
    <section class="sub-login-wrap">
      <div class="inner">
        <div class="login-form">
            <h2>DEVLOG</h2>
            <form action="#">
                <input type="text" name="id" bind:value={id} placeholder="아이디를 입력해주세요.">
                <input type="password" name="password" bind:value={pw} placeholder="비밀번호를 입력해주세요.">
                <button type="button" on:click={loginHandler}>로그인 하기</button>
            </form>
            <p># 가입은 관리자에게 문의해주세요.</p>
        </div>
      </div>
    </section> 
  </main>
<Footer />