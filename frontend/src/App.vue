<script setup>
import { useMcpStore } from '@/stores/mcp'
import { ref, watch } from 'vue'

const mcpStore = useMcpStore()
const searchText = ref('')

let searchTimer = null

// 监听搜索文本变化，使用防抖执行搜索
watch(searchText, (newValue) => {
  if (searchTimer) {
    clearTimeout(searchTimer)
  }
  searchTimer = setTimeout(() => {
    mcpStore.search(newValue)
  }, 500) // 500ms 防抖延迟
})
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
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  padding: 0 40px;
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

.app-main {
  background: #f5f7fa;
  padding: 24px 40px;
  min-height: calc(100vh - 120px);
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
