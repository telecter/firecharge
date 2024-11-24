<script setup lang="ts">

</script>

<template>
  <header class="container">
    <RouterLink to="/" style="text-decoration: none; color: white;">
      <h2><img src="/firecharge.png" alt="firecharge icon" width="50"> firecharge</h2>
    </RouterLink>
  </header>
  <hr>
  <main class="container">
    <RouterView />
  </main>
</template>

<style>
a {
  text-decoration: none;
}
</style>