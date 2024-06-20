<script>
import {getServerGroups} from "@/api/server_group.js";
import {getServer, updateServer} from "@/api/server.js";
export default {
  data() {
    return {
      groups: [],
      form: {
        groupId: 1,
        name: '',
        remark: '',
        sort: 1,
      },
      id: null
    }
  },
  mounted() {
    this.getGroups();

    let id = this.$route.params.id;
    if (!id) {
      this.$message.error('服务器ID不能为空!');
      this.$router.push('/server');
    }
    this.id = id;
    this.getServer(this.id);
  },
  methods: {
    getGroups() {
      getServerGroups().then(data => {
        this.groups = data;
      }).catch(() => {
        this.$message.error('获取分组失败!');
      });
    },
    getServer(id) {
      getServer(id).then(data => {
        this.form = data;
      }).catch(() => {
        this.$message.error('获取服务器失败!');
      });
    },
    updateServer() {
      updateServer(this.id, this.form).then(data => {
        this.$bus.emit('serverChanged');

        this.$router.push({ name: 'server', params: { id: this.id} });
      }).catch(err => {
        this.$message.error('更新失败!' + err);
      });
    },
    generateKey(){
      // 生成uuid
      this.form.key = this.generateUUID()
    },

    generateUUID() {
      return 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random() * 16 | 0,
            v = c === 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
      })
    }
  }
}
</script>

<template>
  <div class="server-add">
    <h1 class="mb-3">编辑服务器</h1>
    <div class="shadow-box big-padding">
      <form>
        <div class="row">
          <div class="col-md-6">
            <div class="mb-3">
              <label class="form-label">服务器分组</label>
              <div class="input-group">
                <select class="form-control"
                        v-model="form.group_id"
                        required>
                  <option v-for="group in groups" :value="group.id">{{group.group_name}}</option>
                </select>
<!--                <button type="button" class="btn btn-primary" aria-label="open modal to ">-->
<!--                  <svg class="svg-inline&#45;&#45;fa fa-plus fa-w-14" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="plus" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path class="" fill="currentColor" d="M416 208H272V64c0-17.67-14.33-32-32-32h-32c-17.67 0-32 14.33-32 32v144H32c-17.67 0-32 14.33-32 32v32c0 17.67 14.33 32 32 32h144v144c0 17.67 14.33 32 32 32h32c17.67 0 32-14.33 32-32V304h144c17.67 0 32-14.33 32-32v-32c0-17.67-14.33-32-32-32z"></path></svg>-->
<!--                </button>-->
              </div>

            </div>
            <div class="mb-3">
              <label class="form-label">服务器名称</label>
              <input type="text" class="form-control" v-model="form.name"
                     required>
            </div>
            <div class="mb-3">
              <label class="form-label">服务器备注信息</label>
              <input type="text" class="form-control" v-model="form.remark"
                     required>
            </div>
          </div>
          <div class="col-md-6">
            <div class="mb-3">
              <label class="form-label">排序</label>
              <input type="number" class="form-control" v-model="form.sort"
                     min="0"
                     required>
            </div>
            <div class="mb-3">
              <label class="form-label">首页展示</label>
              <select class="form-control" v-model="form.show_index">
                <option :value="1">展示</option>
                <option :value="0">不展示</option>
              </select>
            </div>
            <div class="mb-3">
              <label class="form-label">秘钥</label>
              <div class="input-group">
                <input class="form-control" v-model="form.key" required disabled>
                <button type="button" class="btn btn-primary" aria-label="open modal to " @click="generateKey">
                  重新生成
                </button>
              </div>

            </div>
          </div>
        </div>
        <div class="row">
          <button type="button" class="btn btn-primary" @click="updateServer">保存</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>

</style>
