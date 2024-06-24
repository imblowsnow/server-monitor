import request from '@/api/request.js'


export function login(data) {
    return request({
        url: '/admin-api/v1/user/login',
        method: 'post',
        data
    })
}
