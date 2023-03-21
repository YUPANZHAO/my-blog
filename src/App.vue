<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <RouterView v-slot="{ Component }" >
    <keep-alive>
      <component :is="Component" :key="$route.name" v-if="$route.meta.keepAlive" />
    </keep-alive>
    <component :is="Component" :key="$route.name" v-if="!$route.meta.keepAlive" />
  </RouterView>
</template>

<style>
html,body {
  width: 100%;
  height: 100%;
}
body {
  margin: 0;
  padding: 0;
}
</style>
