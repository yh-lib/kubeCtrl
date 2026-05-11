// 封装 axios
/*
    1. 封装可以简化我的接口调用的代码，可以去掉一些重复的代码
    2. 换掉 Axios 也是比较简单的
*/

import axios from 'axios'
import { API_CONFIG, CONFIG } from '../config/index.js'
import router from '../router/index.js'
import { ElMessage, ElLoading } from 'element-plus'
import { ref } from 'vue'
import { load } from 'js-yaml'

let lastErrorMessage = ''
let lastErrorAt = 0

const notifyError = (message) => {
  const text = message || '请求失败'
  const now = Date.now()
  if (text === lastErrorMessage && now - lastErrorAt < 800) {
    return
  }
  lastErrorMessage = text
  lastErrorAt = now
  ElMessage.error(text)
}

const getErrorMessage = (err, fallback = '请求失败') => {
  return error?.response?.data?.message || err?.message || fallback
}

// axios 全局配置
axios.defaults.baseURL = API_CONFIG.baseUrl

// 获取完整请求url
const getFullRequestUrl = (config) => {
  if (!config) {
    return ''
  }
  return axios.getUri({
    baseURL: axios.defaults.baseURL,
    url: config.url,
    params: config.params,
  })
}

let loading = ref()
// axios 拦截器
// 添加请求拦截器
axios.interceptors.request.use(
  (config) => {
    loading = ElLoading.service({
      lock: true,
      text: '加载中。。。',
      background: 'rgba(0, 0, 0, 0.7)',
    })
    console.log('请求拦截器:::config:::', config, 'URL:::', getFullRequestUrl(config))
    // 在发送请求之前做些什么
    // 添加请求时的时间戳，处理浏览器缓存问题
    if (config.method == 'get') {
      let timeStamp = new Date().getTime()
      config.params
        ? (config.params.timeStamp = timeStamp)
        : (config.params = { timeStamp: timeStamp })
    }
    // 在request header中加入token
    config.headers.Authorization = window.localStorage.getItem(CONFIG.TOKEN_NAME)
    return config
  },
  (err) => {
    // 对请求错误做些什么
    return Promise.reject(err)
  }
)

// 添加响应拦截器
axios.interceptors.response.use((response) => {
  // 2xx 范围内的状态码都会触发该函数。
  // 对响应数据做点什么
  console.log('响应拦截器:::Response:::', response, 'URL:::', getFullRequestUrl(response?.config))
  if (response.data.status === 200) {
    return response
  } else if (response.data.status === 400) {
    notifyError(response.data.message)
    return Promise.reject(new Error(response.data.message))
  } else if (response.data.status === 401) {
    // 提示信息
    ElMessage({
      message: response.data.message,
      type: 'warning',
    })
    // token已失效,移除本地token
    window.localStorage.removeItem(CONFIG.TOKEN_NAME)
    // 跳转到登录页
    router.currentRoute.value.path != '/login' && router.push('/login')
    return Promise.reject(new Error(response.data.message))
  } else if (response.data.status === 403) {
    // 提示信息
    console.log('响应拦截器403:::Response:::', response.data.status)
    notifyError(response.data.message)
    return Promise.reject(new Error(response.data.message))
  }
})

const request = (url = '', data = {}, method = 'get', timeout = 3000) => {
  return new Promise((resolve, reject) => {
    // GET POST
    const methodLower = method.toLowerCase()
    if (methodLower == 'get') {
      axios({
        method: methodLower,
        params: data,
        timeout: timeout,
        url: url,
      })
        .then((response) => {
          // 正常拿到响应数据后
          loading.close()
          resolve(response)
        })
        .catch((err) => {
          // 响应失败或错误后
          loading.close()
          if (err.message.includes('timeout')) {
            ElMessage.error('后端服务超时未响应')
          }
          reject(err)
        })
    } else if (methodLower === 'post') {
      axios({
        method: methodLower,
        data: data,
        timeout: timeout,
        url: url,
      })
        .then((response) => {
          // 正常拿到响应数据后
          loading.close()
          resolve(response)
        })
        .catch((err) => {
          // 响应失败或错误后
          loading.close()
          reject(err)
        })
    }
  })
}

export default request
