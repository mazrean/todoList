import type { Message, Error } from "src/api/common";
import { UserInfo, getMe, postSignup, postLogin } from 'src/api/user';
import { writable } from 'svelte/store';

export const user = writable<UserInfo>(null);

export async function getMeAction(): Promise<void> {
  const userInfo = await getMe().catch((err: Error) => {
    throw err;
  })

  user.set(userInfo);
}

export async function signupAction(userInfo: UserInfo): Promise<Message> {
  const message = await postSignup(userInfo).catch((err: Error) => {
    throw err;
  })

  user.set(userInfo);

  return message;
}

export async function loginAction(userInfo: UserInfo): Promise<Message> {
  const message = await postLogin(userInfo).catch((err: Error) => {
    throw err;
  })

  user.set(userInfo);

  return message;
}
