<script setup>
import { useUserStore } from '@/stores/user'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import McpCard from '@/components/McpCard.vue';

const router = useRouter()
const userStore = useUserStore()

// 初始化用户状态
onMounted(() => {
  userStore.init()
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
          <el-icon :size="28" color="#00d4ff"><Grid /></el-icon>
          <h1>MCP Hub</h1>
        </div>

        <nav class="nav-links">
          <router-link to="/" class="nav-link">首页</router-link>
          <router-link to="/mcps" class="nav-link">MCP</router-link>
          <router-link to="/docs" class="nav-link">文档</router-link>
          <router-link to="/community" class="nav-link">社区</router-link>
        </nav>

        <div class="header-actions">
          <!-- 用户信息区域 -->
          <div v-if="userStore.isLoggedIn" class="user-section">
            <el-dropdown @command="handleLogout">
              <span class="user-info">
                <el-avatar v-if="userStore.iconUrl" :size="32" :src="userStore.iconUrl" />
                <el-avatar v-else :size="32" class="user-avatar">
                  {{ userStore.username.charAt(0).toUpperCase() }}
                </el-avatar>
              </span>
              <template #dropdown>
                <el-dropdown-menu class="user-dropdown">
                  <el-dropdown-item disabled>
                    <div class="dropdown-email">
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
          <el-button v-else class="login-btn" @click="goToLogin">
            登录
          </el-button>
        </div>
      </div>
    </el-header>

    <el-main class="app-main">
      <router-view />
    </el-main>
  </el-container>
</template>

<style scoped>


.app-container {
  min-height: 100vh;
  background: var(--mcp-dark-bg);
}

.app-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: var(--mcp-dark-bg-elevated);
  border-bottom: 1px solid var(--mcp-border);
  display: flex;
  align-items: center;
  padding: 0 40px;
  z-index: 1000;
  height: 60px;
}

.header-content {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 32px;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-shrink: 0;
}

.logo-section h1 {
  margin: 0;
  padding: 0;

  font-size: 20px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  font-family: 'Space Grotesk', sans-serif;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 32px;
  flex: 1;
}

.nav-link {
  color: var(--mcp-text-secondary);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  padding: 8px 12px;
  border-radius: var(--radius-sm);
  transition: all 0.3s;
}

.nav-link:hover {
  color: var(--mcp-text-primary);
  background: rgba(255, 255, 255, 0.05);
}

.nav-link.router-link-active {
  color: var(--mcp-accent-cyan);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
  flex-shrink: 0;
  margin-left: auto;
}

.user-section {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
  border-radius: var(--radius-full);
  transition: all 0.3s;
}

.user-info:hover {
  opacity: 0.8;
}

.user-avatar {
  background: linear-gradient(135deg, var(--mcp-accent-cyan), var(--mcp-accent-blue));
  color: var(--mcp-text-primary);
  font-weight: 600;
}

.dropdown-email {
  font-size: 12px;
  color: var(--mcp-text-muted);
}

.login-btn {
  background: var(--mcp-accent-cyan);
  border-color: var(--mcp-accent-cyan);
  color: var(--mcp-dark-bg);
  font-weight: 600;
}

.login-btn:hover {
  background: var(--mcp-accent-blue);
  border-color: var(--mcp-accent-blue);
}

.app-main {
  background: var(--mcp-dark-bg);
  padding: 24px 40px;
  padding-top: 84px;
  min-height: 100vh;
}

@media (max-width: 768px) {
  .nav-links {
    display: none;
  }

  .header-content {
    gap: 16px;
  }

  .search-input {
    width: 180px;
  }

  .app-main {
    padding: 20px 16px;
    padding-top: 80px;
  }
}
</style>
