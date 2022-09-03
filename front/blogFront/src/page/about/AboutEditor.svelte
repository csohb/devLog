<script lang="ts">
import Footer from "../../components/Footer.svelte";
import Header from "../../components/Header.svelte";
import { push, querystring } from "svelte-spa-router";
import aboutStore from "../../stores/about";
import popupStore from "../../stores/popup";
import authStore from "../../stores/auth";
import type { CrateSkillRequest, SkillListType } from "../../api/types/about";
import { onMount, tick } from "svelte";

let careerList = [];
let company = "";
let startDate = "";
let endDate = "";
let jobTitle = "";
let jobDetail = "";

let skillList: SkillListType[] = [];
let skillName = "";
let percentage: number | null = null;

let projectList = [];
let pjStartDate = "";
let pjEndDate = "";
let projectName = "";
let pjDescription = "";
let pjStack = "";
let projectStackList = [];
let pjIsPersonal = false;

onMount(() => {
  if ($authStore.loginNick === "") {
    let loginNick = authStore.getCookie("user_id");
    if (loginNick == null) {
      push("/about");
    } else {
      authStore.setNick(loginNick);
    }
  }
});

function onClickCareerAdd() {
  // api 등록전 커리어 추가
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

  careerList = careerList;

  if (listCount !== careerList.length) {
    popupStore.open({
      title: "커리어 추가",
      text: "추가되었습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        resetCareer();
      },
    });
  } else {
    alert("저장 실패!");
  }
}
function resetCareer() {
  company = "";
  startDate = "";
  endDate = "";
  jobDetail = "";
  jobTitle = "";
}
function onClickCareerSave() {
  if (!validationQuery()) {
    return;
  }
  if (careerList.length === 0) {
    return;
  }
  aboutStore
    .crateCareer($querystring.split("=")[1], careerList)
    .then(() => {
      alert("저장완료");
    })
    .catch((err) => {
      console.log(err);
    });
}
function onClickSkillAdd() {
  if (skillName.trim() === "") {
    alert("스킬명을 입력해주세요.");
    return;
  }
  let skillItem: SkillListType = {
    name: skillName,
  };

  if (percentage != null) {
    skillItem.percentage = Number(percentage);
  }
  skillList.push(skillItem);
  // 반응성을 위해 추가
  skillList = skillList;
}
function validationQuery(): boolean {
  if ($querystring.split("=")[1] == null) {
    return false;
  }
  if ($querystring.split("=")[1] === "") {
    return false;
  }
  return true;
}
function onClickSkillSave() {
  if (!validationQuery()) {
    return;
  }
  if (skillList.length === 0) {
    return;
  }
  let req: CrateSkillRequest = {
    user_id: $querystring.split("=")[1],
    tech: skillList,
  };
  aboutStore
    .crateSkill(req)
    .then(() => {
      alert("저장완료");
    })
    .catch((err) => {
      console.log(err);
    });
}
function resetSkill() {
  skillName = "";
  percentage = 0;
}

function onClickProjectSave() {
  if (!validationQuery()) {
    return;
  }
  if (projectList.length === 0) {
    return;
  }
  aboutStore
    .crateProject($querystring.split("=")[1], projectList)
    .then(() => {
      alert("프로젝트 저장완료");
    })
    .catch((err) => {
      console.log(err);
    });
}

