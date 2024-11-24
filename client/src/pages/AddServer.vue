<template>
    <h2>Add a Server</h2>
    <div class="input-group input-group-lg mb-3">
        <span class="input-group-text">Name</span>
        <input type="text" class="form-control" placeholder="Suspicious Server...">
    </div>
    <div class="input-group w-50 mb-3">
        <span class="input-group-text">Version</span>
        <select class="form-control">
            <option value="1.21.3">1.21.3</option>
        </select>
    </div>
    <div class="input-group w-50 mb-3">
        <span class="input-group-text">Type</span>
        <select class="form-control">
            <option value="paper">PaperMC</option>
            <option value="vanilla">Vanilla</option>
        </select>
    </div>
    <RouterLink to="/">
        <button class="btn btn-success">Create!</button>
    </RouterLink>
</template>