import axios from 'axios'

axios.defaults.baseURL = '/api'

export default (method, url, data) => {
  return axios({
    url,
    method,
    [method.toLowerCase() === 'get' ? 'params' : 'data']: data
  })
}