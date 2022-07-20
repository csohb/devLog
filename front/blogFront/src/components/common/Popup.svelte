<script lang="ts">
import popupStore from "../../stores/popup";

$: console.log("popup stroe:", $popupStore);

function onClickClose() {
  popupStore.cancel();
}

function onClickOK() {
  popupStore.close();
}
</script>

{#if $popupStore.popupTask.length !== 0 && $popupStore.popupTask[0].isShow}
  <div class="popup-bg"></div>
  <div class="popup-wrap">
    <p class="popup-title">{$popupStore.popupTask[0].title}</p>
    <p class="popup-text">{$popupStore.popupTask[0].text}</p>
    <div class="popup-action-wrap">
      {#if $popupStore.popupTask[0].type === "confirm"}
        <button type="button" on:click="{onClickOK}"
          >{$popupStore.popupTask[0].btn}</button>
      {/if}
      <button type="button" on:click="{onClickClose}">닫기</button>
    </div>
  </div>
{/if}
