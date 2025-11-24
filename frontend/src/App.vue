<script setup>
import { useMcpStore } from '@/stores/mcp'
import { useUserStore } from '@/stores/user'
import { ref, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'

const router = useRouter()
const mcpStore = useMcpStore()
const userStore = useUserStore()
const searchText = ref('')

let searchTimer = null

// 初始化用户状态
onMounted(() => {
  userStore.init()
})

// 监听搜索文本变化，使用防抖执行搜索
watch(searchText, (newValue) => {
  if (searchTimer) {
    clearTimeout(searchTimer)
  }
  searchTimer = setTimeout(() => {
    mcpStore.search(newValue)
  }, 500) // 500ms 防抖延迟
})

// 处理登出
const handleLogout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    userStore.logout()
    router.push('/login')
  }).catch(() => {
    // 取消退出
  })
}

// 跳转到登录页
const goToLogin = () => {
  router.push('/login')
}
</script>

<template>
  <el-container class="app-container">
    <el-header class="app-header">
      <div class="header-content">
        <div class="logo-section">
          <el-icon :size="32" color="#409EFF"><Grid /></el-icon>
          <h1>MCP 广场</h1>
        </div>
        <div class="header-actions">
          <el-input
            v-model="searchText"
            placeholder="搜索 MCP 服务..."
            style="width: 300px"
            clearable
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>

          <!-- 用户信息区域 -->
          <div v-if="userStore.isLoggedIn" class="user-section">
            <el-dropdown @command="handleLogout">
              <span class="user-info">
                <el-avatar :size="32" style="background-color: #409EFF;">
                  {{ userStore.username.charAt(0).toUpperCase() }}
                </el-avatar>
                <span class="username">{{ userStore.username }}</span>
                <el-icon class="el-icon--right"><ArrowDown /></el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item disabled>
                    <div style="font-size: 12px; color: #909399;">
                      {{ userStore.email }}
                    </div>
                  </el-dropdown-item>
                  <el-dropdown-item divided command="logout">
                    <el-icon><SwitchButton /></el-icon>
                    退出登录
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>

          <!-- 未登录显示登录按钮 -->
          <el-button v-else type="primary" @click="goToLogin">
            登录
          </el-button>
        </div>
      </div>
    </el-header>

    <el-main class="app-main">
      <router-view />
    </el-main>

    <el-footer class="app-footer" height="60px">
      <div class="footer-content">
        <p>MCP Market © 2025</p>
      </div>
    </el-footer>
  </el-container>
</template>

<style scoped>
.app-container {
  min-height: 100vh;
}

.app-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  padding: 0 40px;
  z-index: 1000;
}

.header-content {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-section h1 {
  margin: 0;
  font-size: 24px;
  color: #303133;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 20px;
}

.user-section {
  margin-left: 10px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 4px 12px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.user-info:hover {
  background-color: #f5f7fa;
}

.username {
  font-size: 14px;
  color: #303133;
  font-weight: 500;
}

.app-main {
  background: #f5f7fa;
  padding: 24px 40px;
  padding-top: 84px;
  min-height: calc(100vh - 60px);
}

.app-footer {
  background: #fff;
  border-top: 1px solid #e4e7ed;
  display: flex;
  align-items: center;
  justify-content: center;
}

.footer-content {
  text-align: center;
  color: #909399;
}

.footer-content p {
  margin: 0;
}
</style>
