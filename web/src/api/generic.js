import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

const BASE_URL = import.meta.env.VITE_BASE_URL

// create
export const createHandler = (resourceType, itemForm) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'create'
  return request(fullUrl, itemForm, 'post')
}
// delete
export const deleteHandler = (clusterId, nameSpace, resourceType, name) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'delete'
  return request(fullUrl, { clusterId, nameSpace, name }, 'post')
}
// update
export const updateHandler = (resourceType, itemForm) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'update'
  return request(fullUrl, itemForm, 'post')
}
// select
export const getListHandler = (clusterId, nameSpace, resourceType) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'list'
  return request(fullUrl, { clusterId, nameSpace }, 'get')
}
export const getHandler = (clusterId, nameSpace, resourceType, name) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'get'
  return request(fullUrl, { clusterId, nameSpace, name }, 'get')
}
export const getMetricsHandler = (clusterId, nameSpace, resourceType) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'metrics'
  return request(fullUrl, { clusterId, nameSpace }, 'get')
}
export const getStatusHandler = (clusterId, nameSpace, resourceType) => {
  const fullUrl = BASE_URL + '/' + resourceType + '/' + 'status'
  return request(fullUrl, { clusterId, nameSpace }, 'get')
}
