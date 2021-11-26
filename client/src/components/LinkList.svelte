<script type=ts>
  import type { Item } from "./LinkList";

  export let items: Item[];
</script>

<div>
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
  ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  li {
    margin: 5px 0;
  }
  a:hover {
    color: #666;
  }
  a {
    text-decoration: none;
    color: #999;
  }
</style>
