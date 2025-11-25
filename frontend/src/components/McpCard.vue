<template>
  <div class="mcp-card" @click="handleClick">
    <div class="card-image-wrapper">
      <img v-if="mcp.icon_url" :src="mcp.icon_url" :alt="mcp.name" class="card-image" />
      <div v-else class="card-image-placeholder">
        <el-icon :size="64" color="var(--mcp-accent-cyan)"><Box /></el-icon>
      </div>
    </div>
    <div class="card-content">
      <h3 class="mcp-name">{{ mcp.name }}</h3>
      <p class="mcp-description">{{ mcp.description || '暂无描述' }}</p>
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

.card-image-wrapper {
  width: 100%;
  height: 160px;
  overflow: hidden;
  background: linear-gradient(135deg, #1a2536 0%, #0d1117 100%);
  position: relative;
}

.card-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.card-image-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(0, 212, 255, 0.1) 0%, rgba(59, 130, 246, 0.1) 100%);
}

.card-content {
  flex: 1;
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.mcp-name {
  margin: 0;
  font-size: 14px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'Space Grotesk', sans-serif;
}

.mcp-description {
  margin: 0;
  font-size: 12px;
  color: var(--mcp-text-secondary);
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  flex: 1;
}
</style>
