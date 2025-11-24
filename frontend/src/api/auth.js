import request from './request'

/**
 * 用户登录
 * @param {string} email - 用户邮箱
 * @param {string} password - 用户密码
 * @returns {Promise} 返回包含 token 和用户信息的响应
 */
export const login = (email, password) => {
  return request({
    url: '/users/login',
    method: 'post',
    data: {
      email,
      password
    }
  })
}

/**
 * 创建用户（注册）
 * @param {Object} userData - 用户注册数据
 * @returns {Promise}
 */
export const register = (userData) => {
  return request({
    url: '/users',
    method: 'post',
    data: userData
  })
}

/**
 * 获取用户信息
 * @param {string} id - 用户ID
 * @returns {Promise}
 */
export const getUserInfo = (id) => {
  return request({
    url: `/users/${id}`,
    method: 'get'
  })
}

/**
 * 更新用户信息
 * @param {string} id - 用户ID
 * @param {Object} userData - 用户数据
 * @returns {Promise}
 */
export const updateUser = (id, userData) => {
  return request({
    url: `/users/${id}`,
    method: 'put',
    data: userData
  })
}
