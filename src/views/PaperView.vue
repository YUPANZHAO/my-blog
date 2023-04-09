<template>
  <div class="root">
    <Navigation class="Navigation"
      AvatarUrl="https://img1.imgtp.com/2023/03/15/Yu7q4P8x.jpg"
      AvatarSize="36px"
      Name="YUPAN ZHAO' BLOG"
      :BackToHomeHandle=backToHome
      :Menus="menuItems"
      :isPhone="isPhone" />
    <div class="Content">
      <div :class="ClassPaper">
        <div class="CoverAndTitle">
          <div style="line-height:0;font-size:0;">
            <img class="Cover" :src="paperData.cover_url">
          </div>
          <div class="Gradient" v-if="paperData.cover_url != ''" />
          <div :class="ClassInfoGroup">
            <div class="IconGroup" v-if="paperData.cover_url != ''">
              <div style="line-height:0;font-size:0;">
                <img class="DateIcon" src="https://img1.imgtp.com/2023/03/21/Wf3rRBJP.png" >
              </div>
              <p class="Date"> {{ paperData.date }} </p>
              <div style="line-height:0;font-size:0;">
                <img class="CategoryIcon" src="https://img1.imgtp.com/2023/03/21/n6SdmknJ.png" >
              </div>
              <p class="Category"> {{ paperData.category }} </p>
            </div>
            <p :class="ClassTitle"> {{ paperData.title }} </p>
            <div class="IconGroup-NoCover" v-if="paperData.cover_url == ''">
              <div style="line-height:0;font-size:0;">
                <img class="DateIcon" src="https://img1.imgtp.com/2023/03/19/1x4xkvM5.png" >
              </div>
              <p class="Date-NoCover"> {{ paperData.date }} </p>
              <div style="line-height:0;font-size:0;">
                <img class="CategoryIcon" src="https://img1.imgtp.com/2023/03/19/DIF6J9jb.png" >
              </div>
              <p class="Category-NoCover"> {{ paperData.category }} </p>
            </div>
          </div>
        </div>
        <v-md-preview class="MarkDownContent" :text="paperData.content" @copy-code-success="codeCopySuccess"></v-md-preview>
      </div>
    </div>
  </div>
</template>

<script>
import { getArticle } from '../api/getArticle.js'
export default {
  data() {
    return {
      screenWidth: window.innerWidth,
      screenHeight: window.innerHeight,
      menuItems: [
        {
          name: "导航",
          items: [
            {
              name: "回到首页",
              handle: () => {
                this.$router.replace({ path: '/' });
              },
            },
          ],
          handle: () => { this.$router.replace({ path: '/' }); },
        },
      ],
      paperData: {
        title: '',
        date: '',
        category: '',
        cover_url: '',
        content: '',
      },
    }
  },
  computed: {
    isPhone() {
      return this.screenHeight / this.screenWidth > 0.6 || this.screenWidth <= 850;
    },
    ClassTitle() {
      return this.paperData.cover_url == '' ? 'Title-NoCover' : 'Title';
    },
    ClassPaper() {
      return this.isPhone ? 'Paper-Phone' : 'Paper';
    },
    ClassInfoGroup() {
      return this.paperData.cover_url == '' ? 'InfoGroup-NoCover' : 'InfoGroup';
    },
  },
  methods: {
    windowSizeChange() {
      this.screenWidth = window.innerWidth;
      this.screenHeight = window.innerHeight;
    },
    backToHome() {
      if(window.history.state.back) {
        this.$router.go(-1);
      }else {
        this.$router.replace({ path: '/' });
      }
    },
    async initData() {
      var reply = await getArticle({
        "name": this.$route.query.paperName
      })
      if(reply.status == 200) {
        this.paperData = reply.data
      }
    },
    codeCopySuccess(code) {
      ElMessage({
        message: '代码复制成功 🍻',
        type: 'success',
        grouping: true,
      })
    },
  },
  mounted() {
    window.scrollTo(0,0);
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
    this.initData();
  },
  destroyed() {
    window.onresize = null;
  },
}
</script>

<style scoped>
.root {
  width: 100%;
  height: 100%;
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
  top: 54px;
  width: 100%;
  position: absolute;
}
.Paper {
  width: 800px;
  max-width: 96%;
  background-color: white;
  box-shadow: 0px 0px 10px #d8d8d8;
  margin: 0 auto;
  margin-top: 20px;
  margin-bottom: 20px;
}
.Paper-Phone {
  width: 100%;
  background-color: white;
}
.CoverAndTitle {
  width: 100%;
  position: relative;
}
.Cover {
  width: 100%;
}
.Gradient {
  width: 100%;
  height: 25vh;
  background: linear-gradient(#00000000, #373737);
  position: absolute;
  transform: translateY(-100%);
}
.Title {
  margin: 0;
  margin-left: 25px;
  margin-right: 25px;
  padding-bottom: 20px;
  font-size: 5vmin;
  color: white;
}
.Title-NoCover {
  margin: 0;
  margin-left: 25px;
  margin-right: 25px;
  padding-top: 30px;
  padding-bottom: 30px;
  font-size: 5vmin;
}
.InfoGroup {
  width: 100%;
  position: absolute;
  transform: translateY(-100%);
}
.InfoGroup-NoCover {
  width: 100%;
}
.IconGroup {
  width: 80%;
  height: 26px;
  margin-bottom: 5px;
  margin-left: 25px;
  display: flex;  
}
.IconGroup-NoCover {
  width: 80%;
  height: 26px;
  margin-left: 25px;
  display: flex;
}
.DateIcon {
  width: 26px;
  height: 26px;
}
.CategoryIcon {
  width: 26px;
  height: 26px;
  margin-left: 15px;
  position: relative;
  top: -1px;
}
.Date {
  margin: 0;
  padding: 0;
  position: relative;
  top: 2.5px;
  margin-left: 5px;
  font-size: 16px;
  color: white;
}
.Category {
  margin: 0;
  padding: 0;
  position: relative;
  top: 2px;
  margin-left: 5px;
  font-size: 16px;
  color: white;
}
.Date-NoCover {
  margin: 0;
  padding: 0;
  position: relative;
  top: 2.5px;
  margin-left: 5px;
  font-size: 16px;
  color: #5a5a5a;
}
.Category-NoCover {
  margin: 0;
  padding: 0;
  position: relative;
  top: 2px;
  margin-left: 5px;
  font-size: 16px;
  color: #5a5a5a;
}
.MarkDownContent {
  width: 100%;
  background-color: white;
  position: relative;
}
</style>