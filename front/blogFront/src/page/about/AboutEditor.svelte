<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { querystring } from "svelte-spa-router";
import aboutStore from "../../stores/about";
import popupStore from "../../stores/popup";

let careerList = [];
let company = "";
let startDate = "";
let endDate = "";
let jobTitle = "";
let jobDetail = "";

function onClickCareerAdd() {
  if (company.trim() === "") {
    alert("회사명을 입력해주세요.");
    return;
  }
  if (startDate.trim() === "") {
    alert("입사일을 입력해주세요.");
    return;
  }
  if (jobTitle.trim() === "") {
    alert("업무명을 입력해주세요.");
    return;
  }

  const listCount = careerList.length;

  careerList.push({
    company,
    start_date: startDate,
    end_date: endDate,
    job_title: jobTitle,
    job_detail: jobDetail,
  });

  if (listCount !== careerList.length) {
    popupStore.open({
      title: "커리어 추가",
      text: "추가되었습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        company = "";
        startDate = "";
        endDate = "";
        jobDetail = "";
        jobTitle = "";
      },
    });
  } else {
    alert("저장 실패!");
  }
}
function onClickCareerSave() {
  if ($querystring.split("=")[1] == null) {
    return;
  }
  if ($querystring.split("=")[1] === "") {
    return;
  }
  if (careerList.length === 0) {
    return;
  }
  aboutStore.crateCareer($querystring.split("=")[1], careerList).then(() => {
    alert("저장완료");
  });
}
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
                <span>취업일:</span>
                <input type="text" bind:value="{startDate}" />
              </li>
              <li>
                <span>퇴사일:</span>
                <input type="text" bind:value="{endDate}" />
              </li>
              <li>
                <span>회사명:</span>
                <input type="text" bind:value="{company}" />
              </li>
              <li>
                <span>업무명:</span>
                <input type="text" bind:value="{jobTitle}" />
              </li>
              <li>
                <span>업무상세:</span>
                <input type="text" bind:value="{jobDetail}" />
              </li>
            </ul>
            <button type="button" on:click="{onClickCareerAdd}">추가</button>
            <button type="button" on:click="{onClickCareerSave}">저장</button>
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
