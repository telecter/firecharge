<script setup lang="ts">


const props = defineProps<{ id: string }>()

</script>

<template>
  <div class="card mt-2 p-2">
    <h2><RouterLink :to="{ name: 'server', params: { id: props.id }}">{{ props.id }}</RouterLink></h2>
    <p class="lead">1.21.3
      <br>
      <strong>PaperMC</strong>
    </p>
  </div>
</template>


