<script>
import {updateServerGroup, getServerGroup} from "@/api/server_group.js";

export default {
  name: "index",
  data() {
    return {
      form: {
        group_name: '',
        sort: 1
      },
      id: null
    }
  },
  mounted() {
    this.id = this.$route.params.id;
    this.getServerGroup(this.id);
  },
  methods: {
    getServerGroup(id) {
      getServerGroup(id).then(data => {
        this.form = data
      });
    },
    updateServerGroup() {
      updateServerGroup(this.id, this.form).then(() => {
        this.$message({
          type: 'success',
          message: '更新成功!'
        });
        this.$router.push('/server_group');
      }).catch((err) => {
        this.$message({
          type: 'error',
          message: '更新失败!' + err
        });
      });
    }
  }
}
</script>

<template>
  <div>
    <h1 class="mb-3">编辑分组</h1>
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
          <button type="button" class="btn btn-primary" @click="updateServerGroup">更新</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>

</style>
