<template>
  <div class="root">
    <home-page-anima>
      <div class="root-SingleHomePage" v-if="isSingleHomePage">
        <NameCard :class="classNameCard" 
          AvatarUrl="https://img1.imgtp.com/2023/03/15/Yu7q4P8x.jpg"
          AvatarSize="100px"
          Name="YUPAN ZHAO"
          Description="Huizhou University"
          :IconGroup="iconGroup" />
        <div :class="classMenu">
          <MenuItem :class="classMenuItem"
            MenuName="最近文章1"
            :Items="menuItems_One" />
          <MenuItem :class="classMenuItem"
            MenuName="最近文章2"
            :Items="menuItems_Two" />
          <MenuItem :class="classMenuItem"
            MenuName="最近文章3"
            :Items="menuItems_Three" />
          <MenuItem :class="classMenuItem"
            MenuName="最近文章4"
            :Items="menuItems_Four" />
        </div>
      </div>
    </home-page-anima>
    <navigation-anima>
      <Navigation class="Navigation" v-if="!isSingleHomePage"
        AvatarUrl="https://img1.imgtp.com/2023/03/15/Yu7q4P8x.jpg"
        AvatarSize="36px"
        Name="YUPAN ZHAO' BLOG"
        :BackToHomeHandle=backToHome
        :Menus="allMenuItems"
        :isPhone="isPhone" />
    </navigation-anima>
    <content-page-anima>
      <div class="Content" v-if="!isSingleHomePage">
        <PaperList class="PaperList">

        </PaperList>
      </div>
    </content-page-anima>
  </div>
</template>

<script>
import NameCard from '../components/NameCard.vue'
import MenuItem from '../components/MenuItem.vue'
import HomePageAnima from '../components/animate/HomePageAnima.vue'
import Navigation from '../components/Navigation.vue'
import NavigationAnima from '../components/animate/NavigationAnima.vue'
import PaperList from '../components/PaperList.vue'
import ContentPageAnima from '../components/animate/ContentPageAnima.vue'
export default {
  components: {
    NameCard,
    MenuItem,
    HomePageAnima,
    Navigation,
    NavigationAnima,
    PaperList,
    ContentPageAnima,
  },
  data() {
    return {
      screenWidth: window.innerWidth,
      screenHeight: window.innerHeight,
      iconGroup: [
        { 
          url: "https://img1.imgtp.com/2023/03/15/6OB2Q7Go.png",
          link: "https://github.com/YUPANZHAO",
          hint: "GitHub",
        },
        { 
          url: "https://img1.imgtp.com/2023/03/15/CCetn0ZT.png",
          hint: "yupanzhao2022@163.com",
        },
        { 
          url: "https://img1.imgtp.com/2023/03/15/PK67VI4A.png", 
          link: "http://47.115.209.240:8080",
          hint: "个人博客",
        },
      ],
      menuItems_One: [
        {
          name: "基于 RTMP 协议实现的家庭监控系统",
          handle: () => {
            this.isSingleHomePage = !this.isSingleHomePage;
          },
        },
        {
          name: "C++ 学习笔记",
          handle: () => {
            console.log("handle2");
          },
        },
        {
          name: "三国杀华容道",
        },
        {
          name: "📝记账APP",
        },
      ],
      menuItems_Two: [
        {
          name: "基于 RTMP 协议实现的家庭监控系统",
          handle: () => {
            this.isSingleHomePage = !this.isSingleHomePage;
          },
        },
        {
          name: "C++ 学习笔记",
          handle: () => {
            console.log("handle2");
          },
        },
        {
          name: "三国杀华容道",
        },
        {
          name: "📝记账APP",
        },
      ],
      menuItems_Three: [
        {
          name: "基于 RTMP 协议实现的家庭监控系统",
          handle: () => {
            this.isSingleHomePage = !this.isSingleHomePage;
          },
        },
        {
          name: "C++ 学习笔记",
          handle: () => {
            console.log("handle2");
          },
        },
        {
          name: "三国杀华容道",
        },
        {
          name: "📝记账APP",
        },
      ],
      menuItems_Four: [
        {
          name: "基于 RTMP 协议实现的家庭监控系统",
          handle: () => {
            this.isSingleHomePage = !this.isSingleHomePage;
          },
        },
        {
          name: "C++ 学习笔记",
          handle: () => {
            console.log("handle2");
          },
        },
        {
          name: "三国杀华容道",
        },
        {
          name: "📝记账APP",
        },
      ],
      isSingleHomePage: true,
    }
  },
  computed: {
    isPhone() {
      return this.screenHeight / this.screenWidth > 0.6 || this.screenWidth <= 850;
    },
    classNameCard() {
      return this.isPhone ? 'NameCard-Phone' : 'NameCard';
    },
    classMenu() {
      return this.isPhone ? 'Menu-Phone' : 'Menu';
    },
    classMenuItem() {
      return this.isPhone ? 'MenuItem-Phone' : 'MenuItem';
    },
    allMenuItems() {
      return [
        {
          name: "最近文章1",
          items: this.menuItems_One,
          handle: () => { this.isSingleHomePage = true; },
        },
        {
          name: "最近文章2",
          items: this.menuItems_Two,
          handle: () => { console.log("clicked"); }
        },
        {
          name: "最近文章3",
          items: this.menuItems_Three,
        },
        {
          name: "最近文章4",
          items: this.menuItems_Four,
        },
      ];
    },
  },
  methods: {
    windowSizeChange() {
      this.screenWidth = window.innerWidth;
      this.screenHeight = window.innerHeight;
    },
    backToHome() {
      this.isSingleHomePage = true;
    }
  },
  mounted() {
    this.windowSizeChange();
    window.onresize = () => {
      if(!this.timer) {
        this.timer = true;
        let that = this;
        setTimeout(function() {
          that.windowSizeChange();
          that.timer = false;
        }, 400);
      }
    };
  },
  destroyed() {
    window.onresize = null;
  },
}
</script>

<style scoped>
.root-SingleHomePage {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  position: fixed;
  top: 0px;
  overflow: auto;
}
.root {
  width: 100%;
  height: 100%;
}
.NameCard-Phone {
  width: fit-content;
  margin: 0 auto;
  position: absolute;
  top: 10%;
}
.NameCard {
  width: fit-content;
  height: fit-content;
  position: relative;
  top: 50%;
  transform: translateY(-50%);
}
.Menu-Phone {
  width: fit-content;
  position: absolute;
  margin: 0 auto;
  top: calc(10% + 240px);
  margin-bottom: 200px;
}
.Menu {
  width: fit-content;
  height: fit-content;
  display: grid;
  grid-auto-flow: column;
  grid-template-rows: 1fr 1fr;
  grid-row-gap: 30px;
  grid-column-gap: 30px;
  position: relative;
  top: 50%;
  transform: translateY(-50%);
  margin-left: 150px;
}
.MenuItem-Phone {
  margin-bottom: 30px;
}
.MenuItem {
  width: fit-content;
}
.Navigation {
  width: 100%;
  height: 54px;
  position: fixed;
  top: 0px;
  background-color: white;
  box-shadow: 0px 0px 10px #c9c9c9;
  z-index: 1;
}
.Content {
  margin-top: 54px;
  width: calc(100% - 2px);
  position: relative;
  padding: 1px;
}
.PaperList {
  width: 600px;
  max-width: 86%;
  margin-left: auto;
  margin-right: auto;
  position: relative;
  margin-top: 30px;
  margin-bottom: 30px;
}
</style>