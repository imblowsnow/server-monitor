import request from "./request";

export function monitorGroups() {
    return request({
        url: '/admin-api/v1/monitor/groups',
        method: 'get'
    })
}


export function dashboardTotal() {
    return request({
        url: '/admin-api/v1/monitor/total',
        method: 'get'
    })
}


export function apiMonitorGroups() {
    return request({
        url: '/api/v1/monitor/groups',
        method: 'get'
    })
}

export function apiDashboardTotal() {
    return request({
        url: '/api/v1/monitor/total',
        method: 'get'
    })
}
