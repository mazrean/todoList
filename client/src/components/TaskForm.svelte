<script type=ts>
  import { createEventDispatcher } from 'svelte';

  export let name = "";
  export let description = "";
  export let label: string;

  const dispatch = createEventDispatcher();
  function submit() {
		dispatch('submit', {
      name,
      description
    });
    name = "";
    description = "";
	}
</script>

<div>
  <p>name</p>
  <input bind:value={name}>
  <p>description</p>
  <textarea bind:value={description} /><br>
  <button on:click={submit}>{label}</button>
</div>

<style>
  div {
    display: inline-block;
    padding: 5px;
  }
  p {
    margin: 0;
  }
  input {
    margin-bottom: 15px;
  }
  button {
    cursor: pointer;
    background-color: transparent;
    color: #222;
    border: 1px solid #e5e5e5;
  }
</style>
