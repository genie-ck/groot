<template>
  <el-config-provider :locale="zhCn">
    <router-view />
  </el-config-provider>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { useDebounceFn, useEventListener } from '@vueuse/core'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import { useLayoutStore } from '@/stores/layout'

const store = useLayoutStore()
const baseWidth = 992

const resizeFun = () => {
  const clientWidth = document.documentElement.clientWidth
  store.updateCollapse(clientWidth < baseWidth)
}
const debounceResize = useDebounceFn(resizeFun, 300)

onMounted(() => {
  useEventListener(window, 'resize', debounceResize)
})
</script>

<style>
#app {
  width: 100%;
  height: 100%;
}
</style>
