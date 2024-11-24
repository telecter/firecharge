<script setup lang="ts">
import ServerListing from '../components/ServerListing.vue';
</script>

<template>
    <h2>Your servers</h2>
    <RouterLink to="/add">
        <button class="btn btn-primary">Add</button>
    </RouterLink>
    
    <ServerListing id="snapd" />
</template>