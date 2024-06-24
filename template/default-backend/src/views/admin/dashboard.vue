<script>
import {defineComponent} from 'vue'
import MonitorFaults from "@/components/monitor-faults.vue";
import {dashboardTotal} from "@/api/backend/monitor.js";

export default defineComponent({
  name: "dashboard",
  components: {MonitorFaults},
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
      dashboardTotal().then(data => {
        this.total = data
      });
    }
  }
})
</script>

<template>
  <div class="dashboard">
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
    <div
        class="shadow-box table-shadow-box"
        style="overflow-x: hidden;"
    >
      <monitor-faults :limit="5"></monitor-faults>
    </div>
  </div>
</template>

<style scoped>
.dashboard .num {
  font-size: 30px;
  color: var(--theme-color);
  font-weight: bold;
  display: block;
}
</style>
