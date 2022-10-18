<script lang="ts">
import { fetchDeleteSkill, fetchUpdateSkill } from "../../api/about";
import aboutStore from "../../stores/about";

export let isEditMode = false;
export let currentTab = "";
let name = "";
let percentage = 0;

async function onClickDelete(id: string) {
  if (!isEditMode) {
    return;
  }
  if (id === "") {
    return;
  }
  await fetchDeleteSkill(id).then((resp) => {
    aboutStore.setIntroduce(currentTab);
  });
}

async function onClickUpdate(id: string, idx: number) {
  if (!isEditMode) {
    aboutStore.skillEditMode(idx);
    return;
  }
  if (id === "") {
    aboutStore.skillEditMode(idx);
    return;
  }

  let request = {
    id: id,
    name,
  };

  if (percentage !== 0) {
    request["percentage"] = percentage;
  }

  await fetchUpdateSkill(request)
    .then((resp) => {
      aboutStore.setIntroduce(currentTab);
    })
    .catch((err) => {
      console.log("skill update err", err);
    })
    .finally(() => {
      aboutStore.skillEditMode(idx);
      name = "";
      percentage = 0;
    });
}
</script>

<h2>Knowledges</h2>
<ul class="keywords-wrap">
  {#if $aboutStore.keywords.length !== 0}
    {#each $aboutStore.keywords as keyword, idx}
      <li>
        {keyword.name}
        {#if isEditMode}
          <div class="action">
            <button
              type="button"
              on:click="{() => {
                onClickDelete(keyword.id);
              }}">삭제</button>
            {#if keyword.isEditMode}
              <input type="text" bind:value="{name}" />
              <input
                type="number"
                min="0"
                max="100"
                bind:value="{percentage}" />
              <button
                type="button"
                on:click="{() => {
                  onClickUpdate(keyword.id, idx);
                }}">확인</button>
            {:else}
              <button
                type="button"
                on:click="{() => {
                  aboutStore.skillEditMode(idx);
                  name = keyword.name;
                  percentage = aboutStore.getSkillPercentage(keyword.id);
                }}">수정</button>
            {/if}
          </div>
        {/if}
      </li>
    {/each}
  {:else}
    <li>등록 중 입니다.</li>
  {/if}
</ul>
