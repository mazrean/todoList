<script type=ts>
import { createEventDispatcher } from "svelte";
import type { TaskInfo } from "../api/task";
import Task from "./Task.svelte";
import TaskForm from "./TaskForm.svelte";

export let name: string;
export let tasks: TaskInfo[] = [];

const dispatch = createEventDispatcher();
function edit(event: any) {
  dispatch('edit', event.detail);
}
</script>

<div class="wrapper">
  <h3>{name}</h3>
  <TaskForm label="create" on:submit></TaskForm><br>
  {#if tasks.length === 0}
    No Task
  {:else}
    {#each tasks as task}
      <div class="container">
        <Task task={task} on:submit={edit} on:delete on:left on:right />
      </div>
    {/each}
  {/if}
</div>

<style>
  h3 {
    margin: 5px;
  }
  .wrapper {
    background-color: #f8f8f8;
    padding: 5px;
    border-radius: 5px;
  }
</style>
