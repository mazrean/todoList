<script type=ts>
  import { goto } from "$app/navigation";
  import { toast } from "@zerodevx/svelte-toast";
import { postDashboard } from "../../api/dashboard";
  import type { Error } from '../../api/common';
  import DashboardForm from "../../components/DashboardForm.svelte";

  async function submit(event: any) {
    postDashboard({
        name: event.detail.name,
        description: event.detail.description,
    }).then(dashboard => {
      goto(`/dashboard/${dashboard.id}`);
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
    <h3>ダッシュボード作成</h3>
    <DashboardForm label="create" on:submit={submit} />
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
  a {
    color: #666;
  }
</style>
