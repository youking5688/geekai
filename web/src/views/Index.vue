<template>
  <div class="index-page" :style="{height: winHeight + 'px'}">
    <div class="menu-box">
      <el-menu mode="horizontal" :ellipsis="false">
        <div class="menu-item">
          <el-image :src="logo" alt="baizeai"/>
          <div class="title">{{ title }}</div>
        </div>
        <div class="menu-item">
          <el-button @click="() => router.push('/login')" style="color:#007bff" :dark="false">
            <i class="iconfont icon-book"></i>
            <span>登录</span>
          </el-button>
        </div>
      </el-menu>
    </div>
    <div class="content">
      <h1>{{ title }}</h1>
      <p>{{ slogan }}</p>
      <el-button @click="() => router.push('/chat')" style="color:#007bff" :dark="false">
        <i class="iconfont icon-chat"></i>
        <span>AI聊天</span>
      </el-button>
      <el-button @click="() => router.push('/mj')" style="color:#424282" :dark="false">
        <i class="iconfont icon-mj"></i>
        <span>AI-MJ绘画</span>
      </el-button>
      <el-button @click="() => router.push('/sd')" style="color:#424282" :dark="false">
        <i class="iconfont icon-sd"></i>
        <span>AI-SD绘画</span>
      </el-button>
      <el-button @click="() => router.push('/xmind')" style="color:#424282" :dark="false">
        <i class="iconfont icon-xmind"></i>
        <span>思维导图</span>
      </el-button>
    </div>
    <div class="footer">
      <footer-bar />
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import FooterBar from '@/components/FooterBar.vue';
import { httpGet } from '@/utils/http';
import { ElMessage } from 'element-plus';

const router = useRouter();

const title = ref("白泽AI 智作系统");
const logo = ref("/images/logo.png");
const slogan = ref("融顶尖技术，绘未来生活");
const winHeight = ref(window.innerHeight - 150);

onMounted(() => {
  httpGet("/api/config/get?key=system")
    .then(res => {
      title.value = res.data.title;
      logo.value = res.data.logo;
    })
    .catch(e => {
      ElMessage.error("获取系统配置失败：" + e.message);
    });
});
</script>

<style lang="stylus" scoped>
@import '@/assets/iconfont/iconfont.css';

.index-page {
  margin: 0;
  background-image: url('@/assets/img/ai-bg.jpg');
  background-size: cover;
  background-position-x: 50%;
  background-position-y: center;
  overflow: hidden;
  color: #ffffff;
  display: flex;
  justify-content: center;
  align-items: baseline;
  padding-top: 150px;

  .menu-box {
    position: absolute;
    top: 0;
    width: 100%;
    display: flex;

    .el-menu {
      padding: 0 30px;
      width: 100%;
      display: flex;
      justify-content: space-between;
      background: none;
      border: none;

      .menu-item {
        display: flex;
        padding: 20px 0;
        color: #ffffff;

        .title {
          font-size: 24px;
          padding: 10px 10px 0 10px;
        }

        .el-image {
          height: 50px;
        }

        .el-button {
          margin-left: 10px;

          span {
            margin-left: 5px;
          }
        }
      }
    }
  }

  .content {
    text-align: center;
    position: relative;

    h1 {
      font-size: 5rem;
      margin-bottom: 1rem;
    }

    p {
      font-size: 1.5rem;
      margin-bottom: 2rem;
    }

    .el-button {
      padding: 25px 20px;
      font-size: 1.3rem;
      transition: all 0.3s ease;

      .iconfont {
        font-size: 1.6rem;
        margin-right: 10px;
      }
    }
  }

  .footer {
    .el-link__inner {
      color: #ffffff;
    }
  }
}
</style>
