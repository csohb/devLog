<script lang="ts">
import { onMount } from "svelte";
import mainStore from "../../stores/main";
import { push } from "svelte-spa-router";

onMount(() => {
  // fetchProfileList().then((res)=>{
  //     console.log(res)
  // })
  mainStore.getProfiles();
});

function imgError() {
  mainStore.imgError();
  console.log("img error");
}

function onClickAboutPage(name: string) {
  push(`/about?name=${name}`);
}

$: console.log("data", $mainStore);
</script>

<div class="inner">
  <div class="main-about-wrap">
    <h1>ABOUT</h1>
    <div class="profile-wrap">
      {#each $mainStore.profiles as item}
        <div
          class="profile-id"
          on:click="{() => {
            onClickAboutPage(item.nick_name);
          }}">
          <div class="profile-img">
            <img
              src="{item.image}"
              alt="{item.nick_name + ' 이미지'}"
              on:error="{imgError}" />
          </div>
          <div class="profile-text">
            <p>{item.nick_name} <span>{item.dev}</span></p>
            <p>{item.intro}</p>
          </div>
        </div>
      {/each}

      <!-- <div class="profile-id">
            <div class="profile-img"><img src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png" alt="프로필 사진"></div>
                <div class="profile-text">
                    <p>이름 <span>Backend Developer</span></p>
                    <p>저는 현재 백엔드 개발을 하고 있으며, 주로 Python 기반 백엔드 프레임워크를 다루고 있습니다.</p>
                </div>
           </div> -->
    </div>
  </div>
</div>
