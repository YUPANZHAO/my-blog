<template>
  <div>
    <img :src="AvatarUrl"
      class="avatar" :style="{ width:AvatarSize, height:AvatarSize }">
    <p class="Name"> {{ Name }} </p>
    <p class="Description"> {{ Description }} </p>
    <ul class="IconGroup" >
      <li v-for="icon in IconGroup" :key="icon">
        <a :href="icon.link" target="_blank" :hint="icon.hint" class="tooltip">
          <img class="Icon" :src="icon.url">
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: [
    'AvatarUrl',
    'AvatarSize',
    'Name',
    'Description',
    'IconGroup',
  ],
}
</script>

<style scoped>
.avatar {
  border-radius: 50%;
  display: block;
  margin: 0 auto;
  box-shadow: 0px 0px 10px #888888;
}
.Name {
  font-family: Georgia, serif; 
  font-size: 22px;
  width: fit-content;
  margin: 0 auto;
  margin-top: 10px;
}
.Description {
  font-family: Georgia, serif;
  color: gray;
  font-size: 14px;
  width: fit-content;
  margin: 0 auto; 
  margin-top: 10px;
}
.IconGroup {
  display: flex;
  justify-content: center;
  List-style-type: none;
  margin: 0;
  padding: 0;
  margin-top: 10px;
  height: 20px;
}
.Icon {
  width: 20px;
  height: 20px;
  padding-inline: 5px;
}
.tooltip {
  position: relative;
}
.tooltip:hover::before {
  word-break: keep-all;
  white-space: nowrap;
  content: attr(hint);
  position: absolute;
  padding: 2px 6px;
  display: block;
  color: #333;
  border: 1px solid #333;
  border-radius: 5px;
  font-size: 12px;
  line-height: 20px;
  top: 25px;
  left: 50%;
  transform: translateX(-50%);
  box-shadow: 0px 0px 10px #eaeaea;
}
</style>