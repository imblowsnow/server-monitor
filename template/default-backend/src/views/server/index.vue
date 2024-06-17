<script>
import HpBar from "@/components/hp-bar.vue";
import MonitorFaults from "@/components/monitor-faults.vue";

export default {
  components: {MonitorFaults, HpBar},
  data() {
    return {
      total: []
    }
  },
  mounted() {
    this.total = Array.from({length: 100}, (v, k) => {
      return {
        name: k,
        status: Math.floor(Math.random() * 11)
      }
    });
  },
  methods: {
    handleEdit(id) {
      this.$router.push(`/edit/${id}`);
    }
  }
}
</script>

<template>
  <div class="server">
    <h1 class="mb-3">服务名称</h1>
    <div class="mb-3">
      备注信息
    </div>
    <div class="functions mb-3">
      <div class="btn-group" role="group">
        <button class="btn btn-normal" style="border-radius: 30px 0 0 30px;">
          编辑
        </button>
        <button class="btn btn-danger" style="border-radius: 0 30px 30px 0;">
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
          <span class="word">检测频率 60 秒</span>
        </div>
        <div class="col-md-4 text-center">
            <span

                class="badge rounded-pill bg-primary"
                style="font-size: 30px;"
            >正常</span
            >
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
                        127.0.0.1
                    </span>
        </div>
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">系统</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            当前
          </p>
          <span class="col-4 col-sm-12 num">
                        ubunut
                    </span>
        </div>
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">在线时间</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            (24小时)
          </p>
          <span class="col-4 col-sm-12 num">
                        <span class="" title="24小时">
                            100%
                        </span>
                    </span>
        </div>
        <div class="col-12 col-sm col row d-flex align-items-center d-sm-block">
          <h4 class="col-4 col-sm-12">在线时间</h4>
          <p class="col-4 col-sm-12 mb-0 mb-sm-2">
            (30天)
          </p>
          <span class="col-4 col-sm-12 num">
                        <span class="" title="30天">100%</span>
                    </span>
        </div>
      </div>
    </div>

    <div class="shadow-box mb-3">
      <monitor-faults></monitor-faults>
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
