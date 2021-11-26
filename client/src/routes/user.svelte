<script type=ts>
  import UserForm from "../components/UserForm.svelte";
  import { toast } from '@zerodevx/svelte-toast';
  import { deleteMeAction, updateMeAction, user } from "../store/user";
  import type { Error } from '../api/common';
  import { UserInfo } from "../api/user";
  import { goto } from "$app/navigation";

  let me: string|null = null;
  user.subscribe(user => {
    me = user;
  });

  async function submit(event: any) {
    updateMeAction(new UserInfo(
      event.detail.name,
      event.detail.password,
    )).then(() => {
      toast.push("ユーザー情報の更新に成功しました");
    }).catch((err: Error) => {
      toast.push(`${err.code}:${err.error}`, {
        theme: {
          '--toastBackground': '#F56565',
          '--toastBarBackground': '#C53030'
        }
      });
    });
  }

  async function deleteAccount(event: any) {
    deleteMeAction().then(() => {
      toast.push("アカウントの削除に成功しました");
      goto("/signup");
    }).catch((err: Error) => {
      toast.push(`${err.code}:${err.error}`, {
        theme: {
          '--toastBackground': '#F56565',
          '--toastBarBackground': '#C53030'
        }
      });
    });
  }
</script>

<div class="container">
  <div class="content">
    <h3>Change Account Information</h3>
    <UserForm name={me} label="update" on:submit={submit} />
  </div>
  <div class="content">
    <h3>Delete Account</h3>
    <button on:click={deleteAccount}>delete</button>
  </div>
</div>

<style>
  .container {
    height: 100%;
    width: 100%;
  }
  .content {
    margin: 0 5px;
    margin-bottom: 25px;
  }
  h3 {
    font-size: 24px;
    line-height: 1.4;
    color: #222;
    margin: 0;
  }
  button {
    cursor: pointer;
    background-color: #f0506e;
    color: #fff;
    border: 1px solid transparent;
  }
</style>
