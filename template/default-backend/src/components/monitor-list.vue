<script>
import MonitorItem from '@/components/monitor-item.vue'
import {monitorGroups} from "@/api/monitor.js";
export default {
  name: "MonitorList",
  components: {
    MonitorItem
  },
  data() {
    return {
      activeGroupIds: [1],
      groups: []
    }
  },
  mounted() {
    this.getMonitorGroups()
    // 监听事件
    this.$bus.on('serverChanged', () => {
      this.getMonitorGroups()
    })
  },
  methods:{
    getMonitorGroups() {
      monitorGroups().then(data => {
        this.groups = data
        console.log('monitorGroups',this.groups);
      })
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
        <monitor-item v-for="server in group.servers" :server="server"></monitor-item>
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
}
.monitor-list .monitor-item:hover {
  background-color: #e7faec;
}
</style>
