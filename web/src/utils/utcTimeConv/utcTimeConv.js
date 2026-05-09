import { ref, computed } from 'vue'

// 计算北京时间
export const utcTime2BjTime = (utcTime) => {
  const date = new Date(utcTime) // 根据浏览器的本地时区进行显示
  const originalTime = date.toLocaleString('zh-CN', { timeZone: 'Asia/Shanghai' }) //   指定时区来确保显示正确
  const formattedTime = originalTime.replace(/\//g, '-') // 修改显示格式
  return formattedTime
}
