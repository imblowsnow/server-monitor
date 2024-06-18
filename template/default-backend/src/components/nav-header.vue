<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
  name: "NavHeader"
})
</script>

<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
    <div class="container-fluid">
      <router-link class="navbar-brand" to="/">XXXXX监控系统</router-link>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarScroll" aria-controls="navbarScroll" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarScroll">
        <ul class="navbar-nav me-auto my-2 my-lg-0 navbar-nav-scroll" style="--bs-scroll-height: 100px;">
          <li class="nav-item">
            <router-link class="nav-link" to="/">首页</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/server_group">分组管理</router-link>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
              通知管理
            </a>
            <ul class="dropdown-menu">
              <li><router-link class="dropdown-item" to="/notify_group">通知组配置</router-link></li>
              <li><router-link class="dropdown-item" to="/notify_channel">通知渠道配置</router-link></li>
              <li><hr class="dropdown-divider"></li>
              <li><router-link class="dropdown-item" to="/notify_log">通知日志</router-link></li>
            </ul>
          </li>
        </ul>
        <form class="d-flex" role="search">

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
