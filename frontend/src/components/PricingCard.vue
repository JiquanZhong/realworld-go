<template>
  <div class="pricing-card">
    <div class="pricing-header">
      <h3 class="plan-name">{{ planName }}</h3>
      <div class="price-container">
        <span class="price">{{ price }}</span>
        <span v-if="period" class="period">/{{ period }}</span>
      </div>
    </div>

    <el-button class="cta-button" type="primary" @click="$emit('select')">
      {{ buttonText }}
    </el-button>

    <div class="features-list">
      <div
        v-for="(feature, index) in features"
        :key="index"
        class="feature-item"
      >
        <el-icon class="check-icon" :size="20">
          <Check />
        </el-icon>
        <span class="feature-text">{{ feature }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Check } from '@element-plus/icons-vue'

defineProps({
  planName: {
    type: String,
    required: true
  },
  price: {
    type: String,
    required: true
  },
  period: {
    type: String,
    default: 'month'
  },
  buttonText: {
    type: String,
    default: 'Get Started'
  },
  features: {
    type: Array,
    default: () => []
  }
})

defineEmits(['select'])
</script>

<style scoped>
.pricing-card {
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border-light);
  border-radius: var(--radius-sm);
  padding: 25px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  max-width: 480px;
  transition: all 0.3s;
}

.pricing-card:hover {
  border-color: var(--mcp-accent-cyan);
  box-shadow: 0 0 20px rgba(0, 212, 255, 0.15);
}

.pricing-header {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.plan-name {
  margin: 0;
  font-size: 16px;
  font-weight: 700;
  color: var(--mcp-text-primary);
}

.price-container {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.price {
  font-size: 36px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  letter-spacing: -1px;
  line-height: 1.25;
}

.period {
  font-size: 16px;
  font-weight: 700;
  color: var(--mcp-text-primary);
}

.cta-button {
  width: 100%;
  background: #242e47;
  border-color: #242e47;
  color: var(--mcp-text-primary);
  font-weight: 700;
  font-size: 14px;
  height: 40px;
  border-radius: var(--radius-sm);
}

.cta-button:hover {
  background: #2d3a54;
  border-color: #2d3a54;
}

.features-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.feature-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.check-icon {
  color: var(--mcp-accent-cyan);
  flex-shrink: 0;
  margin-top: 2px;
}

.feature-text {
  font-size: 13px;
  color: var(--mcp-text-primary);
  line-height: 1.5;
}
</style>
