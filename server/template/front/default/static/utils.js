window.utils = {
    mergeUsage(state, host) {
        if (!state) {
            return 0
        }
        if (!host) {
            return state.cpu
        }
        return (state.cpu + (state.mem_used / host.mem_total * 100) + (state.disk_used / host.disk_total * 100)) / 3
    },
    readableBytes(bytes) {
        if (!bytes) {
            return '0B'
        }
        var i = Math.floor(Math.log(bytes) / Math.log(1024)),
            sizes = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
        return parseFloat((bytes / Math.pow(1024, i)).toFixed(2)) + sizes[i];
    },
    formatByteSize(bs) {
        const x = this.readableBytes(bs)
        return x != "NaN undefined" ? x : '0B'
    },
    secondToDate(s) {
        var d = Math.floor(s / 3600 / 24);
        if (d > 0) {
            return d + " 天"
        }
        var h = Math.floor(s / 3600 % 24);
        var m = Math.floor(s / 60 % 60);
        var s = Math.floor(s % 60);
        return h + ":" + ("0" + m).slice(-2) + ":" + ("0" + s).slice(-2);
    },
    formatTimestamp(t) {
        return new Date(t * 1000).toLocaleString()
    },
}
