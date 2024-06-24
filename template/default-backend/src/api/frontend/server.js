import request from "@/api/request.js";

export async function getServerInfo(id) {
    return request({
        url: '/api/v1/server/info/' + id,
        method: 'get'
    })
}


export async function getServerStatisticsList(id, type=1){
    return request({
        url: `/api/v1/server/statistics/${id}?type=${type}`,
        method: 'get'
    })
}
