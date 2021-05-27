<template>
  <div class="card">
    <slot name="content"></slot>
  </div>
</template>

<style scoped>
.card {
  height: fit-content;
  width: fit-content;
  border-radius: 2px;
  box-shadow: -2px 2px 4px grey;
  padding: 6px 12px;
}
</style>
