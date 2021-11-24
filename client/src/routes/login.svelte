<script type=ts>
  import UserForm from "../components/UserForm.svelte";
  import { toast } from '@zerodevx/svelte-toast';
  import { loginAction } from "../store/user";
  import type { Error } from '../api/common';
  import { UserInfo } from "../api/user";
  import { goto } from "$app/navigation";

  async function submit(event: any) {
    loginAction(new UserInfo(
      event.detail.name,
      event.detail.password,
    )).catch((err: Error) => {
      toast.push(`${err.code}:${err.error}`, {
        theme: {
          '--toastBackground': '#F56565',
          '--toastBarBackground': '#C53030'
        }
      });
    }).then(() => {
      goto('/');
    });
  }
</script>

<div class="wrapper">
  <div class="container">
    <UserForm label="login" on:submit={submit} /><br>
    <a href="/signup">アカウント作成</a>
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
  a {
    color: #666;
  }
</style>
