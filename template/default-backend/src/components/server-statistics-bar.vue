<script>
export default {
  name: "server-statistics-bar",
  props: {
    total: {
      type: Array,
      required: true
    },
    width: {
      type: Number,
      default: 5
    },
    height: {
      type: Number,
      default: 16
    },
    maxNum: {
      type: Number,
      default: null
    }
  },
  data() {
    return {
      list: []
    }
  },
  computed: {
    statusClass() {
      return (status) => {
        if (status === 0) return 'empty';
        if (status < 50 ) return 'error';
        if (status < 90 ) return 'pending';
        return '';
      }
    }
  },
  mounted() {
    this.resizeServerStatisticsBar();
  },
  watch: {
    total() {
      this.resizeServerStatisticsBar();
    }
  },
  methods: {
    resizeServerStatisticsBar() {
      if (this.maxNum) this.list = this.total.reverse().slice(-this.maxNum);
      else this.list = this.total.reverse();
    }
  }
}
</script>

<template>
  <div ref="hpBar" @resize="resizeServerStatisticsBar" class="server-statistics-bar-big" style="transform: translateX(0px);" >
    <div v-for="item in this.list"
         class="beat"
         :class="statusClass(item.status)"
         :title="'在线率' + item.status + '%\n' + item.start_time + ' - ' + item.end_time"
         :style="{'width': this.width + 'px', 'height': this.height + 'px'}"
         style="margin: 2px; --hover-scale: 1.5;"></div>
  </div>
</template>

<style scoped>
.server-statistics-bar-big{

}
.server-statistics-bar-big .beat {
  display: inline-block;
  background-color: var(--theme-color);
  border-radius: 50rem;
}

.server-statistics-bar-big .beat.empty {
  background-color: #dcdcdc;
}

.server-statistics-bar-big .beat.error {
  background-color: #dc3545;
}

.server-statistics-bar-big .beat.pending {
  background-color: #f8a306;
}

.server-statistics-bar-big .beat.maintenance {
  background-color: #1747f5;
}

/*.server-statistics-bar-big .beat.empty:hover{*/
/*    transition: all ease-in-out 0.15s;*/
/*    opacity: 0.8;*/
/*    transform: scale(1.5);*/
/*}*/
.server-statistics-bar-big .beat:not(.empty):hover {
  transition: all ease-in-out .15s;
  opacity: .8;
  transform: scale(var(--hover-scale));
}
</style>
