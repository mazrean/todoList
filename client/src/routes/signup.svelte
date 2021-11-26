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

<div class="container">
  <h3>Signup</h3>
  <UserForm label="signup" on:submit={submit} />
</div>

<style>
  h3 {
    font-size: 24px;
    line-height: 1.4;
    color: #222;
    margin: 0;
  }
  .container {
    height: 100%;
    width: 100%;
  }
</style>
