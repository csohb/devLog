<script lang="ts">
import aboutStore from "../../stores/about";
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
          <div class="time">{career.start_date} ~ <br />{career.end_date}</div>
          <div class="desc">
            <h3>{career.company}</h3>
            <h4>
              {career.job_title}
              {#if career.job_detail !== ""} [ {career.job_detail} ]{/if}
            </h4>
            <div class="people"></div>
          </div>
        </li>
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
