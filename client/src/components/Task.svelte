<script type=ts>
  import { TaskInfo } from "../api/task";
  import { createEventDispatcher } from "svelte";

  export let task: TaskInfo = new TaskInfo();

  let id = task.id;
  let name = task.name;
  let description = task.description;
  let createdAt = new Date(task.createdAt);

  let edit = false;

  const dispatch = createEventDispatcher();
  function toggle() {
    if (edit) {
      dispatch('submit', {
        id,
        name,
        description,
      });
    }
		edit = !edit;
	}
</script>

<div>
  {#if edit}
  <h4><input bind:value={name}></h4>
  {:else}
  <h4>{name}</h4>
  {/if}
  <p>{createdAt.toLocaleString()}</p>
  {#if edit}
  <p><textarea bind:value={description}></textarea></p>
  {:else}
  <p>{description}</p>
  {/if}
  <button on:click={toggle}>
    {#if edit}
    save
    {:else}
    edit
    {/if}
  </button>
  <button on:click={() => dispatch('delete', {id})}>delete</button>
</div>
