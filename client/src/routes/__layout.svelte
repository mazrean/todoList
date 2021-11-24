<script type=ts>
  import Header from '../components/Header.svelte';
  import { SvelteToast, toast } from '@zerodevx/svelte-toast';
  import type { Error } from '../api/common';
  import { getMeAction } from '../store/user';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';

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
    });
  }
</script>

<SvelteToast />
<Header title="ToDo List" />
<slot></slot>
