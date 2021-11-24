<script type=ts>
  import UserForm from "../components/UserForm.svelte";
  import { toast } from '@zerodevx/svelte-toast';
  import { updateMeAction, user } from "../store/user";
  import type { Error } from '../api/common';
  import { UserInfo } from "../api/user";

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
</script>

<div class="wrapper">
  <div class="container">
    <h3>ユーザー情報変更</h3>
    <UserForm name={me} label="update" on:submit={submit} />
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
</style>
