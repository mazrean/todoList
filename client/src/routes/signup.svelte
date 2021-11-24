<script type=ts>
  import UserForm from "../components/UserForm.svelte";
  import { toast } from '@zerodevx/svelte-toast';
  import { signupAction } from "../store/user";
  import type { Error } from '../api/common';
  import { UserInfo } from "../api/user";
  import { goto } from "$app/navigation";

  async function submit(event: any) {
    signupAction(new UserInfo(
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
    <h3>ユーザー作成</h3>
    <UserForm label="signup" on:submit={submit} />
  </div>
</div>

<style>
  h3 {
    margin-bottom: 5px;
  }
  .wrapper {
    margin: 15px auto;
    width: 100%;
    display: flex;
  }
  .container {
    margin: 0 auto;
  }
</style>
