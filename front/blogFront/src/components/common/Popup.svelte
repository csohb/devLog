<script lang="ts">
import popupStore from "../../stores/popup";

function onClickClose() {
  // 팝업 닫기
  popupStore.cancel();
}

function onClickOK() {
  // 확인
  popupStore.close();
}
</script>

{#if $popupStore.popupTask.length !== 0 && $popupStore.popupTask[0].isShow}
  <div class="popup-bg" on:click="{onClickClose}"></div>
  <div class="popup-wrap">
    <p class="popup-title">{$popupStore.popupTask[0].title}</p>
    <p class="popup-text">{@html $popupStore.popupTask[0].text}</p>
    <div class="popup-action-wrap">
      {#if $popupStore.popupTask[0].type === "confirm"}
        <button type="button" on:click="{onClickOK}"
          >{$popupStore.popupTask[0].btn}</button>
        <button type="button" on:click="{onClickClose}">닫기</button>
      {:else}
        <button type="button" on:click="{onClickOK}">닫기</button>
      {/if}
    </div>
  </div>
{/if}
