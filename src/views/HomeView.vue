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
          <MenuItem v-for="menu in menuItems" :key="menu" 
            :class="classMenuItem"
            :MenuName="menu.name"
            :Items="menu.items"
            />
        </div>
      </div>
    </home-page-anima>
    <navigation-anima>
      <Navigation class="Navigation" v-if="!isSingleHomePage"
        AvatarUrl="https://img1.imgtp.com/2023/03/15/Yu7q4P8x.jpg"
        AvatarSize="36px"
        Name="YUPAN ZHAO' BLOG"
        :BackToHomeHandle=backToHome
        :Menus="menuItems"
        :isPhone="isPhone" />
    </navigation-anima>
    <content-page-anima>
      <div class="Content" v-show="!isSingleHomePage">
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
import { getRecentArticle } from '../api/getRecentArticle.js'
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
      menuItems: [
        {
          name: "导航",
          items: [
            {
              name: "首页",
              handle: () => {
                this.isSingleHomePage = true;
              },
            },
            {
              name: "文章列表",
              handle: () => {
                this.isSingleHomePage = false;
              },
            },
          ],
          handle: () => { this.isSingleHomePage = true; },
        },
        {
          name: "最近文章",
          items: this.recentArticle(),
          handle: () => { this.returnTop() },
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
  },
  methods: {
    windowSizeChange() {
      this.screenWidth = window.innerWidth;
      this.screenHeight = window.innerHeight;
    },
    backToHome() {
      this.isSingleHomePage = true;
    },
    returnTop() {
      document.body.scrollTop = 0
      // firefox
      document.documentElement.scrollTop = 0
    },
    async recentArticle(max_num) {
      var reply = await getRecentArticle({
        "max_num": max_num,
      })
      var ret = []
      if(reply.status == 200) {
        for (let i = 0; i < reply.data.recent_articles.length; i++) { 
          ret.push({
            name: reply.data.recent_articles[i].title,
            handle: () => {
              this.$router.push( { 
                name: 'paper', 
                query: {
                  paperName: reply.data.recent_articles[i].name,
                },
              });
            },
          })
        }
      }
      this.menuItems[1].items = ret
      return ret
    },
  },
  activated() {
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
  deactivated() {
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