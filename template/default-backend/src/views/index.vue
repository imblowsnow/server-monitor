<script>
import FrontendMonitorList from "@/components/frontend-monitor-list.vue";
import {apiDashboardTotal} from "@/api/monitor.js";

export default {
  name: "index",
  components: {FrontendMonitorList},
  data() {
    return {
      total: {
        onlineNum: 0,
        offlineNum: 0,
      },
    }
  },
  mounted() {
    this.dashboardTotal();
  },
  methods: {
    dashboardTotal() {
      apiDashboardTotal().then(data => {
        this.total = data
      });
    }
  }
}
</script>

<template>
  <div class="home container mt-3">
    <h1 class="mb-3">状态速览</h1>
    <div class="shadow-box big-padding text-center mb-4">
      <div class="row">
        <div class="col">
          <h3>在线</h3>
          <span
              class="num"
              :class="total.onlineNum === 0 && 'text-secondary'"
          >
              {{ total.onlineNum}}
          </span>
        </div>
        <div class="col">
          <h3>离线</h3>
          <span
              class="num"
              :class="total.offlineNum > 0 ? 'text-danger' : 'text-secondary'"
          >
             {{ total.offlineNum}}
          </span>
        </div>
      </div>
    </div>

    <div class="shadow-box mb-3">
      <frontend-monitor-list></frontend-monitor-list>
    </div>


    <footer class="text-center m-5">
      技术支持：xxx
    </footer>
  </div>

</template>

<style>
.home .num {
  font-size: 30px;
  color: var(--theme-color);
  font-weight: bold;
  display: block;
}
.statistics-bar-box{
  width: 220px!important;
}
@media (max-width: 800px) {
  .statistics-bar-box{
    width: 115px!important;
  }
}
</style>
