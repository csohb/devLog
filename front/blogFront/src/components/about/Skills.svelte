<script lang="ts">
import ProgressBar from "@okrad/svelte-progressbar";
import { beforeUpdate } from "svelte";
import aboutStore from "../../stores/about";

let skills = [];

beforeUpdate(() => {
  skills = $aboutStore.skills.reduce((temp, val) => {
    temp.push({
      title: val.name,
      series: [
        {
          perc: val.percentage,
          color: "#5AB6DF",
        },
      ],
    });
    return temp;
  }, []);
});
</script>

<h2>Major Skills</h2>
<ul class="progress-bar">
  {#if skills.length !== 0}
    {#each skills as item}
      <li class="progress-bar-item">
        <p>{item.title}</p>
        <div>
          <!-- radial -->
          <ProgressBar
            series="{item.series}"
            style="semicircle"
            thickness="{10}" />
        </div>
      </li>
    {/each}
  {:else}
    <li>등록 중 입니다.</li>
  {/if}
</ul>
