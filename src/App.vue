<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import ModelSelector from './components/ModelSelector.vue'
</script>

<template>
  <div class="main">
    <header>
      <nav style="width: 30%;">
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>

      <div style="width: 40%; display: flex; justify-content: center;">
        <ModelSelector />
        <!-- <div>
          center
        </div> -->
      </div>

      <div style="width: 30%; display: flex; justify-content: right;">
        <div>

        </div>
      </div>
    </header>

    <RouterView />
  </div>
</template>

<style scoped>
.main {
  width: 100%;
  padding-left: 2rem;
  padding-right: 2rem;
  padding-top: 0.5rem;
  padding-bottom: 2rem;
  display: flex;
  flex-direction: column;
}

header {
  width: 100%;
  display: flex;
  justify-content: space-between;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}
</style>
