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

<div>
  <UserForm label="login" on:submit={submit} />
</div>

<style>
  div {
    margin: 15px auto;
    width: 100%;
    display: flex;
  }
</style>
