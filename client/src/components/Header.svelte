<script type=ts>
  export let title: string|null;
  export let user: string;
</script>

<div>
  <h1>{title}</h1>
  {#if user}
    <a href="/user">{user}</a>
  {/if}
</div>

<style>
  div {
    width: 100%;
    background-color: #39f;
    color: #fff;
    display: inline-flex;
  }

  h1 {
    font-size: 2em;
    margin: 15px 20px;
  }

  a {
    display: inline-block;
    margin: auto 15px auto auto;
    text-decoration: none;
    color: #fff;
  }
</style>
