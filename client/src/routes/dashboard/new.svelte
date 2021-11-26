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

<div class="container">
  <h3>New Dashboard</h3>
  <DashboardForm label="create" on:submit={submit} />
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
    margin: 0 5px;
  }
</style>
