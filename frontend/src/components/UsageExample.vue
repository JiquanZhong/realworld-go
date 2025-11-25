<template>
  <div class="usage-example">
    <div class="example-content">
      <div class="example-header">
        <span class="example-label">Example {{ number }}</span>
      </div>
      <h3 class="example-title">{{ title }}</h3>
      <p class="example-description">{{ description }}</p>
    </div>
    <div v-if="imageUrl" class="example-visual">
      <img :src="imageUrl" :alt="title" class="example-image" />
    </div>
    <div v-else class="example-visual placeholder">
      <el-icon :size="64" color="var(--mcp-accent-cyan)">
        <TrendCharts />
      </el-icon>
    </div>
  </div>
</template>

<script setup>
import { TrendCharts } from '@element-plus/icons-vue'

defineProps({
  number: {
    type: Number,
    required: true
  },
  title: {
    type: String,
    required: true
  },
  description: {
    type: String,
    required: true
  },
  imageUrl: {
    type: String,
    default: ''
  }
})
</script>

<style scoped>
.usage-example {
  display: flex;
  gap: 24px;
  align-items: flex-start;
  padding: 20px 0;
  border-bottom: 1px solid var(--mcp-border);
}

.usage-example:last-child {
  border-bottom: none;
}

.example-content {
  flex: 1;
  min-width: 0;
}

.example-header {
  margin-bottom: 8px;
}

.example-label {
  font-size: 12px;
  font-weight: 600;
  color: var(--mcp-text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.example-title {
  margin: 0 0 8px;
  font-size: 18px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  line-height: 1.4;
}

.example-description {
  margin: 0;
  font-size: 14px;
  color: var(--mcp-text-secondary);
  line-height: 1.6;
}

.example-visual {
  width: 280px;
  height: 200px;
  flex-shrink: 0;
  border-radius: var(--radius-md);
  overflow: hidden;
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
}

.example-visual.placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
}

.example-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

@media (max-width: 768px) {
  .usage-example {
    flex-direction: column-reverse;
  }

  .example-visual {
    width: 100%;
  }
}
</style>
