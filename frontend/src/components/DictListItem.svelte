<script lang="ts">
  import type { DictFile } from '../types/dict'
  import { notifications, Severity } from '../stores/notification' 

  export let dict: DictFile 
  export let selected = false
  export let select: (name: string) => void 

  const notify = () => notifications.addNotificaton("Hi from the notification service!", Severity.success)
</script>

<li>
  <div>
    <button class={`checkbox ${selected && "selected"}`} aria-label={`select ${dict.Name}`} on:click={() => select(dict.Name)}></button>
    {dict.Display}
  </div>
  <span class="list-btn-container">
    <span>{dict.Modified.split("T")[0]}</span>
    <span>{dict.Size/1000} kb</span>
    <button type="button" on:click={notify} >delete</button>
  </span>
</li>

<style>
  li {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 24px 12px;
    list-style: none;
  }
  span{
    color: #5d5d5d
  }
  .checkbox{
    border: 1px solid black;
    background-color: white;
    height: 16px;
    width: 16px;
    border-radius: 2px;
    cursor: pointer;
  }
  .selected{
    background-color: #1F797E; 
  }
  .list-btn-container > * {
    margin-right: 16px;
  }
</style>
