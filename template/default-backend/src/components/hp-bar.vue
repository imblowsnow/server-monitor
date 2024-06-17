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
      num: 10
    }
  },
  computed: {
    list() {
      // 从后往前取
      return this.total.slice(-this.num);
    },
    statusClass() {
      return (status) => {
        if (status === 10) return 'error';
        if (status === 0) return 'success';
        return '';
      }
    }
  },
  mounted() {
    this.resizeHpBar();
  },
  methods: {
    resizeHpBar() {
      console.log('resizeHpBar', this.$refs.hpBar.clientWidth);
      // 计算最多显示多少个
      this.num = Math.floor(this.$refs.hpBar.clientWidth / (this.width + 4));
      console.log('resizeHpBar', this.num);
    }
  }
}
</script>

<template>
  <div ref="hpBar" @resize="resizeHpBar" class="hp-bar-big" style="transform: translateX(0px);" >
    <div v-for="item in this.list"
         class="beat"
         :class="statusClass(item.status)"
         :title="item.name"
         :style="{'width': width + 'px', 'height': height + 'px'}"
         style="margin: 2px; --hover-scale: 1.5;"></div>
  </div>
</template>

<style scoped>

</style>
