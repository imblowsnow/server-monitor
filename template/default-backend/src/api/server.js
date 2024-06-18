import request from "./request";

export async function serverGroups() {
    return request({
        url: '/admin-api/v1/server/groups',
        method: 'get'
    })
}
