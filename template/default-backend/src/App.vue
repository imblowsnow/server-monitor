<script setup>
import { RouterLink, RouterView } from 'vue-router'
import {getSetting} from "@/api/setting.js";
import {useSettingStore} from "@/stores/setting.js";

const settingStore = useSettingStore();

getSetting().then(data => {
  settingStore.updateSetting(data);
})
</script>

<template>
  <RouterView />
</template>

<style scoped>

</style>
