import axios from 'axios'

export default async function (options) {
    return axios(options).then(response => {
        if (response.status === 200) {
            if (options.nocheck || response.data.code === 200) {
                return response.data.data
            } else {
                throw new Error(response.data.message)
            }
        }else{
            throw new Error('Request failed with status code ' + response.status)
        }
    })
}
