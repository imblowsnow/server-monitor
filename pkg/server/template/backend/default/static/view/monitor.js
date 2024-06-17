import { registerComponent } from '../lib/component.js'

// 等待页面加载完毕
window.onload = function () {
    const { createApp } = Vue

    createApp({
        components: {
            monitorLayout: registerComponent('monitor-layout'),
            monitorList: registerComponent('monitor-list'),
        },
        data() {
            return {
                total: {
                    up: 0,
                    down: 0,
                },
                faults: [
                    { id: 1, name: 'fault1', status: 'up' },
                ],
                clientHeight: document.body.clientHeight,
            }
        },
        mounted() {
            console.log('mounted')
        }
    }).use(ElementPlus).mount('#app')
}
