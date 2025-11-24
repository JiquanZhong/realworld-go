<template>
  <el-card class="mcp-card" shadow="hover" @click="handleClick">
    <div class="card-header">
      <div class="icon-wrapper">
        <img v-if="mcp.icon_url" :src="mcp.icon_url" :alt="mcp.name" class="mcp-icon" />
        <el-icon v-else :size="48" color="#909399"><Box /></el-icon>
      </div>
      <div class="header-info">
        <h3 class="mcp-name">{{ mcp.name }}</h3>
        <div class="mcp-category">
          <el-tag size="small" type="info">{{ mcp.category || '未分类' }}</el-tag>
        </div>
      </div>
    </div>

    <div class="card-body">
      <p class="mcp-description">{{ mcp.description || '暂无描述' }}</p>
    </div>

    <div class="card-footer">
      <div class="tags-section">
        <el-tag
          v-for="tag in displayTags"
          :key="tag.id"
          size="small"
          effect="plain"
          class="tag-item"
        >
          {{ tag.name }}
        </el-tag>
        <span v-if="mcp.tags && mcp.tags.length > 3" class="more-tags">
          +{{ mcp.tags.length - 3 }}
        </span>
      </div>

      <div class="stats-section">
        <div class="stat-item">
          <el-icon><Download /></el-icon>
          <span>{{ formatNumber(mcp.install_count || 0) }}</span>
        </div>
        <div class="stat-item">
          <el-icon><Star /></el-icon>
          <span>{{ mcp.rating_avg ? mcp.rating_avg.toFixed(1) : '0.0' }}</span>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  mcp: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['click'])

const displayTags = computed(() => {
  if (!props.mcp.tags || props.mcp.tags.length === 0) return []
  return props.mcp.tags.slice(0, 3)
})

const formatNumber = (num) => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'k'
  }
  return num.toString()
}

const handleClick = () => {
  emit('click', props.mcp)
}
</script>

<style scoped>
.mcp-card {
  cursor: pointer;
  transition: all 0.3s;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.mcp-card:hover {
  transform: translateY(-4px);
}

.card-header {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
}

.icon-wrapper {
  flex-shrink: 0;
  width: 64px;
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f7fa;
  border-radius: 8px;
  overflow: hidden;
}

.mcp-icon {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.header-info {
  flex: 1;
  min-width: 0;
}

.mcp-name {
  margin: 0 0 8px 0;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.mcp-category {
  display: flex;
  align-items: center;
}

.card-body {
  flex: 1;
  margin-bottom: 16px;
}

.mcp-description {
  margin: 0;
  font-size: 14px;
  color: #606266;
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-footer {
  border-top: 1px solid #e4e7ed;
  padding-top: 12px;
}

.tags-section {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 12px;
  min-height: 24px;
}

.tag-item {
  font-size: 12px;
}

.more-tags {
  font-size: 12px;
  color: #909399;
}

.stats-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #909399;
  font-size: 14px;
}

.stat-item .el-icon {
  font-size: 16px;
}
</style>
