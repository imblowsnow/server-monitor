<script>
import MonitorItem from '@/components/monitor-item.vue'
import {monitorGroups} from "@/api/backend/monitor.js";
export default {
  name: "MonitorList",
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
      monitorGroups().then(data => {
        this.groups = data
        if (this.groups.length === 0) {
          return
        }
        // 设置当前默认选中的分组
        this.activeGroupIds = [this.groups[0].group_id]
        console.log('activeGroupIds',this.activeGroupIds);
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
    }
  }
}
</script>

<template>
  <div class="monitor-list">
    <el-collapse v-model="activeGroupIds">
      <el-collapse-item v-for="group in groups" :name="group.group_id">
        <template #title>
          {{ group.group_name }}
        </template>
        <router-link v-for="server in group.servers"
                     :to="'/server/' + server.server_id" style="margin-left: 0px;">
          <monitor-item
                        :class="{'active': activeServerId == server.server_id}"
                  :server="server">
          </monitor-item>
        </router-link>
      </el-collapse-item>
    </el-collapse>
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

.monitor-list a{
  color: #000;
  text-decoration: none;
}
</style>
