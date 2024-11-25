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
    <code class="w-100 mb-1" style="font-size: 15px; display: block;">
        Hello
    </code>
    <input type="text" class="form-control" placeholder="Enter a command here">
    <button type="button" class="btn btn-danger mt-5" data-bs-toggle="modal" data-bs-target="#deleteModal">Delete Server</button>

<div class="modal fade" id="deleteModal" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h1 class="modal-title fs-5" id="exampleModalLabel">Delete this server?</h1>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <p>Are you sure you want to delete this server?</p>
        <p style="color: var(--bs-danger)">This action is irreversible.</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">No!</button>
        <button type="button" class="btn btn-danger">Delete</button>
      </div>
    </div>
  </div>
</div>


</template>