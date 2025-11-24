import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi } from '@/api/auth'
import { ElMessage } from 'element-plus'

export const useUserStore = defineStore('user', () => {
  // 状态
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('userInfo') || 'null'))

  // 计算属性
  const isLoggedIn = computed(() => !!token.value)
  const username = computed(() => userInfo.value?.username || '')
  const email = computed(() => userInfo.value?.email || '')
  const userId = computed(() => userInfo.value?.id || '')
  const role = computed(() => userInfo.value?.role || '')
  const isAdmin = computed(() => userInfo.value?.role === 'admin')
  const iconUrl = computed(() => userInfo.value?.icon_url || '')

  // 登录方法
  const login = async (email, password) => {
    try {
      const response = await loginApi(email, password)

      // 保存 token 和用户信息
      token.value = response.token
      userInfo.value = response.user

      // 持久化到 localStorage
      localStorage.setItem('token', response.token)
      localStorage.setItem('userInfo', JSON.stringify(response.user))

      ElMessage.success('登录成功')
      return true
    } catch (error) {
      ElMessage.error(error.message || '登录失败')
      return false
    }
  }

  // 登出方法
  const logout = () => {
    token.value = ''
    userInfo.value = null

    // 清除 localStorage
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')

    ElMessage.success('已退出登录')
  }

  // 更新用户信息
  const setUserInfo = (info) => {
    userInfo.value = info
    localStorage.setItem('userInfo', JSON.stringify(info))
  }

  // 初始化：从 localStorage 恢复状态
  const init = () => {
    const savedToken = localStorage.getItem('token')
    const savedUserInfo = localStorage.getItem('userInfo')

    if (savedToken) {
      token.value = savedToken
    }

    if (savedUserInfo) {
      try {
        userInfo.value = JSON.parse(savedUserInfo)
      } catch (e) {
        console.error('Failed to parse user info:', e)
        localStorage.removeItem('userInfo')
      }
    }
  }

  return {
    // 状态
    token,
    userInfo,

    // 计算属性
    isLoggedIn,
    username,
    email,
    userId,
    role,
    isAdmin,
    iconUrl,

    // 方法
    login,
    logout,
    setUserInfo,
    init
  }
})
