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

<div class="wrapper">
  <div class="container">
    <h3>ユーザー情報変更</h3>
    <UserForm name={me} label="update" on:submit={submit} />
    <h3>アカウント削除</h3>
    <button on:click={deleteAccount}>delete</button>
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
  h3 {
    margin-bottom: 5px;
  }
  button {
    cursor: pointer;
    background-color: transparent;
    color: #222;
    border: 1px solid #e5e5e5;
  }
</style>
