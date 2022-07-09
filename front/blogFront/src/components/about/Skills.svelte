<script lang="ts">
import ProgressBar from "@okrad/svelte-progressbar";
let skills = [
  {
    title: "Python",
    series: [
      {
        perc: 10,
        color: "#5AB6DF",
      },
    ],
  },
  {
    title: "JavaScript",
    series: [
      {
        perc: 80,
        color: "#5AB6DF",
      },
    ],
  },
];
</script>

<h2>Major Skills</h2>
<ul class="progress-bar">
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
</ul>
