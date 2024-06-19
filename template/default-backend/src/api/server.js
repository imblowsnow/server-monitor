import request from "./request";



export async function getServer(id) {
    return request({
        url: '/admin-api/v1/server/' + id,
        method: 'get'
    })
}
export async function getServerInfo(id) {
    return request({
        url: '/admin-api/v1/server/info/' + id,
        method: 'get'
    })
}

export async function addServer(data) {
    return request({
        url: '/admin-api/v1/server',
        method: 'post',
        data
    })
}

export async function updateServer(id,data) {
    return request({
        url: '/admin-api/v1/server/' + id,
        method: 'put',
        data
    })
}

export async function getServerStatisticsList(id, type=1){
    return request({
        url: `/admin-api/v1/server/statistics/${id}?type=${type}`,
        method: 'get'
    })
}

export async function pageServerFaults(server_id, page=1, size=10) {
    return request({
        url: `/admin-api/v1/server/fault/page?page=${page}&size=${size}&serverId=${server_id || ''}`,
        method: 'get'
    })
}

export async function editServerFaultRemark(id, remark) {
    return request({
        url: `/admin-api/v1/server/fault/${id}`,
        method: 'put',
        data: {remark}
    })
}

export async function deleteServer(id) {
    return request({
        url: '/admin-api/v1/server/' + id,
        method: 'delete'
    })
}
