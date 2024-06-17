<script>
import HpBar from "@/components/hp-bar.vue";

export default {
  name: "MonitorItem",
  components: {HpBar},
  props: {
    monitor: {
      type: Object,
      required: true
    }
  },
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
  }
}
</script>

<template>
  <div class="monitor-item">
    <router-link to="/server/1" style="margin-left: 0px;">
      <div class="row">
        <div class="col-8 small-padding">
          <div class="info">
                    <span class="badge rounded-pill bg-primary" title="24小时">
                        100%
                    </span>
            ces
          </div>
        </div>
        <div class="col-4">
          <div class="wrap" style="padding: 4px 1.25px; width: 100%;" >
            <hp-bar :total="total"></hp-bar>
          </div>
        </div>
      </div>
    </router-link>
  </div>
</template>

<style scoped>

</style>
