<script>
import MonitorItem from '@/components/monitor-item.vue'
import {apiMonitorGroups} from "@/api/backend/monitor.js";
export default {
  name: "FrontendMonitorList",
  components: {
    MonitorItem
  },
  data() {
    return {
      activeGroupIds: [1],
      groups: [],

      activeServerId: null
    }
  },
  mounted() {
    this.getMonitorGroups()
    // 监听事件
    this.$bus.on('serverChanged', () => {
      this.getMonitorGroups()
    })
    this.$bus.on('serverSelect', (serverId) => {
      console.log('serverSelect', serverId);
      this.activeServerId = serverId;
      if (!this.activeServerId) return;
      this.selectGroupByServerId(this.activeServerId)
    })
  },
  methods:{
    getMonitorGroups() {
      apiMonitorGroups().then(data => {
        this.groups = data
        if (this.groups.length === 0) {
          return
        }
        // 设置当前默认选中的分组
        this.activeGroupIds = this.groups.map(group => group.group_id)
      })
    },
    selectGroupByServerId(serverId){
      for (let group of this.groups) {
        for (let server of group.servers) {
          if (server.serverId === serverId) {
            this.activeGroupIds = [group.group_id]
            return
          }
        }
      }
    },
    goServer(id){
      this.$router.push({
        name: "server",
        params: {id}
      })
    }
  }
}
</script>

<template>
  <div class="monitor-list">
    <el-collapse  v-if="groups.length > 0" v-model="activeGroupIds">
      <el-collapse-item v-for="group in groups" :name="group.group_id">
        <template #title>
          {{ group.group_name }}
        </template>
        <div v-for="server in group.servers">
          <router-link @click="goServer(server.server_id)" to="">
            <monitor-item
                :width="5" :height="16"
                :class="{'active': activeServerId == server.server_id}"
                :server="server">
              <template v-slot:statistics-bar-after>
                <span class="desc" style="float: left;">30天</span>
                <span class="desc" style="float: right">today</span>
              </template>
            </monitor-item>
          </router-link>
        </div>
      </el-collapse-item>
    </el-collapse>
    <el-empty v-else description="无服务器列表"></el-empty>
  </div>
</template>

<style scoped>
.monitor-list .monitor-item{
  display: block;
  text-decoration: none;
  padding: 13px 15px 10px;
  border-radius: 10px;
  transition: all ease-in-out .15s;
  font-size: 16px;
  cursor: pointer;
  margin-bottom: 10px;
}
.monitor-list .monitor-item:hover , .monitor-list .monitor-item.active{
  background-color: #e7faec;
}
.desc{
  font-size: 13px;
  color: #aaa;
}

.monitor-list a {
  color: #000;
  text-decoration: none;
}
</style>
