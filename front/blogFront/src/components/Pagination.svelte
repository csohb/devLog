<script lang="ts">
import { beforeUpdate, createEventDispatcher } from "svelte";

const dispatch = createEventDispatcher();

export let count: number = 0;
export let page: number = 0;
export let pageSize: number = 10;

let buttons = [];
export const labels = {
  first: "<<",
  last: ">>",
  next: ">",
  previous: "<",
};
beforeUpdate(() => {
  // console.log(
  //   "page update count:",
  //   count,
  //   "page:",
  //   page,
  //   "pageSize:",
  //   pageSize,
  //   "pageCount:",
  //   pageCount,
  //   "block",
  //   block
  // );
  range(pageSize);
});

const getPageCount = (cnt: number): number => {
  let val = cnt;
  if (cnt === 0) {
    return 0;
  }
  let ret = Math.ceil(val / pageSize);
  return ret;
};

$: pageCount = getPageCount(count);
$: block = Math.floor(page / 10) * 10;

function range(size: number): number[] {
  let tmp = [];
  if (pageCount == 0) {
    tmp = [0];
  } else {
    for (let i = 0; i < size; i++) {
      let n = i + block;
      if (n >= pageCount) {
        break;
      }
      tmp.push(i + block);
    }
  }
  buttons = tmp;
  return buttons;
}

function onChange(p: number) {
  page = p < 0 ? 0 : p < pageCount ? p : pageCount;
  dispatch("change", p);
}
</script>

<div class="pagination-wrap">
  <ul>
    <li>
      <button
        class="box"
        disabled="{block === 0}"
        on:click="{() => onChange(0)}">
        {labels.first}
      </button>
    </li>
    <li>
      <button
        class="box"
        disabled="{block === 0}"
        on:click="{() => onChange(page - pageSize)}">
        {labels.previous}
      </button>
    </li>
    {#each buttons as idx}
      <li>
        <button class:active="{page === idx}" on:click="{() => onChange(idx)}">
          {idx + 1}
        </button>
      </li>
    {/each}
    <li>
      <button
        class="box"
        disabled="{block + pageSize > pageCount}"
        on:click="{() => onChange(page + pageSize)}">
        {labels.next}
      </button>
    </li>
    <li>
      <button
        class="box"
        disabled="{block + pageSize > pageCount}"
        on:click="{() => onChange(pageCount)}">
        {labels.last}
      </button>
    </li>
  </ul>
</div>
