<script>
import ServerStatisticsBar from "@/components/server-statistics-bar.vue";

export default {
  name: "MonitorItem",
  components: {ServerStatisticsBar},
  props: {
    server: {
      type: Object,
      required: true
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
    <router-link :to="'/server/' + server.server_id" style="margin-left: 0px;">
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
        <div style="width: 180px;">
          <div class="wrap" style="padding: 4px 1.25px; width: 100%;" >
            <server-statistics-bar :width="10" :height="10" :total="server.online_statistics"></server-statistics-bar>
          </div>
        </div>
      </div>
    </router-link>
  </div>
</template>

<style scoped>
.monitor-item .box{
  display: flex;
  align-items: center;
}
.monitor-item a{
  color: #000;
  text-decoration: none;
}
</style>
