<script type=ts>
  import type { Item } from "./Sidebar";
  export let title: string;
  export let items: Item[];
</script>

<div>
  <h3>{title}</h3>
  <ul>
    {#each items as item}
      <li>
        <a href="{item.link}">{item.label}</a>
      </li>
    {/each}
  </ul>
</div>

<style>
  h3 {
    font-size: 24px;
    line-height: 1.4;
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
    text-decoration: none;
  }
</style>
