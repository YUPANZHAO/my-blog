<template>
  <div>
    <p class="MenuName">{{ MenuName }}</p>
    <a v-for="item in Items" :key="item" class="Item" @click="item.handle">
      <p class="ItemText"> {{ item.name }} </p> 
    </a>
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
}
.ItemText {
  font-family: Georgia, serif;
  font-size: 16px;
  margin: 0;
  margin-bottom: 10px;
  max-width: 260px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.ItemText:hover {
  color: rgb(0, 196, 23);
  background-color: rgba(171, 171, 171, 0.069);
  border-radius: 5px;
}
</style>