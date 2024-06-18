<script>
export default {
  name: "hp-bar",
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
    }
  },
  data() {
    return {
      num: 10,
      list: []
    }
  },
  computed: {
    statusClass() {
      return (status) => {
        if (status === 0) return 'empty';
        if (status < 80 ) return 'error';
        return '';
      }
    }
  },
  mounted() {
    this.resizeHpBar();
  },
  watch: {
    total() {
      this.resizeHpBar();
    }
  },
  methods: {
    resizeHpBar() {
      // 计算最多显示多少个
      this.num = Math.floor(this.$refs.hpBar.clientWidth / (this.width + 4));
      this.list = this.total.reverse().slice(-this.num);
    }
  }
}
</script>

<template>
  <div ref="hpBar" @resize="resizeHpBar" class="hp-bar-big" style="transform: translateX(0px);" >
    <div v-for="item in this.list"
         class="beat"
         :class="statusClass(item.status)"
         :title="'在线率' + item.status + '%\n' + item.start_time + ' - ' + item.end_time"
         :style="{'width': this.width + 'px', 'height': this.height + 'px'}"
         style="margin: 2px; --hover-scale: 1.5;"></div>
  </div>
</template>

<style scoped>
.hp-bar-big{
  display: flex;
}
.hp-bar-big .beat {
  display: inline-block;
  background-color: var(--theme-color);
  border-radius: 50rem;
}

.hp-bar-big .beat.empty {
  background-color: #dcdcdc;
}

.hp-bar-big .beat.error {
  background-color: #dc3545;
}

.hp-bar-big .beat.pending {
  background-color: #f8a306;
}

.hp-bar-big .beat.maintenance {
  background-color: #1747f5;
}

/*.hp-bar-big .beat.empty:hover{*/
/*    transition: all ease-in-out 0.15s;*/
/*    opacity: 0.8;*/
/*    transform: scale(1.5);*/
/*}*/
.hp-bar-big .beat:not(.empty):hover {
  transition: all ease-in-out .15s;
  opacity: .8;
  transform: scale(var(--hover-scale));
}
</style>
