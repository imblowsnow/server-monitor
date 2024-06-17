import {registerComponent} from "../../static/lib/component.js";

export default {
    components: {
        navHeader: registerComponent('nav-header'),
        navFooter: registerComponent('nav-footer'),
        'monitor-list': registerComponent('monitor-list')
    },
    data(){
        return {
            // 宽度小于 800
            isMobile: document.body.clientWidth < 800
        }
    },
    mounted() {
        // 监听窗口大小变化
        window.onresize = () => {
            this.isMobile = document.body.clientWidth < 800
            console.log('isMobile',this.isMobile);
        }
    }
}
