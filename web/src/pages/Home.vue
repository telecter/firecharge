<script setup lang="ts">
import ServerListing from '../components/ServerListing.vue';

const servers = await (await fetch("http://localhost:3000/api/servers")).json()
</script>

<template>
    <h2>Servers</h2>
    <RouterLink to="/add">
        <button class="btn btn-primary">Add</button>
    </RouterLink>
    
    <ServerListing :id="server.name" :type="server.type" :version="server.version" v-for="server in servers" />
</template>