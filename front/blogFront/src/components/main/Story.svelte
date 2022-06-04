<script lang="ts">
  import { Swiper, SwiperSlide } from 'svelte-swiper';
  
  const options = {
    slidesPerView: 1,
   
    pagination: {
      el: '.swiper-pagination',
      clickable: true,
    },
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
    breakpoints: {
        768: {
          slidesPerView: 3,
          spaceBetween: 10,
        },
  }
  };
</script>

<div class="inner">
   <div class="main-stroy-wrap">
       <h1>STORY</h1>
       <Swiper {options}>
        <SwiperSlide> 
            <div class="main-stroy-item">
            <div class="main-story-img"><img src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png" alt="프로필 사진"></div>
                <div class="main-story-text">
                    <p class="main-story-text-label">Backend</p>
                    <h3>제목</h3>
                    <p>dev blog 를 만들기 시작한 계기???????이런 설명 간단하게 노출</p>
                </div>
            </div>
        </SwiperSlide>
        <SwiperSlide> 
            <div class="main-stroy-item">
            <div class="main-story-img"><img src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png" alt="프로필 사진"></div>
                <div class="main-stro-text">
                    <p>Backend</p>
                    <h3>제목</h3>
                    <p>dev blog 를 만들기 시작한 계기???????이런 설명 간단하게 노출</p>
                </div>
            </div>
        </SwiperSlide>
        <SwiperSlide> 
            <div class="main-stroy-item">
            <div class="main-story-img"><img src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png" alt="프로필 사진"></div>
                <div class="main-stro-text">
                    <p>Backend</p>
                    <h3>제목</h3>
                    <p>dev blog 를 만들기 시작한 계기???????이런 설명 간단하게 노출</p>
                </div>
            </div>
        </SwiperSlide>
        <SwiperSlide> 
            <div class="main-stroy-item">
            <div class="main-story-img"><img src="https://ridicorp.com/wp-content/uploads/2022/03/main-rsz-1-940x627.png" alt="프로필 사진"></div>
                <div class="main-stro-text">
                    <p>Backend</p>
                    <h3>제목</h3>
                    <p>dev blog 를 만들기 시작한 계기???????이런 설명 간단하게 노출</p>
                </div>
            </div>
        </SwiperSlide>
      
 
      
        <div class="swiper-pagination" slot="pagination"></div>
        <div class="swiper-button-next" slot="button-next"></div>
        <div class="swiper-button-prev" slot="button-prev"></div>
      </Swiper>
   </div> 
</div>