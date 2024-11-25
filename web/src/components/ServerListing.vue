<script setup lang="ts">


const props = defineProps<{ id: string, type: string, version: string }>()

</script>

<template>
  <div class="card mt-2 p-2">
    <h2><RouterLink :to="{ name: 'server', params: { id: props.id }}">{{ props.id }}</RouterLink></h2>
    <p class="lead">{{ version }}
      <br>
      <strong>{{ type }}</strong>
    </p>
  </div>
</template>


