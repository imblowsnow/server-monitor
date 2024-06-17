<script>
import MonitorItem from '@/components/monitor-item.vue'
export default {
  name: "MonitorList",
  components: {
    MonitorItem
  }
}
</script>

<template>
  <div class="monitor-list">
    <el-collapse accordion>
      <el-collapse-item v-for="i in 10" :name="i">
        <template #title>
          分组{{ i }}
        </template>
        <monitor-item v-for="i in 10"></monitor-item>
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
