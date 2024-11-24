<script setup>
    const type = "PaperMC";
    const version = "1.21.3";
</script>


<template>
    <h2>{{ $route.params.id }}</h2>
    <p class="lead"><strong>Version: </strong>{{ version }} ({{ type }})</p>
    <p>Currently <span style="color: red">offline.</span></p>
    <div class="btn-group" role="group">
        <button class="btn btn-success">Start</button>
        <button class="btn btn-danger">Stop</button>
    </div>
    <hr>
    <h3>Console</h3>
    <code class="w-100 mb-1"style="font-size: 15px; display: block;">
        Hello
    </code>
    <input type="text" class="form-control" placeholder="Enter a command here">
</template>