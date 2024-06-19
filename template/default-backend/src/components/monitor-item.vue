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
      <div class="row">
        <div class="col-8 small-padding">
          <div class="info">
            <span class="badge rounded-pill bg-primary" title="最近24小时">
                {{ server.online_rate }}%
            </span>
            {{ server.server_name }}
          </div>
        </div>
        <div class="col-4">
          <div class="wrap" style="padding: 4px 1.25px; width: 100%;" >
            <server-statistics-bar :max-num="10"
                :total="server.online_statistics"></server-statistics-bar>
          </div>
        </div>
      </div>
    </router-link>
  </div>
</template>

<style scoped>
.monitor-item a{
  color: #000;
  text-decoration: none;
}
</style>
