<template>
  <div>
    <img :src="AvatarUrl"
      class="avatar" :style="{ width:AvatarSize, height:AvatarSize }"
      @click="BackToHomeHandle">
    <p class="Name"> {{ Name }} </p>
  </div>
</template>

<script>
export default {
  props: [
    'AvatarUrl',
    'AvatarSize',
    'Name',
    'BackToHomeHandle',
  ],
}
</script>

<style scoped>
.avatar {
  border-radius: 50%;
  display: block;
  box-shadow: 0px 0px 5px #595959;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  left: 18px;
}
.Name {
  font-family: Georgia, serif; 
  font-size: 22px;
  width: fit-content;
  margin: 0;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  left: 68px;
}
</style>