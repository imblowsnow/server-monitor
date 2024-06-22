
function readableBytes(bytes) {
    if (!bytes) {
        return '0B'
    }
    var i = Math.floor(Math.log(bytes) / Math.log(1024)),
        sizes = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
    return parseFloat((bytes / Math.pow(1024, i)).toFixed(2)) + sizes[i];
}


function formatByteSize(bs) {
    const x = readableBytes(bs)
    return x != "NaN undefined" ? x : '0B'
}

function formatPercent( total, used) {
    return parseInt(used / total * 100) || 0
}
export default {
    readableBytes,
    formatByteSize,
    formatPercent
}
