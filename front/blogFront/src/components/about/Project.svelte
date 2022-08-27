<script lang="ts">
import { tick } from "svelte";

import type { CrateProjectsReq } from "../../api/types/about";

import aboutStore from "../../stores/about";

export let isEditMode = false;
export let currentTab = "";

let projectName = "";
let isPersonal = false;
let startDate = "";
let endDate = "";
let projectDescription = "";
let projectStack = "";
let projectStackList = [];

function imgError(target: any) {
  target.path[0].style.display = "none";
}

function onClcikEdit(id: string) {
  let req: CrateProjectsReq = {
    name: projectName,
    is_personal: isPersonal,
    start_date: startDate,
    end_date: endDate,
    description: projectDescription,
    stack: projectStackList,
  };

  aboutStore
    .updateProject(id, req)
    .then(() => {
      if (currentTab !== "") {
        aboutStore.setIntroduce(currentTab);
      } else {
        alert("수정 성공! 새로고침을 눌러주세요.");
      }
    })
    .finally(() => {
      projectName = "";
      isPersonal = false;
      startDate = "";
      endDate = "";
      projectDescription = "";
      projectStack = "";
      projectStackList = [];
    });
}

function onClickDelete(id: string) {
  aboutStore
    .deleteProject(id)
    .then(() => {
      if (currentTab !== "") {
        aboutStore.setIntroduce(currentTab);
      } else {
        alert("삭제 성공! 새로고침을 눌러주세요.");
      }
    })
    .catch((err) => {
      console.log("프로젝트 삭제 err:", err);
    });
}

function onClickStackAdd() {
  if (projectStack.trim() === "") {
    return;
  }
  projectStackList.push(projectStack);

  projectStackList = projectStackList;
  projectStack = "";
}
</script>

<h2>Project</h2>
<div class="timeline">
  <ul>
    {#if $aboutStore.project.length !== 0}
      {#each $aboutStore.project as project, idx}
        <li>
          <div
            class="bullet"
            class:pink="{idx % 3 === 0}"
            class:green="{idx % 3 === 1}"
            class:orange="{idx % 3 === 2}">
          </div>
          <div class="time">
            {project.start_date} - <br />{project.end_date !== "0001-01-01"
              ? project.end_date
              : "진행중"}
          </div>
          <div class="desc">
            <span>{project.is_personal ? "[ 개인 프로젝트 ]" : ""}</span>
            <h3>
              {project.name}
            </h3>
            <h4>
              {project.description}
              {#if isEditMode}
                <span
                  on:click="{() => {
                    if (project.isEditMode) {
                      onClcikEdit(String(project.id));
                    } else {
                      projectName = project.name;
                      startDate = project.start_date;
                      endDate = project.end_date;
                      projectDescription = project.description;
                      projectStackList = project.stack;
                      isPersonal = project.is_personal;
                    }
                    aboutStore.projectEditMode(idx);
                  }}">
                  {#if project.isEditMode}완료{:else}수정{/if}
                </span>
                <span
                  on:click="{() => {
                    onClickDelete(project.id);
                  }}">
                  삭제
                </span>
              {/if}
            </h4>
            <div class="people">
              {#each project.stack as st}
                <img
                  src="/icon_image/{st}_icon.png"
                  on:error="{imgError}"
                  alt="{st}" />
              {/each}
            </div>
          </div>
        </li>
        {#if project.isEditMode}
          <li>
            <div
              class="bullet"
              class:pink="{idx % 3 === 0}"
              class:green="{idx % 3 === 1}"
              class:orange="{idx % 3 === 2}">
            </div>
            <div class="time">
              <input type="text" bind:value="{startDate}" /> ~
              <br /><input type="text" bind:value="{endDate}" />
            </div>
            <div class="desc">
              <span class="switch">
                <!-- <div class="switch"> -->
                <input
                  type="checkbox"
                  id="switch"
                  bind:checked="{isPersonal}" />
                <label for="switch" class="switch_label">
                  <span class="switch_btn"></span>
                </label>
                <!-- </div> -->
              </span>
              <h3>
                <input type="text" bind:value="{projectName}" />
              </h3>
              <h4>
                <input type="text" bind:value="{projectDescription}" />
              </h4>
              <div class="people">
                {projectStackList}
                <input type="text" bind:value="{projectStack}" />
                <button type="button" on:click="{onClickStackAdd}"
                  >스텍 추가</button>
                <button
                  type="button"
                  on:click="{() => {
                    projectStack = '';
                    projectStackList = [];
                  }}">스텍 초기화</button>
              </div>
            </div>
          </li>
        {/if}
      {/each}
    {:else}
      <li>
        <div class="bullet pink"></div>
        <div class="desc">
          <h3>등록 중 입니다.</h3>
        </div>
      </li>
    {/if}
  </ul>
</div>
