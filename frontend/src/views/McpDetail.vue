<template>
  <div class="mcp-detail-page">
    <div class="page-container">
      <div class="page-header">
        <h1 class="page-title">MCP Details</h1>
        <p class="page-subtitle">Explore the capabilities and usage of this MCP</p>
      </div>

      <div v-if="loading" class="loading-container">
        <el-icon class="is-loading" :size="48" color="var(--mcp-accent-cyan)">
          <Loading />
        </el-icon>
      </div>

      <div v-else-if="mcp" class="detail-content">
        <!-- Tabs -->
        <el-tabs v-model="activeTab" class="detail-tabs">
          <el-tab-pane label="Overview" name="overview">
            <div class="tab-content">
              <!-- Overview Section -->
              <section class="content-section">
                <h2 class="section-title">Overview</h2>
                <p class="section-text">
                  {{ mcp.description || 'This MCP provides a comprehensive suite of tools for managing and optimizing your digital operations. It integrates seamlessly with various platforms, offering features like automated processing, performance tracking, and intelligent optimization. The interface is designed with a futuristic, tech-forward aesthetic, incorporating modern design elements and dynamic visualizations to enhance the user experience.' }}
                </p>
              </section>

              <!-- Features Section -->
              <section class="content-section">
                <h2 class="section-title">Features</h2>
                <div class="features-grid">
                  <FeatureCard
                    icon="PieChart"
                    title="Data Visualization"
                  />
                  <FeatureCard
                    icon="TrendCharts"
                    title="Trend Identification"
                  />
                  <FeatureCard
                    icon="Document"
                    title="Report Generation"
                  />
                </div>
              </section>

              <!-- Usage Examples Section -->
              <section class="content-section">
                <h2 class="section-title">Usage Examples</h2>
                <div class="examples-list">
                  <UsageExample
                    :number="1"
                    title="Creating a New Campaign"
                    description="Learn how to quickly set up a new marketing campaign using the MCP's intuitive interface, now with a futuristic, tech-inspired design."
                  />
                  <UsageExample
                    :number="2"
                    title="Analyzing Campaign Performance"
                    description="Discover how to use the MCP's analytics dashboard to monitor and improve your campaign results, presented in a sleek, tech-forward style."
                  />
                  <UsageExample
                    :number="3"
                    title="Adjusting Budget Allocation"
                    description="Find out how to optimize your budget by reallocating funds based on campaign performance, all within a modern, tech-themed interface."
                  />
                </div>
              </section>

              <!-- Pricing Section -->
              <section class="content-section">
                <h2 class="section-title">Pricing</h2>
                <div class="pricing-grid">
                  <PricingCard
                    plan-name="Basic Plan"
                    price="$49"
                    period="month"
                    button-text="Choose Basic"
                    :features="[
                      'Up to 5 campaigns',
                      'Basic analytics',
                      'Email support'
                    ]"
                    @select="handlePricingSelect('basic')"
                  />
                  <PricingCard
                    plan-name="Pro Plan"
                    price="$99"
                    period="month"
                    button-text="Choose Pro"
                    :features="[
                      'Up to 20 campaigns',
                      'Advanced analytics',
                      'Priority support'
                    ]"
                    @select="handlePricingSelect('pro')"
                  />
                  <PricingCard
                    plan-name="Enterprise"
                    price="Custom"
                    period="month"
                    button-text="Contact Us"
                    :features="[
                      'Unlimited campaigns',
                      'Custom analytics',
                      'Dedicated account manager'
                    ]"
                    @select="handlePricingSelect('enterprise')"
                  />
                </div>
              </section>
            </div>
          </el-tab-pane>

          <el-tab-pane label="Examples" name="examples">
            <div class="tab-content">
              <p class="section-text text-secondary">Examples content coming soon...</p>
            </div>
          </el-tab-pane>

          <el-tab-pane label="Reviews" name="reviews">
            <div class="tab-content">
              <p class="section-text text-secondary">Reviews content coming soon...</p>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>

      <div v-else class="error-container">
        <el-empty description="MCP not found" />
        <el-button type="primary" @click="$router.push('/mcps')">
          Back to MCPs
        </el-button>
      </div>
    </div>

    <PageFooter />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getMcpDetail } from '@/api/mcp'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'
import FeatureCard from '@/components/FeatureCard.vue'
import PricingCard from '@/components/PricingCard.vue'
import UsageExample from '@/components/UsageExample.vue'
import PageFooter from '@/components/PageFooter.vue'

const route = useRoute()
const mcp = ref(null)
const loading = ref(true)
const activeTab = ref('overview')

const fetchMcpDetail = async () => {
  try {
    loading.value = true
    const id = route.params.id
    const data = await getMcpDetail(id)
    mcp.value = data
  } catch (error) {
    console.error('Failed to fetch MCP details:', error)
    ElMessage.error('Failed to load MCP details')
  } finally {
    loading.value = false
  }
}

const handlePricingSelect = (plan) => {
  ElMessage.success(`You selected ${plan} plan`)
}

onMounted(() => {
  fetchMcpDetail()
})
</script>

<style scoped>
.mcp-detail-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.page-container {
  max-width: 960px;
  margin: 0 auto;
  width: 100%;
  flex: 1;
  padding-bottom: 40px;
}

.page-header {
  padding: 20px 16px;
  margin-bottom: 20px;
}

.page-title {
  margin: 0 0 8px;
  font-size: 32px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  font-family: 'Space Grotesk', sans-serif;
}

.page-subtitle {
  margin: 0;
  font-size: 16px;
  color: var(--mcp-text-secondary);
}

.loading-container,
.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  gap: 24px;
}

.detail-content {
  background: var(--mcp-dark-bg-elevated);
  border-radius: var(--radius-md);
  overflow: hidden;
}

.detail-tabs {
  padding: 0;
}

.detail-tabs :deep(.el-tabs__header) {
  margin: 0;
  background: transparent;
  border-bottom: 1px solid var(--mcp-border-light);
  padding: 0 16px;
}

.detail-tabs :deep(.el-tabs__nav-wrap::after) {
  display: none;
}

.detail-tabs :deep(.el-tabs__item) {
  color: var(--mcp-text-secondary);
  font-weight: 700;
  font-size: 14px;
  padding: 16px 0;
  margin-right: 32px;
}

.detail-tabs :deep(.el-tabs__item.is-active) {
  color: var(--mcp-text-primary);
}

.detail-tabs :deep(.el-tabs__active-bar) {
  background-color: var(--mcp-text-primary);
  height: 3px;
}

.tab-content {
  padding: 0;
}

.content-section {
  padding: 20px 16px 12px;
}

.content-section + .content-section {
  padding-top: 20px;
}

.section-title {
  margin: 0 0 16px;
  font-size: 22px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  font-family: 'Space Grotesk', sans-serif;
}

.section-text {
  margin: 0;
  font-size: 16px;
  color: var(--mcp-text-primary);
  line-height: 1.6;
}

.section-text.text-secondary {
  color: var(--mcp-text-secondary);
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 12px;
}

.examples-list {
  display: flex;
  flex-direction: column;
}

.pricing-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 10px;
}

@media (max-width: 768px) {
  .page-container {
    padding-bottom: 20px;
  }

  .page-title {
    font-size: 24px;
  }

  .features-grid,
  .pricing-grid {
    grid-template-columns: 1fr;
  }

  .detail-tabs :deep(.el-tabs__item) {
    margin-right: 16px;
  }
}
</style>
