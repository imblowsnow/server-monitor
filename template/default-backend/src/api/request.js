import axios from 'axios'
import router from '@/router'
// 设置请求头
export default async function (options) {
    options.headers = options.headers || {}
    return axios({
        ...options,
        headers: {
            ...options.headers,
            'Authorization': localStorage.getItem('token') || ''
        }
    }).then(response => {
        if (response.status === 200) {
            if (options.nocheck || response.data.code === 200) {
                return response.data.data
            } else {
                throw new Error(response.data.message)
            }
        }else{
            throw new Error('Request failed with status code ' + response.status)
        }
    }).catch(error => {
        if (error.response && error.response.status === 401) {
            console.log('未登录，跳转到登录页')
            localStorage.setItem('token', '')
            // router跳转到登录页
            router.push({
                name: 'login'
            })
        }else{
            throw error
        }   
    })
}
