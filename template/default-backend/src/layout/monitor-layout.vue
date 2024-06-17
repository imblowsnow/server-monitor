<script lang="ts">
import {defineComponent} from 'vue'
import NavHeader from '@/components/nav-header.vue'
import NavFooter from '@/components/nav-footer.vue'
import MonitorList from "@/components/monitor-list.vue";
export default defineComponent({
  name: "MonitorLayout",
  components: {
    NavHeader,
    NavFooter,
    MonitorList
  },
  data(){
    return {
      // 宽度小于 800
      isMobile: document.body.clientWidth < 1000
    }
  },
  mounted() {
    // 监听窗口大小变化
    window.onresize = () => {
      this.isMobile = document.body.clientWidth < 1000
      console.log('isMobile',this.isMobile);
    }
  }
})
</script>

<template>
  <div class="monitor-layout" :class="{'mobile': isMobile}">
    <nav-header></nav-header>
    <div v-if="!isMobile" class="row" style="width: 100%;padding: 20px;">
      <div class="col-4">
        <div class="shadow-box mb-3" style="height: calc(-160px + 100vh);">
          <slot name="monitor-list">
            <div>
              <router-link to="/server/add" class="btn btn-primary mb-3">
                <svg class="svg-inline--fa fa-plus fa-w-14" style="width: 0.875em;"
                     aria-hidden="true" focusable="false" data-prefix="fas" data-icon="plus" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path class="" fill="currentColor" d="M416 208H272V64c0-17.67-14.33-32-32-32h-32c-17.67 0-32 14.33-32 32v144H32c-17.67 0-32 14.33-32 32v32c0 17.67 14.33 32 32 32h144v144c0 17.67 14.33 32 32 32h32c17.67 0 32-14.33 32-32V304h144c17.67 0 32-14.33 32-32v-32c0-17.67-14.33-32-32-32z"></path></svg>
                添加服务器
              </router-link>
            </div>
            <monitor-list style="height: calc(100% - 107px);" class="scrollbar"></monitor-list>
          </slot>
        </div>
      </div>
      <div class="col-8 mt-3">
        <router-view></router-view>
      </div>
    </div>
    <div v-else style="width: 100%">
      <div class="m-3" style="margin-bottom: 70px!important;">
        <router-view></router-view>
      </div>
    </div>
    <nav-footer v-if="isMobile"></nav-footer>
  </div>

</template>

<style scoped>

</style>
