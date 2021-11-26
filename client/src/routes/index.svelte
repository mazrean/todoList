<script type=ts>
  import { getMyDashboards } from "../api/dashboard";
  import type { Item } from "../components/LinkList";
  import { toast } from '@zerodevx/svelte-toast';
  import type { Error } from '../api/common';
  import LinkList from "../components/LinkList.svelte";
import { goto } from "$app/navigation";

  let dashboards: Item[] = [];
  getMyDashboards().then(dashboardInfos => {
    let dashboardList: Item[] = [];
    for (let dashboard of dashboardInfos) {
      dashboardList.push({
        label: dashboard.name,
        link: `/dashboard/${dashboard.id}`,
      })
      dashboards = dashboardList;
    }
  }).catch((err: Error) => {
    toast.push(`${err.code}:${err.error}`, {
      theme: {
        '--toastBackground': '#F56565',
        '--toastBarBackground': '#C53030'
      }
    });
  });
</script>

<div class="container">
  <div class="header">
    <h3>My Dashboards</h3>
    <button on:click={()=>goto("/dashboard/new")}>New</button>
  </div>
  <div class="content">
    <LinkList items={dashboards} />
  </div>
</div>


<style>
  .container {
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
  }
  .content {
    margin: 5px;
  }
  .header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 5px;
  }
  h3 {
    font-size: 24px;
    line-height: 1.4;
    color: #222;
    display:inline;
    margin: 0;
  }
  button {
    cursor: pointer;
    background-color: #1e87f0;
    color: #fff;
    border: 1px solid transparent;
    margin: 5px;
    height: 24px;
  }
</style>
