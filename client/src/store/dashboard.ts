import { DashboardInfo, getMyDashboards } from "../api/dashboard";
import { writable } from "svelte/store";

export const dashboards = writable<DashboardInfo[]>([]);

export async function getDashboardsAction(): Promise<void> {
  const dashboardInfos = await getMyDashboards().catch((err: Error) => {
    throw err;
  });

  dashboards.set(dashboardInfos);
}
