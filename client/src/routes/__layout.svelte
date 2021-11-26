<script type=ts>
  import Header from '../components/Header.svelte';
  import { SvelteToast, toast } from '@zerodevx/svelte-toast';
  import type { Error } from '../api/common';
  import { getMeAction, user } from '../store/user';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { getDashboardsAction } from '../store/dashboard';
  import Sidebar from '../components/Sidebar.svelte';
import { beforeUpdate } from 'svelte';

  let me: string|null = null;
  user.subscribe(user => {
    me = user;
  });

  beforeUpdate(() => {
    if ($page.path !== '/login' && $page.path !== '/signup') {
      getMeAction().catch((err: Error) => {
        if (err.code === 401) {
          toast.push("login required.redirect to login page.");
          goto('/login');
        } else {
          toast.push(`${err.code}:${err.error}`, {
            theme: {
              '--toastBackground': '#F56565',
              '--toastBarBackground': '#C53030'
            }
          });
        }
      }).then(() => {
        getDashboardsAction().catch((err: Error) => {
          toast.push(`${err.code}:${err.error}`, {
            theme: {
              '--toastBackground': '#F56565',
              '--toastBarBackground': '#C53030'
            }
          });
        });
      });
    }
  });
</script>

<SvelteToast />
<div class="body">
  <div class="header">
    <Header title="ToDo List" user={me} />
  </div>
  <div class="main">
    <div class="sidebar">
      <Sidebar />
    </div>
    <div class="slot">
      <slot></slot>
    </div>
  </div>
</div>

<style>
  .body {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    width: 100%;
  }
  .header {
    display: inline-block;
  }
  .main {
    flex: 1;
    display: flex;
    flex-direction: row;
    height: 100%;
  }
  .sidebar {
    width: 240px!important;
    min-height: 100%;
    border-right: 1px #e5e5e5 solid;
  }
  .slot {
    width: 100%;
    padding: 15px;
  }
</style>
