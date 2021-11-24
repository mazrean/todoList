import { Message, Error } from "./common";

export class UserInfo {
  name: string;
  password: string;
}

export async function postSignup(userInfo: UserInfo): Promise<Message> {
  return await fetch('/api/signup', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userInfo)
  }).catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 201) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as Message;
  });
}

export async function postLogin(userInfo: UserInfo): Promise<Message> {
  return await fetch('/api/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userInfo)
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

export async function getMe(): Promise<UserInfo> {
  return await fetch('/api/users/me').catch(err => {
    throw new Error(-1, err.message);
  }).then(async res => {
    if (res.status !== 200) {
      const err: Error = await res.json();
      throw new Error(res.status, err.error);
    }

    return await res.json() as UserInfo;
  });
}

export async function patchMe(userInfo: UserInfo): Promise<Message> {
  return await fetch('/api/users/me', {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userInfo)
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

export async function deleteMe(): Promise<Message> {
  return await fetch('/api/users/me', {
    method: 'DELETE'
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