function onClickProjectAdd() {
  if (projectName.trim() === "") {
    alert("프로젝트명을 입력해주세요.");
    return;
  }
  if (pjStartDate.trim() === "") {
    // TODO: 2022-05-03 형식으로만 입력할 수 있도록 수정 필요
    alert("프로젝트 시작일을 입력해주세요.");
    return;
  }

  const listCount = projectList.length;

  projectList.push({
    name: projectName,
    is_personal: pjIsPersonal,
    start_date: pjStartDate,
    end_date: pjEndDate,
    description: pjDescription,
    stack: projectStackList,
  });
  projectList = projectList;

  if (listCount !== projectList.length) {
    popupStore.open({
      title: "프로젝트 추가",
      text: "추가되었습니다.",
      type: "alert",
      isShow: false,
      action: () => {
        resetProject();
      },
    });
  } else {
    alert("저장 실패!");
  }
}
function resetProject() {
  projectName = "";
  pjIsPersonal = false;
  pjStartDate = "";
  pjEndDate = "";
  pjDescription = "";
  pjStack = "";
  projectStackList = [];
}
async function onClickStackAdd() {
  console.log(pjStack, projectStackList);
  if (pjStack.trim() === "") {
    return;
  }

  projectStackList.push(pjStack);
  await tick();
  projectStackList = projectStackList;
  pjStack = "";
}
function onClickRouter() {
  popupStore.open({
    title: `${$querystring.split("=")[1]}소개 수정`,
    text: "확인하시면 작성하신 내용이 사라집니다.<br />내용을 저장하시려면 저장 버튼을 눌러주세요.<br />페이지를 이동하시겠습니까?",
    type: "confirm",
    btn: "확인",
    isShow: false,
    action: () => {
      resetCareer();
      resetSkill();
      push("/about");
    },
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
                <input
                  type="text"
                  bind:value="{startDate}"
                  placeholder="####-##-##" />
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
            <ul>
              {#each careerList as career}
                회사명: {career.company}
              {/each}
            </ul>
            <button type="button" on:click="{onClickCareerAdd}">추가</button>
            <button type="button" on:click="{onClickCareerSave}">저장</button>
          </div>
          <div class="sub-about-edit">
            <h2>SKILLS</h2>
            <ul>
              <li>
                <span>skill :</span>
                <input type="text" bind:value="{skillName}" />
              </li>
              <li>
                <span>percentage :</span>
                <input type="number" min="0" bind:value="{percentage}" />
              </li>
            </ul>

            <ul>
              {#each skillList as skill}
                <li>
                  name: {skill.name} <br />
                  %: {skill.percentage}
                </li>
              {/each}
            </ul>
            <button type="button" on:click="{onClickSkillAdd}">추가</button>
            <button type="button" on:click="{onClickSkillSave}">저장</button>
          </div>
          <div class="sub-about-edit">
            <h2>Project</h2>
            <ul>
              <li style="display: flex;align-items: center;">
                <span>개인작업 여부:</span>
                <div class="switch">
                  <input
                    type="checkbox"
                    id="switch"
                    bind:checked="{pjIsPersonal}" />
                  <label for="switch" class="switch_label">
                    <span class="switch_btn"></span>
                  </label>
                </div>
              </li>
              <li>
                <span>시작일:</span>
                <input type="text" bind:value="{pjStartDate}" />
              </li>
              <li>
                <span>종료일:</span>
                <input type="text" bind:value="{pjEndDate}" />
              </li>
              <li>
                <span>프로젝트명:</span>
                <input type="text" bind:value="{projectName}" />
              </li>
              <li>
                <span>프로젝트 설명:</span>
                <input type="text" bind:value="{pjDescription}" />
              </li>
              <li>
                <span>사용 스텍:</span>
                <input type="text" bind:value="{pjStack}" />
                <button type="button" on:click="{onClickStackAdd}"
                  >스텍 추가</button>
                {projectStackList}
              </li>
            </ul>
            {#if projectList.length !== 0}
              <ul>
                {#each projectList as item}
                  <li>개인작업 여부 :{item.is_personal}</li>
                  <li>프로젝트명: {item.name}</li>
                  <li>시작일 : {item.start_date}</li>
                {/each}
              </ul>
            {/if}
            <button type="button" on:click="{onClickProjectAdd}">추가</button>
            <button type="button" on:click="{onClickProjectSave}">저장</button>
          </div>
        </div>
        <div class="sub-about-edit-action">
          <button type="button" on:click="{onClickRouter}"
            >소개 페이지로 이동하기</button>
        </div>
      </div>
    </div>
  </section>
</main>
<Footer />
