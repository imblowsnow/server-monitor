<script>
export default {
  name: "monitor-faults",
  props: {
    serverId: {
      type: Number,
      default: null
    }
  },
  data() {
    return {
      faults: [
        { id: 1, name: 'fault1', status: 'up' },
      ],
      clientHeight: document.body.clientHeight,
    }
  },
}
</script>

<template>
  <div class="monitor-faults">
    <table class="table table-borderless table-hover">
      <thead>
      <tr>
        <th>名称</th>
        <th>开始时间</th>
        <th>结束时间</th>
        <th>备注</th>
      </tr>
      </thead>
      <tbody>
      <tr
          v-for="(beat, index) in faults"
          :key="index"
          :class="{ 'shadow-box': clientHeight <= 550}"
      >
        <td>服务器A</td>
        <td>2024-06-17 10:00:00</td>
        <td>2024-06-17 10:00:00</td>
        <td class="border-0">离线故障</td>
      </tr>

      <tr v-if="faults.length === 0">
        <td colspan="4">No important events</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>

</style>
