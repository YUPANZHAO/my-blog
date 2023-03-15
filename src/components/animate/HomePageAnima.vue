<template>
  <Transition
    name = "home-page-transition">
    <slot></slot>
  </Transition>
</template>

<script>
export default {

}
</script>

<style>
.home-page-transition-enter-active,
.home-page-transition-leave-active {
  transition: all 1s ease;
}
.home-page-transition-enter-from,
.home-page-transition-leave-to {
  transform: translateY(-600px) !important;
  opacity: 0;
}
</style>