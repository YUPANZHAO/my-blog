<template>
  <Transition
    name = "content-page-transition">
    <slot></slot>
  </Transition>
</template>

<script>
export default {

}
</script>

<style>
.content-page-transition-enter-active {
  transition: all 2s ease;
}
.content-page-transition-leave-active {
  transition: all 1s ease;
}
.content-page-transition-enter-from,
.content-page-transition-leave-to {
  transform: translateY(800px) !important;
  opacity: 0;
}
</style>