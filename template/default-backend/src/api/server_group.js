import request from './request'

export function pageServerGroupList(params) {
    return request({
        url: '/admin-api/v1/server_group/page',
        method: 'get',
        params
    })
}

export function getServerGroups() {
    return request({
        url: '/admin-api/v1/server_group',
        method: 'get'
    })
}

export function getServerGroup(id) {
    return request({
        url: `/admin-api/v1/server_group/${id}`,
        method: 'get'
    })
}

export function addServerGroup(data) {
    return request({
        url: '/admin-api/v1/server_group',
        method: 'post',
        data
    })
}

export function updateServerGroup(id, data) {
    return request({
        url: `/admin-api/v1/server_group/${id}`,
        method: 'put',
        data
    })
}

export function deleteServerGroup(id) {
    return request({
        url: `/admin-api/v1/server_group/${id}`,
        method: 'delete'
    })
}
