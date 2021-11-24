<script type=ts>
  import { goto } from "$app/navigation";
  import type { Item } from "./Sidebar";

  export let title: string;
  export let items: Item[];
</script>

<div>
  <h3>{title}</h3>
  <button on:click={()=>goto("/dashboard/new")}>Create Dashboard</button>
  {#if items.length == 0}
  <p>No items</p>
  {:else}
  <ul>
    {#each items as item}
      <li>
        <a href="{item.link}">{item.label}</a>
      </li>
    {/each}
  </ul>
  {/if}
</div>

<style>
  h3 {
    font-size: 24px;
    line-height: 1.4;
    color: #222;
    margin-bottom: 5px;
  }
  ul {
    list-style: none;
    padding: 0;
  }
  li {
    margin: 5px 0;
  }
  a {
    color: #333;
  }
  button {
    cursor: pointer;
    background-color: transparent;
    color: #222;
    border: 1px solid #e5e5e5;
  }
</style>
