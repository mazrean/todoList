import { Message, Error } from "./common";
import type { TaskStatusDetail } from "./taskStatus";

export class DashboardMeta {
  name: string;
  description: string;
}

export class DashboardInfo {
  id: string;
  name: string;
  description: string;
  createdAt: Date;
}

export class DashboardDetail {
  id: string;
  name: string;
  description: string;
  createdAt: Date;
  taskStatusList: TaskStatusDetail[];
}

export async function postDashboard(dashboard: DashboardMeta): Promise<DashboardInfo> {
  return fetch("/api/dashboards", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(dashboard)
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 201) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as DashboardInfo;
  });
}

export async function patchDashboard(dashboardID: string, dashboard: DashboardMeta): Promise<DashboardInfo> {
  return fetch(`/api/dashboards/${dashboardID}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(dashboard)
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as DashboardInfo;
  });
}

export async function deleteDashboard(dashboardID: string): Promise<Message> {
  return fetch(`/api/dashboards/${dashboardID}`, {
    method: "DELETE"
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as Message;
  });
}

export async function getMyDashboards(): Promise<DashboardInfo[]> {
  return fetch(`/api/users/me/dashboards`).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as DashboardDetail[];
  });
}

export async function getDashboardInfo(dashboardID: string): Promise<DashboardDetail> {
  return fetch(`/api/dashboards/${dashboardID}`).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as DashboardDetail;
  });
}
