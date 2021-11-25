<script type=ts>
  import { DashboardDetail, getDashboardInfo } from "../../api/dashboard";
  import { page } from '$app/stores';
  import TaskStatus from "../../components/TaskStatus.svelte";
  import { toast } from "@zerodevx/svelte-toast";
  import type { Error } from '../../api/common';
  import type { TaskStatusDetail } from "../../api/taskStatus";
  import { deleteTask, patchTask, postTask } from "../../api/task";

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

  updateInfo()

  function submit(id: string) {
    return async function(event: any) {
      postTask(id, {
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
    patchTask(event.detail.id, {
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
    deleteTask(event.detail.id).then(message => {
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
</script>

<div class="wrapper">
  <div class="container">
    <div>
      <h3>{dashboardName}</h3>
      <p>{dashboardDescription}</p>
    </div>
    {#if taskStatusList.length === 0}
      <div>
        <p>No task Status</p>
      </div>
    {:else}
      <div style="display: grid;grid-template-columns: repeat({taskStatusList.length}, 1fr)">
        {#each taskStatusList as taskStatus}
          <div class="column">
            <TaskStatus name={taskStatus.name} tasks={taskStatus.tasks} on:submit={submit(taskStatus.id)} on:edit={edit} on:delete={deleteHandler} />
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>

<style>
  h3 {
    margin-bottom: 5px;
  }
  .wrapper {
    margin: 15px auto;
    width: 100%;
  }
  .container {
    margin: 0 auto;
    width: 75%;
    display: grid;
    grid-template-rows: auto auto;
  }
</style>
