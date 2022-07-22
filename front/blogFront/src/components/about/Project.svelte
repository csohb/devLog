<script lang="ts">
import aboutStroe from "../../stores/about";
</script>

<h2>Project</h2>
<div class="timeline">
  <ul>
    {#if $aboutStroe.project.length !== 0}
      <li>
        <div class="bullet pink"></div>
        <div class="time">5pm</div>
        <div class="desc">
          <h3>New Icon</h3>
          <h4>Mobile App</h4>
        </div>
      </li>
      <li>
        <div class="bullet green"></div>
        <div class="time">3 - 4pm</div>
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
      </li>
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
