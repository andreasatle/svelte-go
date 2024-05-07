<script>
  import { onMount } from "svelte";

  let pingResponse = "";
  let users = [];

  // fetchPing updates the `pingResponse` variable with the response from the server
  async function fetchPing() {
    const response = await fetch("https://localhost/api/ping");
    if (!response.ok) {
      pingResponse = "Error fetching ping data";
    } else {
      const data = await response.json();
      pingResponse = data.message;
    }
  }
  async function fetchUsers() {
    const response = await fetch("https://localhost/api/users");
    if (!response.ok) {
      users = [];
    } else {
      const data = await response.json();
      console.log(data);
      users = data;
      console.log(users);
    }
  }

  async function fetchAll() {
    const promise1 = fetchPing();
    const promise2 = fetchUsers();
    await Promise.all([promise1, promise2]);
  }
  onMount(fetchAll);
</script>

<h1 class="text-blue-600">Response: {pingResponse}</h1>
{#each users as user}
<p>{user.Name}</p>
{/each}
