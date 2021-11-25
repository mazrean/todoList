<script type=ts>
  import { getMyDashboards } from "../api/dashboard";
  import type { Item } from "../components/Sidebar";
  import { toast } from '@zerodevx/svelte-toast';
  import type { Error } from '../api/common';
  import LinkList from "../components/LinkList.svelte";

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

<div class="wrapper">
  <div class="container">
    <LinkList title="Dashboards" items={dashboards} />
  </div>
</div>

<style>
  .wrapper {
    margin: 15px auto;
    width: 100%;
    display: flex;
  }
  .container {
    margin: 0 auto;
  }
</style>
