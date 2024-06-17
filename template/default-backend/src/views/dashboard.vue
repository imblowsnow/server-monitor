<script lang="ts">
import {defineComponent} from 'vue'
import MonitorFaults from "@/components/monitor-faults.vue";

export default defineComponent({
  name: "dashboard",
  components: {MonitorFaults},
  data() {
    return {
      total: {
        up: 10,
        down: 10,
      },
      faults: [
        { id: 1, name: 'fault1', status: 'up' },
      ],
      clientHeight: document.body.clientHeight,
    }
  },
  mounted() {
    console.log('mounted')
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
              :class="total.up === 0 && 'text-secondary'"
          >
                                    0
                                </span>
        </div>
        <div class="col">
          <h3>离线</h3>
          <span
              class="num"
              :class="total.down > 0 ? 'text-danger' : 'text-secondary'"
          >
                                    0
                                </span>
        </div>
      </div>
    </div>
    <div
        class="shadow-box table-shadow-box"
        style="overflow-x: hidden;"
    >
      <monitor-faults></monitor-faults>
    </div>
  </div>
</template>

<style scoped>

</style>
