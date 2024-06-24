<script>
import {addServerGroup} from "@/api/backend/server_group.js";

export default {
  name: "index",
  data() {
    return {
      form: {
        group_name: '',
        sort: 1
      }
    }
  },
  methods: {
    addServerGroup() {
      addServerGroup(this.form).then(() => {
        this.$message({
          type: 'success',
          message: '添加成功!'
        });
        this.$router.push('/server_group');
      }).catch((err) => {
        this.$message({
          type: 'error',
          message: '添加失败!' + err
        });
      });
    }
  }
}
</script>

<template>
  <div>
    <h1 class="mb-3">添加分组</h1>
    <div class="shadow-box big-padding">
      <form>
        <div class="row">
          <div class="mb-3">
            <label class="form-label">分组名称</label>
            <input type="text" class="form-control" v-model="form.group_name"
                   required>
          </div>
          <div class="mb-3">
            <label class="form-label">分组排序</label>
            <input type="number" :min="1"
                   class="form-control" v-model="form.sort"
                   required>
          </div>
        </div>
        <div class="row">
          <button type="button" class="btn btn-primary" @click="addServerGroup">添加</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>

</style>
