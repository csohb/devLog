<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";

</script>

<Header />
<main class="sub-page">
    <section class="sub-login-wrap">
      <div class="inner">
        <div class="login-form">
            <h2>DEVLOG</h2>
            <form action="#">
                <input type="text" name="id" placeholder="아이디를 입력해주세요.">
                <input type="password" name="password" placeholder="비밀번호를 입력해주세요.">
                <button type="button">로그인 하기</button>
            </form>
            <p># 가입은 관리자에게 문의해주세요.</p>
        </div>
      </div>
    </section> 
  </main>
<Footer />