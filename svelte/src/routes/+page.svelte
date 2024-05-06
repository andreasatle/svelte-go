<script>
  import { onMount } from "svelte";

  let message;

  onMount(async () => {
    //try {
    //  const response = await fetch("http://localhost:8080/ping2");
    //  data = await response.json();
    //} catch (error) {
    //  data = { message: "Error fetching data" };
    //}
    const response = await fetch("https://localhost/api/ping");
    if (!response.ok) {
      message = "Error fetching data";
    } else {
      const data = await response.json();
      message = data.message;
    }
    console.log(message);
  });
</script>

<h1>Response: {message}</h1>
