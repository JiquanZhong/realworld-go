<template>
  <div class="mcp-card" @click="handleClick">
    <div class="card-body">
      <div class="card-icon-wrapper">
        <img
          v-if="mcp.icon_url"
          :src="mcp.icon_url"
          :alt="mcp.name"
          class="card-icon"
          loading="lazy"
        />
        <div v-else class="card-icon-placeholder">
          <el-icon :size="36" color="var(--mcp-accent-cyan)"><Box /></el-icon>
        </div>
      </div>

      <div class="card-main">
        <div class="card-header">
          <h3 class="mcp-name">{{ mcp.name }}</h3>
          <span v-if="mcp.category" class="mcp-category">{{ mcp.category }}</span>
        </div>
        <p class="mcp-description">{{ mcp.description || '暂无描述' }}</p>

        <div class="card-stats">
          <div class="stat-item">
            <span class="stat-label">安装量</span>
            <span class="stat-value">{{ formatNumber(mcp.install_count) }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">评分</span>
            <span class="stat-value">{{ formatRating(mcp.rating_avg) }}</span>
          </div>
        </div>

        <div v-if="displayTags.length" class="card-tags">
          <span v-for="tag in displayTags" :key="tag.id" class="tag-chip">{{ tag.name }}</span>
        </div>
      </div>
    </div>
  </div>
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

const formatNumber = (num = 0) => {
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'k'
  }
  return num.toString()
}

const formatRating = (value) => {
  if (value === null || value === undefined || value <= 0) return 'N/A'
  return Number(value).toFixed(1)
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
  border-radius: var(--radius-md);
  overflow: hidden;
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
}

.mcp-card:hover {
  transform: translateY(-4px);
  border-color: var(--mcp-accent-cyan);
  box-shadow: 0 8px 24px rgba(0, 212, 255, 0.15);
}

.card-body {
  display: flex;
  gap: 16px;
  padding: 16px;
  align-items: flex-start;
}

.card-icon-wrapper {
  flex: 0 0 72px;
  height: 72px;
  border-radius: var(--radius-md);
  overflow: hidden;
  background: rgba(0, 212, 255, 0.08);
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-icon {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.card-icon-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 8px;
}

.mcp-name {
  margin: 0;
  font-size: 16px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  font-family: 'Space Grotesk', sans-serif;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.mcp-category {
  font-size: 12px;
  color: var(--mcp-text-secondary);
  border: 1px solid var(--mcp-border);
  border-radius: 999px;
  padding: 2px 8px;
}

.mcp-description {
  margin: 0;
  font-size: 13px;
  color: var(--mcp-text-secondary);
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  flex: 1;
}

.card-stats {
  display: flex;
  gap: 16px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  gap: 2px;
  min-width: 80px;
}

.stat-label {
  font-size: 12px;
  color: var(--mcp-text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 14px;
  color: var(--mcp-text-primary);
  font-weight: 600;
}

.card-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.tag-chip {
  font-size: 11px;
  padding: 2px 8px;
  border-radius: 999px;
  background: rgba(0, 212, 255, 0.1);
  color: var(--mcp-accent-cyan);
}

@media (max-width: 640px) {
  .card-body {
    flex-direction: column;
    align-items: flex-start;
  }

  .card-icon-wrapper {
    width: 64px;
    height: 64px;
  }
}
</style>
