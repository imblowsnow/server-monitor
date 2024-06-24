<script setup>
import { RouterLink, RouterView } from 'vue-router'
import {getSetting} from "@/api/backend/setting.js";
import {useSettingStore} from "@/stores/setting.js";

const settingStore = useSettingStore();


console.log(import.meta);
getSetting().then(data => {
  settingStore.updateSetting(data);
})
</script>

<template>
  <RouterView />
</template>

<style scoped>

</style>
