<script>
import HpBar from "@/components/hp-bar.vue";
import MonitorFaults from "@/components/monitor-faults.vue";
import {getServerInfo, getServerStatisticsList, deleteServer} from "@/api/server.js";

export default {
  components: {MonitorFaults, HpBar},
  data() {
    return {
      total: [],
      server: null,
      serverId: null,

      // 1. 分 2. 时 3. 日 4. 月
      serverStatisticsType: 3
    }
  },
  watch: {
    '$route.params.id'() {
      this.serverId = this.$route.params.id;
      this.getServerInfo(this.serverId);
      this.getServerStatisticsList(this.serverId);
    }
  },
  mounted() {
    this.serverId = this.$route.params.id;
    this.getServerInfo(this.serverId);
    this.getServerStatisticsList(this.serverId);
  },
  methods: {
    handleEdit(id) {
      this.$router.push(`/server/edit/${id}`);
    },
    handleDelete(id) {
      this.$confirm('此操作将永久删除该服务器, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        deleteServer(id).then(() => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          // 路由返回
          this.$router.go(-1);
        }).catch((err) => {
          this.$message({
            type: 'error',
            message: '删除失败!' + err
          });
        });
      }).catch(() => {
      });
    },
    getServerInfo(id) {
      getServerInfo(id).then(data => {
        this.server = data
      });
    },
    getServerStatisticsList(id) {
      getServerStatisticsList(id, this.serverStatisticsType).then(data => {
        this.total = data
      });
    }
  }
}
</script>

<template>
  <div v-if="server" class="server">
    <h1 class="mb-3">{{ server.name }}</h1>
    <div class="mb-3">
      {{ server.remark }}
    </div>
    <div class="functions mb-3">
      <div class="btn-group" role="group">
        <button class="btn btn-normal" @click="handleEdit(server.id)"
                style="border-radius: 30px 0 0 30px;">
          编辑
        </button>
        <button class="btn btn-normal">
          终端
        </button>
        <button class="btn btn-normal">
          安装脚本
        </button>
        <button class="btn btn-danger" @click="handleDelete(server.id)"
                style="border-radius: 0 30px 30px 0;">
          删除
        </button>
      </div>
    </div>

    <div class="shadow-box mb-3">
      <div class="row">
        <div class="col-md-8">
          <div
              class="wrap"
              style="padding: 7.5px 2.6px; width: 100%;"
          >
            <hp-bar :width="10" :height="30" :total="total"></hp-bar>
          </div>
          <span class="word">最近30天</span>
        </div>
        <div class="col-md-4 text-center">
            <span v-if="server.status === 1" class="badge rounded-pill bg-primary"
                  style="font-size: 30px;">
              在线
            </span>
          <span v-else class="badge rounded-pill bg-danger"
                style="font-size: 30px;">
            离线
          </span>
        </div>
      </div>
    </div>

    <div v-if="server.server_info" class="shadow-box big-padding text-center stats mb-3">
      <div class="row">
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">IP</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            当前
          </p>
          <span class="col-4 col-sm-12 num">
            {{ server.ip }}
          </span>
        </div>
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">系统</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            os
          </p>
          <span class="col-4 col-sm-12 num">
              {{ server.server_info.os }}
          </span>
        </div>
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">系统型号</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            platform
          </p>
          <span class="col-4 col-sm-12 num">
              {{ server.server_info.platform }}
          </span>
        </div>
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">最后在线时间</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            Last online
          </p>
          <span class="col-4 col-sm-12 num">
              <span class="" title="24小时">
                  {{ server.last_online_time }}
              </span>
          </span>
        </div>
      </div>
    </div>

    <div class="shadow-box mb-3">
      <monitor-faults :server-id="serverId"></monitor-faults>
    </div>
  </div>
</template>

<style scoped>
.stats p {
  font-size: 13px;
  color: #aaa;
}

@media (max-width: 550px) {
  .stats h4 {
    font-size: 1.1rem;
  }
}
</style>
