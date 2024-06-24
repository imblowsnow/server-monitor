<script>
import {getSetting, updateSetting} from "@/api/setting.js";
import {useSettingStore} from "@/stores/setting.js";

export default {
  name: "setting",
  data() {
    return {
      settingStore: useSettingStore(),
      form: {
        site_name: '',
        frontend_template: 'default',
        backend_template: 'default',

        admin_username: '',
        admin_password: '',

        server_ip: '',
        server_port: '',
      },
      frontendTemplates: [
        {
          value: 'default',
          label: '默认',
        }
      ],
      backendTemplates: [
        {
          value: 'default',
          label: '默认',
        }
      ],
    }
  },
  mounted() {
    this.getSetting();
  },
  methods: {
    save() {
      updateSetting(this.form).then(() => {
        this.$message({
          type: 'success',
          message: '保存成功!'
        });
        this.settingStore.updateSetting(this.form);
      }).catch((err) => {
        this.$message({
          type: 'error',
          message: '保存失败!' + err
        });
      });
    },
    getSetting() {
      getSetting().then(data => {
        this.form = data;
      });
    }
  }
}
</script>

<template>
  <div>
    <h1 class="mb-3">系统设置</h1>
    <div class="shadow-box big-padding">
      <form>
        <div class="row">
          <div class="col-md-6">
            <div class="mb-3">
              <label class="form-label">站点名称</label>
              <input type="text" class="form-control" v-model="form.site_name"
                     required>
            </div>
            <div class="mb-3">
              <label class="form-label">管理员账号</label>
              <input type="text" class="form-control" v-model="form.admin_username"
                     required>
            </div>
            <div class="mb-3">
              <label class="form-label">管理员密码</label>
              <input type="password" class="form-control" v-model="form.admin_password"
                     required>
            </div>
            <div class="mb-3">
              <label class="form-label">服务器IP</label>
              <input type="text" class="form-control" v-model="form.server_ip"
                     required>
            </div>
<!--            <div class="mb-3">-->
<!--              <label class="form-label">服务器端口</label>-->
<!--              <input type="text" class="form-control" v-model="form.server_port"-->
<!--                     required>-->
<!--            </div>-->
          </div>
          <div class="col-md-6">
            <div class="mb-3">
              <label class="form-label">首页模版</label>
              <select class="form-control" v-model="form.frontend_template">
                <option v-for="item in frontendTemplates" :value="item.value">{{item.label}}</option>
              </select>
            </div>
            <div class="mb-3">
              <label class="form-label">后台模版</label>
              <select class="form-control" v-model="form.backend_template">
                <option v-for="item in frontendTemplates" :value="item.value">{{item.label}}</option>
              </select>
            </div>
          </div>
        </div>
        <div class="row">
          <button type="button" class="btn btn-primary" @click="save">保存</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>

</style>
