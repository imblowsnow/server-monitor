<script>
import ServerStatisticsBar from "@/components/server-statistics-bar.vue";
import MonitorFaults from "@/components/monitor-faults.vue";
import {deleteServer, getServerInfo, getServerStatisticsList} from "@/api/backend/server.js";
import help from "@/utils/help";
import {getSetting} from "@/api/backend/setting.js";
import {useSettingStore} from "@/stores/setting.js";

export default {
  computed: {
    help() {
      return help
    }
  },
  components: {MonitorFaults, ServerStatisticsBar},
  data() {
    return {
      settingStore: useSettingStore(),

      total: [],
      server: null,
      serverId: null,

      // 1. 分 2. 时 3. 日 4. 月
      serverStatisticsType: 3,

      colors:  [
        { color: '#5cb87a', percentage: 50 },
        { color: '#e6a23c', percentage: 80 },
        { color: '#f56c6c', percentage: 100 },
      ],

      ws: null
    }
  },
  watch: {
    '$route.params.id'() {
      this.serverId = this.$route.params.id;
      this.getServerInfo(this.serverId);
      this.getServerStatisticsList(this.serverId);
      this.$bus.emit('serverSelect', this.serverId);
    }
  },
  mounted() {
    this.serverId = this.$route.params.id;
    this.getServerInfo(this.serverId);
    this.getServerStatisticsList(this.serverId);
    this.$bus.emit('serverSelect', this.serverId);
    this.registerServerListen();
  },
  unmounted() {
    this.$bus.emit('serverSelect', null);
    if (this.ws){
      let ws = this.ws;
      this.ws = null
      ws.close();
    }
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
          this.$bus.emit('serverChanged');
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
    },
    registerServerListen(){
      // ws 链接 /ws/web
      // Create a new WebSocket connection
      this.ws = new WebSocket('/ws/web');

      // Listen for messages
      this.ws.addEventListener('message', (event) => {
        let messagePacket = JSON.parse(event.data);
        let data = JSON.parse(messagePacket.message);
        console.log('Message from server ', messagePacket.message_type, data);
        // 5000  服务上线
        // 5001 服务下线
        // 5002 服务状态更新
        if (messagePacket.message_type === 5000){
          if (data.server_id == this.serverId){
            this.server.status = 1;
            this.server.server_info = data.server_info;
            this.server.server_state = data.server_state;
          }
        } else if (messagePacket.message_type === 5001){
          if (data == this.serverId){
            this.server.status = 0;
          }
        } else if (messagePacket.message_type === 5002){
            if (data.server_id == this.serverId){
              console.log('Server state update: ', data);
              this.server.server_state = data;
            }
        }
      });

      // Connection closed
      this.ws.addEventListener('close', (event) => {
        console.log('Server connection closed: ', event);
        if (this.ws == null) return;
        this.$message({
          type: 'error',
          message: '服务器连接已断开，正在重新链接!'
        });
        setTimeout(() => {
          this.registerServerListen();
        }, 1000);
      });
    },

    copyInstallScript() {
      let serverIp = this.settingStore.setting.server_ip;
      let serverPort = this.settingStore.setting.server_port;
      let key = this.server.key;
      let installScript = `curl https://cdn.jsdelivr.net/gh/imblowsnow/server-monitor@master/deploy.sh | bash -s -- ${serverIp} ${serverPort} ${key}`;
      console.log(installScript);

      // 复制到剪贴板
      let input = document.createElement('input');
      input.value = installScript;
      document.body.appendChild(input);
      input.select();
      document.execCommand('copy');
      document.body.removeChild(input);

      this.$message({
        type: 'success',
        message: '复制成功!'
      });
    },
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
        <button class="btn btn-primary" @click="handleEdit(server.id)">
          编辑
        </button>
        <!--        <button class="btn btn-normal">-->
        <!--          终端-->
        <!--        </button>-->
        <button class="btn btn-primary" @click="copyInstallScript" >
          安装脚本
        </button>
        <button class="btn btn-danger" @click="handleDelete(server.id)">
          删除
        </button>
      </div>
    </div>

    <template v-if="server.server_info">
      <div class="shadow-box mb-3">
        <div class="row">
          <div class="col-md-8">
            <div
                class="wrap"
                style="padding: 7.5px 2.6px; width: 100%;"
            >
              <server-statistics-bar :width="10" :height="30" :total="total"></server-statistics-bar>
            </div>
            <span class="desc">最近30天</span>
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

      <div class="shadow-box big-padding text-center stats mb-3">
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
        </div>
      </div>
      <div class="shadow-box big-padding text-center stats mb-3">
        <div class="row">
          <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
            <h4 class="col-4 col-sm-12">CPU型号</h4>
            <p class="col-4 col-sm-12 mb-0 mb-sm-2">
              cpu
            </p>
            <span class="col-4 col-sm-12 num">
                <el-text line-clamp="1" :title="server.server_info.cpu">
                  {{ server.server_info.cpu }}
                </el-text>
            </span>
          </div>
          <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
            <h4 class="col-4 col-sm-12">虚拟化</h4>
            <p class="col-4 col-sm-12 mb-0 mb-sm-2">
              virtualization
            </p>
            <span class="col-4 col-sm-12 num">
                <span class="">
                    {{ server.server_info.virtualization }}
                </span>
            </span>
          </div>
          <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
            <h4 class="col-4 col-sm-12">最后在线时间</h4>
            <p class="col-4 col-sm-12 mb-0 mb-sm-2">
              Last online
            </p>
            <span class="col-4 col-sm-12 num">
                <span class="">
                    {{ server.last_online_time }}
                </span>
            </span>
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col-md-6">
          <div class="shadow-box mb-3">
            <div style="display: flex;justify-content: space-between;">
              <h4 class="mb-3">CPU Load</h4>
              <span class="desc">最后更新于：{{server.server_state.update_time}}</span>
            </div>
            <div class="text-center mb-5">
              <el-progress type="dashboard" :stroke-width="20" :width="200" striped striped-flow
                           :percentage="help.formatPercent(100, server.server_state.cpu)"
                           :color="colors" />
            </div>
            <div class="row mb-3">
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">Load1</div>
                <span class="col-6 col-sm-12 num desc">
                {{ server.server_state.load1.toFixed(2) }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">Load5</div>
                <span class="col-6 col-sm-12 num desc">
                {{ server.server_state.load5.toFixed(2) }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">Load15</div>
                <span class="col-6 col-sm-12 num desc">
                {{ server.server_state.load15.toFixed(2) }}
              </span>
              </div>
            </div>
            <div class="row">
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">TCP连接数</div>
                <span class="col-6 col-sm-12 num desc">
                {{ server.server_state.tcp_conn_count }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">UDP连接数</div>
                <span class="col-6 col-sm-12 num desc">
                {{ server.server_state.udp_conn_count }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">进程数</div>
                <span class="col-6 col-sm-12 num desc">
                {{ server.server_state.process_count }}
              </span>
              </div>
            </div>
          </div>
        </div>
        <div class="col-md-6">
          <div class="shadow-box mb-3">
            <h4 class="mb-3">内存</h4>
            <div class="mb-3">
              <el-progress :percentage="help.formatPercent(server.server_info.mem_total, server.server_state.mem_used)"
                           :color="colors" striped striped-flow />
            </div>
            <div class="row">
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">内存大小</div>
                <span class="col-6 col-sm-12 num desc">
                {{ help.formatByteSize(server.server_info.mem_total) }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">内存占用</div>
                <span class="col-6 col-sm-12 num desc">
                {{ help.formatByteSize(server.server_state.mem_used) }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">交换分区大小</div>
                <span class="col-6 col-sm-12 num desc">
                {{ help.formatByteSize(server.server_info.swap_total) }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">交换分区占用</div>
                <span class="col-6 col-sm-12 num desc">
                {{ help.formatByteSize(server.server_state.swap_used) }}
              </span>
              </div>
            </div>
          </div>
          <div class="shadow-box mb-3">
            <h4 class="mb-3">磁盘</h4>
            <div class="mb-3">
              <el-progress :percentage="help.formatPercent(server.server_info.disk_total, server.server_state.disk_used)"
                           :format="() => help.formatByteSize(server.server_state.disk_used) + '/' + help.formatByteSize(server.server_info.disk_total)"
                           :color="colors" striped striped-flow />
            </div>
          </div>
          <div class="shadow-box mb-3">
            <h4 class="mb-3">网络</h4>
            <div class="row">
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">入站速度</div>
                <span class="col-6 col-sm-12 desc">
                {{ help.formatByteSize(server.server_state.net_in_speed) }}/s
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">出站速度</div>
                <span class="col-6 col-sm-12 desc">
                {{ help.formatByteSize(server.server_state.net_out_speed) }}/s
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">入站流量</div>
                <span class="col-6 col-sm-12 desc">
                {{ help.formatByteSize(server.server_state.net_in_transfer) }}
              </span>
              </div>
              <div class="col-12 col-sm col row d-flex align-items-center d-sm-block ">
                <div class="col-6 col-sm-12">出站流量</div>
                <span class="col-6 col-sm-12 desc">
                {{ help.formatByteSize(server.server_state.net_out_transfer) }}
              </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="shadow-box mb-3">
        <monitor-faults :server-id="serverId" :limit="5"></monitor-faults>
      </div>
    </template>

    <template v-else>
      <div class="shadow-box">
        <el-empty description="未安装" />
      </div>
    </template>


  </div>
</template>

<style scoped>
.stats p {
  font-size: 13px;
  color: #aaa;
}
.desc{
  font-size: 13px;
  color: #aaa;
}
@media (max-width: 550px) {
  .stats h4 {
    font-size: 1.1rem;
  }
}
</style>
