<script lang="ts">
import { onMount } from "svelte";
import Carousel from "svelte-carousel";
import { chevron_left, chevron_right } from "../../icon/Icon";
import { fetchStoryList } from "../../api/story";
import storyStore from "../../stores/story";
import { push } from "svelte-spa-router";

let loading = false;
let screenCount = 3;
let auto = true;
let dots = true;
let arrows = true;
let x = 0;

onMount(() => {
  init();
});
let storyMainList = [];

async function init() {
  await fetchStoryList(1, 5).then((resp) => {
    if (resp === null) {
      return;
    }
    loading = true;
    // TODO: 리스트가 다 불러와져서 5개만 추출
    storyMainList = resp.list.slice(0, 5);
  });
}

function onClickPageMove(id: string) {
  storyStore.viewCount(id);
  push(`/story/${id}`);
}

function resize(x: number) {
  if (x < 769) {
    screenCount = 1;
    auto = true;
    dots = false;
    arrows = false;
  } else {
    screenCount = 3;
    auto = true;
    dots = true;
    arrows = true;
  }
}

$: resize(x);
</script>

<svelte:window bind:innerWidth="{x}" />
<div class="inner">
  <div class="main-stroy-wrap">
    <h1>STORY</h1>
    {#if loading}
      <Carousel
        bind:particlesToShow="{screenCount}"
        particlesToScroll="{1}"
        bind:autoplay="{auto}"
        autoplayDuration="{2000}"
        bind:dots
        bind:arrows
        pauseOnFocus="{true}"
        let:showPrevPage
        let:showNextPage>
        {#each storyMainList as item, index}
          <div
            class="main-stroy-item"
            on:click="{() => {
              onClickPageMove(item.id);
            }}">
            <div class="main-story-img">
              <!-- TODO: 이미지 등록 시 적용 -->
              <img
                src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png"
                alt="프로필 사진" />
            </div>
            <div class="main-story-text">
              <p class="main-story-text-label">Backend/Frontend</p>
              <h3>{item.title}</h3>
              <p class="mian-story-text-description">
                {item.description}
              </p>
            </div>
          </div>
        {/each}
        <div
          slot="prev"
          on:click="{showPrevPage}"
          class="custom-arrow custom-arrow-prev">
          {@html chevron_left}
        </div>
        <div
          slot="next"
          on:click="{showNextPage}"
          class="custom-arrow custom-arrow-next">
          {@html chevron_right}
        </div>
      </Carousel>
    {:else}
      등록된 스토리가 없습니다.
    {/if}
    <!-- <Swiper slidesPerView={1}>
        {#each Array(5) as _ }
        <SwiperSlide> 
            <div class="main-stroy-item">
            <div class="main-story-img"><img src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png" alt="프로필 사진"></div>
                <div class="main-story-text">
                    <p class="main-story-text-label">Backend</p>
                    <h3>제목</h3>
                    <p class="mian-story-text-description">dev blog 를 만들기 시작한 계기???????이런 설명 간단하게 노출</p>
                </div>
            </div>
        </SwiperSlide>
        {/each}
        <div class="swiper-pagination" slot="pagination"></div>
        <div class="swiper-button-next" slot="button-next"></div>
        <div class="swiper-button-prev" slot="button-prev"></div>
      </Swiper> -->
  </div>
</div>
