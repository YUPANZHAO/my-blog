<template>
  <div>
    <img :src="AvatarUrl"
      class="avatar" :style="{ width:AvatarSize, height:AvatarSize }"
      @click="BackToHomeHandle">
    <p class="Name"> {{ Name }} </p>
    <div class="MenuGroup" v-if="!isPhone">
      <el-dropdown v-for="menu in Menus" :key="menu" class="MeunItem">
        <span style="outline: none;">
          {{ menu.name }}
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item v-for="item in menu.items" :key="item" @click="item.handle">
              {{ item.name }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <div class="MenuGroup" style="right:20px" v-else>
      <el-dropdown>
        <img src="https://img1.imgtp.com/2023/03/16/xe3gZxGm.png" style="width:25px;height:20px;outline:none;">
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item v-for="menu in Menus" :key="menu" @click="menu.handle">
              <span style="outline: none;">
                {{ menu.name }}
              </span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
export default {
  props: [
    'AvatarUrl',
    'AvatarSize',
    'Name',
    'BackToHomeHandle',
    'Menus',
    'isPhone',
  ],
}
</script>

<style lang="less" scoped>
.avatar {
  border-radius: 50%;
  display: block;
  box-shadow: 0px 0px 5px #595959;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  left: 18px;
  cursor: pointer;
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
.MenuGroup {
  position: absolute;
  width: fit-content;
  display: flex;
  justify-content: center;
  right: 40px;
  top: 50%;
  transform: translateY(-50%);
}
.MeunItem {
  margin-left: 25px;
  font-size: 18px;
  cursor: pointer;
  font-family: Georgia, serif;
  font-weight: bold;
}
</style>