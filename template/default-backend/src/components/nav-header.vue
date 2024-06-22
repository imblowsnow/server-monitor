<script>
import {defineComponent} from 'vue'
import {useSettingStore} from "@/stores/setting.js";
import {useUserStore} from "@/stores/user.js";

export default defineComponent({
  name: "NavHeader",
  props: {
    login: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      settingStore: useSettingStore(),
      userStore: useUserStore(),
      title: 'Server Monitor',
    }
  },
  watch: {
    'settingStore.setting': {
      handler() {
        console.log('setting changed', this.settingStore.setting);
        this.title = this.settingStore.setting.site_name;
      },
      deep: true
    }
  },
  mounted() {
    this.title = this.settingStore.setting.site_name || 'Server Monitor';
  },
  methods: {
    logout() {
      this.userStore.logout();
      this.$router.push('/login');
    }
  }
})
</script>

<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
    <div class="container-fluid">
      <router-link class="navbar-brand" to="/">{{ title }}</router-link>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarScroll" aria-controls="navbarScroll" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div v-if="login === false" class="collapse navbar-collapse" id="navbarScroll">
        <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
          <li class="nav-item">
            <router-link class="nav-link" to="/">首页</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/server_group">分组管理</router-link>
          </li>
<!--          <li class="nav-item dropdown">-->
<!--            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">-->
<!--              通知管理-->
<!--            </a>-->
<!--            <ul class="dropdown-menu">-->
<!--              <li><router-link class="dropdown-item" to="/notify_group">通知组配置</router-link></li>-->
<!--              <li><router-link class="dropdown-item" to="/notify_channel">通知渠道配置</router-link></li>-->
<!--              <li><hr class="dropdown-divider"></li>-->
<!--              <li><router-link class="dropdown-item" to="/notify_log">通知日志</router-link></li>-->
<!--            </ul>-->
<!--          </li>-->
          <li class="nav-item">
            <router-link class="nav-link" to="/setting">系统设置</router-link>
          </li>
        </ul>
        <form class="d-flex" role="search">
          <button class="btn btn-outline-success" type="submit" @click="logout">退出</button>
        </form>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.navbar .navbar-collapse a.router-link-exact-active {
  color: var(--theme-color);
  font-weight: bold;
}
</style>
