<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { querystring } from "svelte-spa-router";
</script>

<Header />
<main class="sub-page">
  <section class="sub-about-wrap">
    <div class="inner">
      <div class="sub-about">
        <h1 class="sub-page-title">About Me</h1>
        {$querystring.split("=")[1]}
        <div class="sub-about-edit-wrap">
          <div class="sub-about-edit">
            <h2>CAREER</h2>
            <ul>
              <li>
                <span>취업일:</span> <input type="text" />
              </li>
              <li>
                <span>퇴사일:</span> <input type="text" />
              </li>
              <li>
                <span>회사명:</span> <input type="text" />
              </li>
              <li>
                <span>업무명:</span> <input type="text" />
              </li>
              <li>
                <span>업무상세:</span> <input type="text" />
              </li>
            </ul>
          </div>
          <div class="sub-about-edit">
            <h2>SKILLS</h2>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
