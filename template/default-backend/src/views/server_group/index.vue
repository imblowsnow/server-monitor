<script>

import {deleteServerGroup, pageServerGroupList} from "@/api/server_group.js";

export default {
  name: "index",
  data() {
    return {
      list: [],
      page: {
        current: 1,
        size: 10,
        total: 0
      }
    }
  },
  mounted() {
    this.pageChange();
  },
  methods: {
    pageChange(){
      pageServerGroupList({
        page: this.page.current,
        size: this.page.size
      }).then(data => {
        this.list = data.list;
        this.page.total = data.total;
        console.log('pageServerGroupList', data);
      }).catch(() => {
        this.$message({
          type: 'error',
          message: '获取数据失败!'
        });
      });
    },
    handleEdit(id) {
      this.$router.push(`/server_group/edit/${id}`);
    },
    handleAdd() {
      this.$router.push(`/server_group/add`);
    },
    handleDelete(id) {
      this.$confirm('此操作将永久删除该分组, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteServerGroup(id).then(() => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          this.pageChange();
        }).catch((err) => {
          this.$message({
            type: 'error',
            message: '删除失败!' + err
          });
        });
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    }

  }
}
</script>

<template>
  <div>
    <h1 class="mb-3">分组管理</h1>
    <div class="functions mb-3">
      <div class="btn-group" role="group">
        <button class="btn btn-primary" @click="handleAdd()"
        >
          新增
        </button>
<!--        <button class="btn btn-normal" style="border-radius: 0 30px 30px 0;">-->
<!--          编辑-->
<!--        </button>-->
      </div>
    </div>

    <div class="shadow-box mb-3">
      <table class="table table-borderless table-hover">
        <thead>
        <tr>
          <th>名称</th>
          <th>排序</th>
          <th>操作</th>
        </tr>
        </thead>
        <tbody>
        <tr
            v-for="(item, index) in list"
            :key="index"
        >
          <td>{{ item.group_name }}</td>
          <td>{{ item.sort }}</td>
          <td style="width: 150px">
            <el-button class="btn btn-primary" type="primary"
                       size="small" :disabled="item.system === 1"
                       @click="handleEdit(item.id)">编辑</el-button>
            <el-button class="btn btn-danger" size="small" type="danger"
                       :disabled="item.system === 1"
                       @click="handleDelete(item.id)">删除</el-button>
          </td>
        </tr>
        <tr v-if="list.length === 0">
          <td colspan="3" style="text-align: center">无数据</td>
        </tr>
        </tbody>
      </table>
      <nav v-if="page.total > 0" aria-label="Page navigation example">
        <el-pagination v-model:page-size="page.limit" v-model:current-page="page.current" @size-change="pageChange" @current-change="pageChange"
                       layout="prev, pager, next" :total="page.total" />
      </nav>
    </div>
  </div>
</template>

<style scoped>

</style>
