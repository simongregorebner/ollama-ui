<script setup lang="ts">
import ChatBox from '../components/ChatBox.vue'
import ModelSelector from '../components/ModelSelector.vue'
</script>

<template>
  <main>
    <div class="flex-grid">
      <div class="main-col">
        <ChatBox />
      </div>
      <div class="right-col">
        <ModelSelector />
      </div>
    </div>
  </main>
</template>

<style scoped>
.flex-grid {
  display: flex;
  height: 100%;
}

.main-col {
  width: 75%;
}

.right-col {
  width: 25%;
}
</style>