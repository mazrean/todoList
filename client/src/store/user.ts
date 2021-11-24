import type { Message, Error } from "../api/common";
import { UserInfo, getMe, postSignup, postLogin, patchMe, deleteMe } from '../api/user';
import { writable } from 'svelte/store';

export const user = writable<string>(null);

export async function getMeAction(): Promise<void> {
  const userInfo = await getMe().catch((err: Error) => {
    throw err;
  })

  user.set(userInfo.name);
}

export async function signupAction(userInfo: UserInfo): Promise<Message> {
  const message = await postSignup(userInfo).catch((err: Error) => {
    throw err;
  })

  user.set(userInfo.name);

  return message;
}

export async function loginAction(userInfo: UserInfo): Promise<Message> {
  const message = await postLogin(userInfo).catch((err: Error) => {
    throw err;
  })

  user.set(userInfo.name);

  return message;
}

export async function updateMeAction(userInfo: UserInfo): Promise<Message> {
  const message = await patchMe(userInfo).catch((err: Error) => {
    throw err;
  })

  user.set(userInfo.name);

  return message;
}

export async function deleteMeAction(): Promise<Message> {
  const message = await deleteMe().catch((err: Error) => {
    throw err;
  })

  user.set(null);

  return message;
}
