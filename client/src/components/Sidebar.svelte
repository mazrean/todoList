<script type=ts>
import { derived } from "svelte/store";

import { dashboards } from "../store/dashboard";
import type { Item } from "./LinkList";
import LinkList from "./LinkList.svelte";

let dashboardList: Item[] = [];
dashboards.subscribe(dashboardInfos => {
  let newDashboardList = [];
  for (let dashboard of dashboardInfos) {
    newDashboardList.push({
      label: dashboard.name,
      link: `/dashboard/${dashboard.id}`,
    })
  }
  dashboardList = newDashboardList;
});

let userList: Item[] = [
  {
    label: "setting",
    link: "/user",
  },
];
</script>

<div class="container">
  <div class="content">
    <h4>dashboards</h4>
    <div class="list">
      <LinkList items={dashboardList} />
    </div>
  </div>
  <div class="content">
    <h4>user</h4>
    <div class="list">
      <LinkList items={userList} />
    </div>
  </div>
</div>

<style>
  .container {
    width: 100%;
  }
  .content {
    margin: 15px;
  }
  .list {
    padding-left: 10px;
  }
  h4 {
    margin: 0;
  }
</style>
