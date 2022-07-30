<script lang="ts">
import type { CareerListType } from "../../api/types/about";

import aboutStore from "../../stores/about";
export let isEditMode = false;
export let currentTab = "";

let company = "";
let startDate = "";
let endDate = "";
let jobTitle = "";
let jobDetail = "";

async function onClickEdit(id: string) {
  // career 에 get 에는 id 가 따로 없는데 index를 의미하는 건지 확인 필요
  console.log("수정 id:", id);
  let req: CareerListType = {
    company,
    start_date: startDate,
    end_date: endDate,
    job_title: jobTitle,
    job_detail: jobDetail,
  };
  await aboutStore
    .updateCareer(id, req)
    .then((resp) => {
      console.log(resp);
      if (currentTab !== "") {
        aboutStore.setIntroduce(currentTab);
      }
    })
    .finally(() => {
      company = "";
      startDate = "";
      endDate = "";
      jobTitle = "";
      jobDetail = "";
    });
}
function onClickDelete(id: string) {
  console.log("삭제 id:", id);
}
</script>

<h2>Career</h2>
<div class="timeline">
  <ul>
    {#if $aboutStore.careers.length !== 0}
      {#each $aboutStore.careers as career, idx}
        <li>
          <div
            class="bullet"
            class:pink="{idx % 3 === 0}"
            class:green="{idx % 3 === 1}"
            class:orange="{idx % 3 === 2}">
          </div>
          <div class="time">
            {career.start_date} ~ <br />{career.end_date !== "0001-01-01"
              ? career.end_date
              : "재직중"}
          </div>
          <div class="desc">
            <h3>
              {career.company}
            </h3>
            <h4>
              {career.job_title}
              {#if career.job_detail !== ""} [ {career.job_detail} ]{/if}
              {#if isEditMode}
                <span
                  class="sub-about-edit-btn"
                  on:click="{() => {
                    if (career.isEditMode) {
                      onClickEdit(String(career.id));
                    } else {
                      company = career.company;
                      startDate = career.start_date;
                      endDate = career.end_date;
                      jobTitle = career.job_title;
                      jobDetail = career.job_detail;
                    }
                    aboutStore.careerEditMode(idx);
                  }}">
                  {#if career.isEditMode}완료{:else}수정{/if}</span>
                <span
                  on:click="{() => {
                    onClickDelete(career.id);
                  }}">삭제</span>
              {/if}
            </h4>
            <div class="people"></div>
          </div>
        </li>
        {#if career.isEditMode}
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
              <h3>
                <input type="text" bind:value="{company}" />
              </h3>
              <h4>
                <input type="text" bind:value="{jobTitle}" />
                <br />
                <input type="text" bind:value="{jobDetail}" />
              </h4>
              <div class="people"></div>
            </div>
          </li>
        {/if}
        <!-- <li>
        <div class="bullet green"></div>
        <div class="time">2021-03-19</div>
        <div class="desc">
          <h3>Design Stand Up</h3>
          <h4>Hangouts</h4>
          <div class="people">
            <img
              src="https://s3.amazonaws.com/uifaces/faces/twitter/ashleyford/128.jpg"
              alt="" />
            <img
              src="https://s3.amazonaws.com/uifaces/faces/twitter/kfriedson/128.jpg"
              alt="" />
            <img
              src="https://s3.amazonaws.com/uifaces/faces/twitter/mattsince87/128.jpg"
              alt="" />
          </div>
        </div>
      </li> -->
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
