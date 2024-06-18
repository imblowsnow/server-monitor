<script>
import {pageServerFaults} from "@/api/server.js";

export default {
  name: "monitor-faults",
  props: {
    serverId: {
      type: Number,
      default: null
    }
  },
  data() {
    return {
      faults: [
        { id: 1, name: 'fault1', status: 'up' },
      ],
      page: {
        current: 1,
        limit: 10,
        total: 0
      }
    }
  },
  mounted() {
    this.pageServerFaults();
  },
  methods: {
    pageServerFaults() {
      pageServerFaults(this.serverId, this.page.current, this.page.limit).then(data => {
        this.faults = data.list
        this.page.total = data.total
      });
    }
  }
}
</script>

<template>
  <div class="monitor-faults">
    <table class="table table-borderless table-hover">
      <thead>
      <tr>
        <th>名称</th>
        <th>开始时间</th>
        <th>结束时间</th>
        <th>备注</th>
      </tr>
      </thead>
      <tbody>
      <tr
          v-for="(fault, index) in faults"
          :key="index"
      >
        <td>{{ fault.name }}</td>
        <td>{{ fault.start_time }}</td>
        <td>{{ fault.end_time }}</td>
        <td class="border-0">
          {{ fault.remark }}
          <el-button class="btn btn-primary" size="small">备注</el-button>
        </td>
      </tr>

      <tr v-if="faults.length === 0">
        <td colspan="4" style="text-align: center">无故障事件</td>
      </tr>
      </tbody>
    </table>
    <nav aria-label="Page navigation example">
      <el-pagination v-model:page-size="page.limit" v-model:current-page="page.current" @size-change="pageServerFaults" @current-change="pageServerFaults"
          layout="prev, pager, next" :total="this.page.total" />
    </nav>
  </div>
</template>

<style scoped>

</style>
