<template>
  <div>
    <p class="MenuName">{{ MenuName }}</p>
    <p v-for="item in Items" :key="item" class="Item">
      {{ item.name }}
    </p>
  </div>
</template>

<script>
export default {
    props: [
        'MenuName',
        'Items',
    ],
}
</script>

<style>
.MenuName {
  font-family: Georgia, serif;
  font-weight: bold;
  font-size: 22px;
  margin: 0;
  margin-bottom: 15px;
}
.Item {
  font-family: Georgia, serif;
  font-size: 16px;
  margin: 0;
  margin-bottom: 5px;
  max-width: 260px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>