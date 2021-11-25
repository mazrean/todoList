import { Message, Error } from "./common";

export class TaskMeta {
  name: string;
  description: string;
}

export class TaskInfo {
  id: string;
  name: string;
  description: string;
  createdAt: string;
}

export async function postTask(taskStatusID: string, task: TaskMeta): Promise<TaskInfo> {
  return fetch(`/api/tasks/status/${taskStatusID}/tasks`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(task)
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 201) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as TaskInfo;
  })
}

export async function patchTask(taskID: string, task: TaskMeta): Promise<TaskInfo> {
  return fetch(`/api/tasks/${taskID}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(task)
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as TaskInfo;
  })
}

export async function deleteTask(taskID: string): Promise<Message> {
  return fetch(`/api/tasks/${taskID}`, {
    method: "DELETE"
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as Message;
  })
}

export async function patchMoveTask(taskID: string, taskStatusID: string): Promise<Message> {
  return fetch(`/api/tasks/${taskID}/move`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      "dest": taskStatusID
    })
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as Message;
  })
}
