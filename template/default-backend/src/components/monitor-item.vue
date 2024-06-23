<script>
import ServerStatisticsBar from "@/components/server-statistics-bar.vue";

export default {
  name: "MonitorItem",
  components: {ServerStatisticsBar},
  props: {
    server: {
      type: Object,
      required: true
    },
    width: {
      type: Number,
      default: 10
    },
    height: {
      type: Number,
      default: 10
    }
  },
  data() {
    return {
    }
  },
  mounted() {

  }
}
</script>

<template>
  <div class="monitor-item">
      <div class="box">
        <div class="small-padding" style="flex: 1">
          <div class="info">
            <span class="badge rounded-pill " :class="{'bg-primary' : server.server_status === 1, 'bg-danger' : server.server_status === 0}"
                  title="最近24小时">
                {{ server.online_rate }}%
            </span>
            {{ server.server_name }}
          </div>
        </div>
        <div class="statistics-bar-box">
          <div class="wrap" style="padding: 4px 1.25px; width: 100%;" >
            <server-statistics-bar :width="width" :height="height" :total="server.online_statistics"></server-statistics-bar>
          </div>
          <slot name="statistics-bar-after"></slot>
        </div>
      </div>
  </div>
</template>

<style scoped>
.monitor-item .box{
  display: flex;
  align-items: center;
}

.statistics-bar-box{
  width: 180px;
}
</style>
