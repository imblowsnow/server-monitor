<script>
import {pageServerFaults, editServerFaultRemark} from "@/api/backend/server.js";

export default {
  name: "monitor-faults",
  props: {
    serverId: {
      type: Number,
      default: null
    },
    limit: {
      type: Number,
      default: 10
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
  watch: {
    serverId() {
      this.pageServerFaults();
    }
  },
  mounted() {
    this.page.limit = this.limit
    this.pageServerFaults();
  },
  methods: {
    pageServerFaults() {
      pageServerFaults(this.serverId, this.page.current, this.page.limit).then(data => {
        this.faults = data.list
        this.page.total = data.total
      });
    },
    editRemark(id, oldRemark) {
      this.$prompt('请输入备注', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /\S/,
        inputErrorMessage: '备注不能为空',
        inputValue: oldRemark
      }).then(({ value }) => {
        editServerFaultRemark(id, value).then(() => {
          this.$message({
            type: 'success',
            message: '修改成功!'
          });
          this.pageServerFaults();
        }).catch((err) => {
          this.$message({
            type: 'error',
            message: '修改失败!' + err
          });
        });
      });
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
  <div class="monitor-faults">
    <el-table :data="faults" class="table table-borderless table-hover">
      <el-table-column prop="server_name" label="名称">
        <template v-slot:default="scope">
          <router-link to="" @click="goServer(scope.row.server_id)">{{ scope.row.server_name }}</router-link>
        </template>
      </el-table-column>
      <el-table-column prop="start_time" label="开始时间"></el-table-column>
      <el-table-column prop="end_time" label="结束时间"></el-table-column>
      <el-table-column prop="duration" label="耗时(秒)"></el-table-column>
      <el-table-column prop="remark" label="备注"></el-table-column>
      <el-table-column label="操作">
        <template v-slot="scope">
          <el-button type="primary" size="small" @click="editRemark(scope.row.id, scope.row.remark)">修改备注</el-button>
        </template>
      </el-table-column>
    </el-table>

    <nav aria-label="Page navigation example">
      <el-pagination v-model:page-size="page.limit" v-model:current-page="page.current"
                     @size-change="pageServerFaults" @current-change="pageServerFaults"
          layout="prev, pager, next" :total="this.page.total" />
    </nav>
  </div>
</template>

<style scoped>

</style>
