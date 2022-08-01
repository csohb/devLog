<script lang="ts">
import aboutStore from "../../stores/about";
</script>

<div class="sub-about-introduce">
  <pre>{#if $aboutStore.description === ""}description 등록 중 입니다.{/if}</pre>
</div>
<div class="sub-about-info">
  <h2>Info</h2>
  <ul>
    <li>
      <strong>Name</strong>{$aboutStore.info.name} ( {$aboutStore.info
        .nick_name}
      )
    </li>
    <li><strong>Date of birth</strong>{$aboutStore.info.birthday}</li>
    <li><strong>Address</strong>{$aboutStore.info.addr}</li>
    <li><strong>E-mail</strong>{$aboutStore.info.email}</li>
  </ul>
</div>
