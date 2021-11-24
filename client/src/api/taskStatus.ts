import { Message, Error } from "./common";
import type { TaskInfo } from "./task";

export class TaskStatusInfo {
  id: string;
  name: string;
}

export class TaskStatusDetail {
  id: string;
  name: string;
  tasks: TaskInfo[];
}

export async function postTaskStatus(dashboardID: string, name: string): Promise<TaskStatusInfo> {
  return fetch(`/api/dashboards/${dashboardID}/status`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      name
    })
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 201) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as TaskStatusInfo;
  });
}

export async function patchTaskStatus(dashboardID: string, id: string, name: string): Promise<TaskStatusInfo> {
  return fetch(`/api/tasks/status/${id}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      name
    })
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as TaskStatusInfo;
  });
}

export async function deleteTaskStatus(id: string): Promise<Message> {
  return fetch(`/api/tasks/status/${id}`, {
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
