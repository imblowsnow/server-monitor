<script lang="ts">
import {defineComponent} from 'vue'

export default defineComponent({
  name: "NavFooter"
})
</script>

<template>
  <nav class="bottom-nav">
    <router-link to="/" class="nav-link">
      <div><font-awesome-icon icon="tachometer-alt" /></div>
      首页
    </router-link>

    <router-link to="/monitor" class="nav-link">
      <div><font-awesome-icon icon="list" /></div>
      列表
    </router-link>

    <!--    <a href="/add" class="nav-link">-->
    <!--        <div><font-awesome-icon icon="plus" /></div>-->
    <!--        Add-->
    <!--    </a>-->

    <!--    <a href="/settings" class="nav-link">-->
    <!--        <div><font-awesome-icon icon="cog" /></div>-->
    <!--        Settings-->
    <!--    </a>-->
  </nav>
</template>

<style scoped>

.bottom-nav {
  z-index: 1000;
  position: fixed;
  bottom: 0;
  height: calc(50px + env(safe-area-inset-bottom));
  line-height: calc(50px + env(safe-area-inset-bottom));
  width: 100%;
  left: 0;
  background-color: #fff;
  box-shadow: 0 15px 47px 0 rgba(0, 0, 0, 0.05), 0 5px 14px 0 rgba(0, 0, 0, 0.05);
  text-align: center;
  white-space: nowrap;
}
.bottom-nav a{
  text-align: center;
  width: 25%;
  display: inline-block;
  height: 100%;
  font-size: 13px;
  color: #c1c1c1;
  overflow: hidden;
  text-decoration: none;
}
.bottom-nav a.router-link-exact-active, .bottom-nav a.active {
  color: var(--theme-color);
  font-weight: bold;
}
.bottom-nav a div{
  font-size: 20px;
}
</style>
