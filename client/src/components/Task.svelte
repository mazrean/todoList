<script type=ts>
  import { TaskInfo } from "../api/task";
  import { createEventDispatcher } from "svelte";

  export let task: TaskInfo = new TaskInfo();

  let edit = false;

  const dispatch = createEventDispatcher();
  function toggle() {
    if (edit) {
      dispatch('submit', {
        id: task.id,
        name: task.name,
        description: task.description,
      });
    }
		edit = !edit;
	}
</script>

<div>
  {#if edit}
  <h4><input bind:value={task.name}></h4>
  {:else}
  <h4>{task.name}</h4>
  {/if}
  <p>{new Date(task.createdAt).toLocaleString()}</p>
  {#if edit}
  <p><textarea bind:value={task.description}></textarea></p>
  {:else}
  <p>{task.description}</p>
  {/if}
  <button class="primary" on:click={toggle}>
    {#if edit}
    save
    {:else}
    edit
    {/if}
  </button>
  <button class="primary" on:click={() => dispatch('delete', {id: task.id})}>delete</button><br>
  <button class="normal" on:click={() => dispatch('left', {id:task.id})}>left</button>
  <button class="normal" on:click={() => dispatch('right', {id:task.id})}>right</button>
</div>

<style>
  h4 {
    margin: 5px 0;
    color: #222;
  }
  p {
    margin: 5px 0;
  }
  div {
    overflow-wrap: anywhere;
    padding: 5px;
    margin: 5px 0;
    transition: box-shadow .1s ease-in-out;
    background-color: #fff;
    color: #666;
    box-shadow: 0 5px 15px rgb(0 0 0 / 8%);
    border-radius: 5px;
  }
  .normal {
    cursor: pointer;
    background-color: transparent;
    color: #222;
    border: 1px solid #e5e5e5;
    margin: 5px;
    height: 24px;
    width: 51px;
  }
  .primary {
    cursor: pointer;
    background-color: #1e87f0;
    color: #fff;
    border: 1px solid transparent;
    margin: 5px;
    height: 24px;
    width: 51px;
  }
</style>
