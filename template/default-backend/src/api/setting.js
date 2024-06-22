import request from '@/api/request'
export function getSetting() {
    if (!localStorage.getItem("token")){
        return request({
            url: '/api/v1/config',
            method: 'get'
        })
    }
    return request({
        url: '/admin-api/v1/config',
        method: 'get'
    })
}

export function updateSetting(data) {
    return request({
        url: '/admin-api/v1/config',
        method: 'put',
        data
    })
}
