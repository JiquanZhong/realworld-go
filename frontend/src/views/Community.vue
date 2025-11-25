<template>
  <div class="placeholder-page">
    <div class="placeholder-content">
      <el-icon :size="80" color="var(--mcp-accent-cyan)">
        <ChatDotRound />
      </el-icon>
      <h1>Community</h1>
      <p>Coming Soon</p>
    </div>
  </div>
</template>

<script setup>
import { ChatDotRound } from '@element-plus/icons-vue'
</script>

<style scoped>
.placeholder-page {
  min-height: calc(100vh - 120px);
  display: flex;
  align-items: center;
  justify-content: center;
}

.placeholder-content {
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.placeholder-content h1 {
  margin: 0;
  font-size: 36px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  font-family: 'Space Grotesk', sans-serif;
}

.placeholder-content p {
  margin: 0;
  font-size: 18px;
  color: var(--mcp-text-secondary);
}
</style>
