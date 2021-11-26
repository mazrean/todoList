<script type=ts>
  import { DashboardDetail, getDashboardInfo } from "../../api/dashboard";
  import { page } from '$app/stores';
  import TaskStatus from "../../components/TaskStatus.svelte";
  import { toast } from "@zerodevx/svelte-toast";
  import type { Error } from '../../api/common';
  import type { TaskStatusDetail } from "../../api/taskStatus";
  import { deleteTask, patchMoveTask, patchTask, postTask } from "../../api/task";
import { beforeUpdate } from "svelte";

  let dashboardID = "";
  let dashboardName = "";
  let dashboardDescription = "";
  let taskStatusList: TaskStatusDetail[] = [];
  async function updateInfo() {
    await getDashboardInfo($page.params.slug).then(res => {
      dashboardName = res.name;
      dashboardDescription = res.description;
      taskStatusList = res.taskStatusList;
    }).catch((err: Error) => {
      toast.push(`${err.code}:${err.error}`, {
        theme: {
          '--toastBackground': '#F56565',
          '--toastBarBackground': '#C53030'
        }
      });
    });
  }

  beforeUpdate(async () => {
    if (dashboardID != $page.params.slug) {
      await updateInfo();
      dashboardID = $page.params.slug;
    }
  });

  function submit(id: string) {
    return async function(event: any) {
      await postTask(id, {
        name: event.detail.name,
        description: event.detail.description,
      }).then(task => {
        toast.push("task created");
        for (let i = 0; i < taskStatusList.length; i++) {
          if (taskStatusList[i].id === event.detail.id) {
            taskStatusList[i].tasks.push(task);
            break;
          }
        }
        updateInfo();
      }).catch((err: Error) => {
        toast.push(`${err.code}:${err.error}`, {
          theme: {
            '--toastBackground': '#F56565',
            '--toastBarBackground': '#C53030'
          }
        });
      });
    }
  }

  async function edit(event: any) {
    await patchTask(event.detail.id, {
      name: event.detail.name,
      description: event.detail.description,
    }).then(task => {
      toast.push("task updated");
      for (let i = 0; i < taskStatusList.length; i++) {
        if (taskStatusList[i].id === event.detail.id) {
          for (let j = 0; j < taskStatusList[i].tasks.length; j++) {
            if (taskStatusList[i].tasks[j].id === event.detail.id) {
              taskStatusList[i].tasks[i] = task;
              break;
            }
          }
          break;
        }
      }
    }).catch((err: Error) => {
      toast.push(`${err.code}:${err.error}`, {
        theme: {
          '--toastBackground': '#F56565',
          '--toastBarBackground': '#C53030'
        }
      });
    });
  }

  async function deleteHandler(event: any) {
    await deleteTask(event.detail.id).then(() => {
      toast.push("task deleted");
      updateInfo();
    }).catch((err: Error) => {
      toast.push(`${err.code}:${err.error}`, {
        theme: {
          '--toastBackground': '#F56565',
          '--toastBarBackground': '#C53030'
        }
      });
    });
  }

  function move(statusID: string) {
    return async function(event: any) {
      await patchMoveTask(event.detail.id, statusID).then(() => {
        toast.push("task moved");
        updateInfo();
      }).catch((err: Error) => {
        toast.push(`${err.code}:${err.error}`, {
          theme: {
            '--toastBackground': '#F56565',
            '--toastBarBackground': '#C53030'
          }
        });
      });
    }
  }
</script>

<div class="container">
  <div class="row header">
    <h3>{dashboardName}</h3>
    <p>{dashboardDescription}</p>
  </div>
  {#if taskStatusList.length === 0}
    <div class="row">
      <p>No task Status</p>
    </div>
  {:else}
    <div class="row" style="display: grid;grid-template-columns: repeat({taskStatusList.length}, 1fr)">
      {#each taskStatusList as taskStatus, i}
        <div class="column">
          <TaskStatus
            name={taskStatus.name}
            tasks={taskStatus.tasks}
            on:submit={submit(taskStatus.id)}
            on:edit={edit}
            on:delete={deleteHandler}
            on:left={move(taskStatusList[i-1].id)}
            on:right={move(taskStatusList[i+1].id)}
          />
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  h3 {
    font-size: 24px;
    line-height: 1.4;
    color: #222;
    display:inline;
    margin: 0;
  }
  .container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
  }
  .row {
    width: 100%;
  }
  .column {
    margin: 5px;
  }
  .header {
    padding: 5px;
  }
</style>
