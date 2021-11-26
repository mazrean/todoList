<script type=ts>
  import { createEventDispatcher } from 'svelte';

  export let name = "";
  export let password = "";
  export let label: string;

  const dispatch = createEventDispatcher();
  function submit() {
		dispatch('submit', {
      name,
      password
    });
	}
</script>

<div>
  <p>name</p>
  <input bind:value={name}>
  <p>password</p>
  <input type="password" bind:value={password}><br>
  <button on:click={submit}>{label}</button>
</div>

<style>
  div {
    display: inline-block;
  }
  p {
    margin: 0;
  }
  input {
    margin-bottom: 15px;
  }
  button {
    cursor: pointer;
    background-color: #1e87f0;
    color: #fff;
    border: 1px solid transparent;
  }
</style>
