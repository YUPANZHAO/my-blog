<template>
  <Transition
    name = "navigation-transition">
    <slot></slot>
  </Transition>
</template>

<script>
export default {

}
</script>

<style>
.navigation-transition-enter-active,
.navigation-transition-leave-active {
  transition: all 0.3s ease-out;
}
.navigation-transition-enter-from,
.navigation-transition-leave-to {
  transform: translateY(-54px) !important;
  opacity: 0;
}
</style>