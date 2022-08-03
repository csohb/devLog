<script lang="ts">
import aboutStroe from "../../stores/about";

function imgError(target: any) {
  target.path[0].style.display = "none";
}
</script>

<h2>Project</h2>
<div class="timeline">
  <ul>
    {#if $aboutStroe.project.length !== 0}
      {#each $aboutStroe.project as project, idx}
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
            <h4>{project.description}</h4>
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
