import request from "./request";

export function monitorGroups() {
    return request({
        url: '/admin-api/v1/monitor/groups',
        method: 'get'
    })
}
