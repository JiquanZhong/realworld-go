<template>
  <div class="feature-card">
    <div class="feature-icon">
      <el-icon :size="24">
        <component :is="icon" />
      </el-icon>
    </div>
    <div class="feature-content">
      <h3 class="feature-title">{{ title }}</h3>
      <p v-if="description" class="feature-description">{{ description }}</p>
    </div>
  </div>
</template>

<script setup>
defineProps({
  icon: {
    type: String,
    required: true
  },
  title: {
    type: String,
    required: true
  },
  description: {
    type: String,
    default: ''
  }
})
</script>

<style scoped>
.feature-card {
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border-light);
  border-radius: var(--radius-sm);
  padding: 17px;
  display: flex;
  align-items: center;
  gap: 12px;
  transition: all 0.3s;
}

.feature-card:hover {
  background: var(--mcp-dark-bg-card-hover);
  border-color: var(--mcp-accent-cyan);
}

.feature-icon {
  color: var(--mcp-accent-cyan);
  flex-shrink: 0;
}

.feature-content {
  flex: 1;
  min-width: 0;
}

.feature-title {
  margin: 0;
  font-size: 16px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  line-height: 1.25;
}

.feature-description {
  margin: 4px 0 0;
  font-size: 14px;
  color: var(--mcp-text-secondary);
  line-height: 1.5;
}
</style>
