<template>
  <div class="mcp-detail-page">
    <div class="page-container">
      <div v-if="loading" class="loading-container">
        <el-icon class="is-loading" :size="48" color="var(--mcp-accent-cyan)">
          <Loading />
        </el-icon>
      </div>

      <template v-else-if="mcp">
        <div class="top-bar">
          <el-button round text @click="handleGoMarket">
            <el-icon><ArrowLeft /></el-icon>
            返回市场
          </el-button>
        </div>
        <section class="hero-card">
          <div class="hero-main">
            <div class="title-row">
              <div class="hero-icon" :style="{ background: iconBg }">
                <img v-if="mcp.icon_url" :src="mcp.icon_url" :alt="mcp.name" />
                <span v-else>{{ initials }}</span>
              </div>
              <div>
                <p class="eyebrow">MCP 服务 · {{ formattedCreatedAt }}</p>
                <div class="name-row">
                  <h1 class="hero-title">{{ mcp.name }}</h1>
                  <div class="name-chips">
                    <span class="chip primary">MCP</span>
                    <span v-if="mcp.category" class="chip">{{ mcp.category }}</span>
                    <span v-for="tag in mcp.tags" :key="tag.id" class="chip subtle">#{{ tag.name }}</span>
                  </div>
                </div>
                <p class="hero-desc">
                  {{ mcp.description || defaultDescription }}
                </p>
              </div>
            </div>

            <div class="meta-grid">
              <div class="meta-item">
                <el-icon><TrendCharts /></el-icon>
                <div>
                  <p class="meta-label">安装量</p>
                  <p class="meta-value">{{ formattedInstallCount }}</p>
                </div>
              </div>
              <div class="meta-item">
                <el-icon><StarFilled /></el-icon>
                <div>
                  <p class="meta-label">评分</p>
                  <p class="meta-value">{{ formattedRating }}</p>
                </div>
              </div>
              <div class="meta-item">
                <el-icon><Timer /></el-icon>
                <div>
                  <p class="meta-label">最近更新</p>
                  <p class="meta-value">{{ formattedUpdatedAt }}</p>
                </div>
              </div>
              <div class="meta-item endpoint">
                <el-icon><Link /></el-icon>
                <div>
                  <p class="meta-label">Endpoint</p>
                  <p class="meta-value code">{{ mcp.endpoint || '尚未提供' }}</p>
                </div>
              </div>
            </div>
          </div>

          <div class="hero-actions">
            <div class="stat-block">
              <p class="stat-label">提交者</p>
              <p class="stat-value">{{ mcp.submitter_name || '未提供' }}</p>
            </div>
            <div class="action-buttons">
              <el-button round text @click="handleCopyEndpoint">
                <el-icon><DocumentCopy /></el-icon>
                复制 Endpoint
              </el-button>
              <el-button round text @click="handleShare">
                <el-icon><Share /></el-icon>
                分享
              </el-button>
              <el-button round type="primary" @click="scrollToContent">
                立即体验
              </el-button>
            </div>
          </div>
        </section>

        <section ref="detailSectionRef" class="doc-section">
          <div class="section-header">
            <div>
              <p class="section-eyebrow">服务详情</p>
              <h2 class="section-title">MCP {{ mcp.name }} 服务</h2>
              <p class="section-desc">
                {{ mcp.description || defaultDescription }}
              </p>
            </div>
          </div>
          <div class="markdown-body" v-html="renderedDetail" />
        </section>
      </template>

      <div v-else class="error-container">
        <el-empty description="未找到该 MCP" />
        <el-button type="primary" @click="$router.push('/mcps')">
          返回列表
        </el-button>
      </div>
    </div>

    <PageFooter />
  </div>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getMcpDetail } from '@/api/mcp'
import { ElMessage } from 'element-plus'
import MarkdownIt from 'markdown-it'
import { ArrowLeft, DocumentCopy, Link, Loading, Share, StarFilled, Timer, TrendCharts } from '@element-plus/icons-vue'
import PageFooter from '@/components/PageFooter.vue'

const route = useRoute()
const router = useRouter()
const mcp = ref(null)
const loading = ref(true)
const detailSectionRef = ref(null)

const md = new MarkdownIt({
  html: true,
  linkify: true,
  breaks: true
})

const defaultDescription = '探索高质量 MCP 服务，快速集成到你的智能体工作流中。'

const formattedInstallCount = computed(() => {
  if (!mcp.value?.install_count) return '0'
  const value = mcp.value.install_count
  if (value >= 1000) return `${(value / 1000).toFixed(1)}k`
  return `${value}`
})

const formattedRating = computed(() => {
  if (mcp.value?.rating_avg === 0 || mcp.value?.rating_avg) {
    return mcp.value.rating_avg.toFixed(1)
  }
  return '—'
})

const formattedCreatedAt = computed(() => formatDate(mcp.value?.created_at))
const formattedUpdatedAt = computed(() => formatDate(mcp.value?.updated_at))

const renderedDetail = computed(() => {
  const detail = mcp.value?.detail || mcp.value?.description || defaultDescription
  return md.render(detail || '')
})

const initials = computed(() => (mcp.value?.name ? mcp.value.name.slice(0, 1).toUpperCase() : 'M'))
const iconBg = computed(() => 'linear-gradient(135deg, rgba(99, 102, 241, 0.35), rgba(217, 70, 239, 0.25))')

const fetchMcpDetail = async () => {
  try {
    loading.value = true
    const id = route.params.id
    const data = await getMcpDetail(id)
    mcp.value = data
  } catch (error) {
    console.error('Failed to fetch MCP details:', error)
    ElMessage.error('获取 MCP 详情失败')
  } finally {
    loading.value = false
  }
}

const formatDate = (dateStr) => {
  if (!dateStr) return '—'
  const date = new Date(dateStr)
  if (Number.isNaN(date.getTime())) return '—'
  return date.toLocaleDateString('zh-CN')
}

const copyToClipboard = async (text, successMessage) => {
  if (!text) {
    ElMessage.warning('没有可复制的内容')
    return
  }
  try {
    await navigator.clipboard.writeText(text)
    ElMessage.success(successMessage)
  } catch (error) {
    console.error('Clipboard copy failed:', error)
    ElMessage.error('复制失败，请手动复制')
  }
}

const handleCopyEndpoint = () => {
  copyToClipboard(mcp.value?.endpoint, 'Endpoint 已复制')
}

const handleShare = () => {
  copyToClipboard(window.location.href, '页面链接已复制')
}

const handleGoMarket = () => {
  router.push('/mcps')
}

const scrollToContent = () => {
  if (detailSectionRef.value) {
    detailSectionRef.value.scrollIntoView({ behavior: 'smooth', block: 'start' })
  }
}

onMounted(() => {
  fetchMcpDetail()
})

watch(
  () => route.params.id,
  () => {
    fetchMcpDetail()
  }
)
</script>

<style scoped>
.mcp-detail-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: linear-gradient(180deg, rgba(29, 33, 45, 0.6) 0%, rgba(13, 17, 23, 1) 40%);
}

.page-container {
  max-width: 95vw;
  margin: 0 auto;
  width: 100%;
  flex: 1;
  padding: 10px 16px 48px;
}

.top-bar {
  display: flex;
  justify-content: flex-start;
  margin-bottom: 8px;
}

.top-bar :deep(.el-button) {
  padding-left: 6px;
  padding-right: 10px;
  color: var(--mcp-text-primary);
  border-color: transparent;
}

.top-bar :deep(.el-button:hover),
.top-bar :deep(.el-button:focus) {
  background: rgba(255, 255, 255, 0.06);
  color: var(--mcp-text-primary);
  border-color: var(--mcp-border);
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

.hero-card {
  background: rgba(21, 26, 38, 0.95);
  border: 1px solid var(--mcp-border);
  border-radius: 18px;
  padding: 20px 24px;
  display: grid;
  grid-template-columns: 3fr 1.2fr;
  gap: 18px;
  margin-bottom: 24px;
}

.hero-main {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.title-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.hero-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: grid;
  place-items: center;
  overflow: hidden;
  color: #fff;
  font-weight: 700;
  font-size: 20px;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.hero-icon img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.eyebrow {
  margin: 0;
  color: var(--mcp-text-muted);
  font-size: 12px;
  letter-spacing: 0.5px;
}

.hero-title {
  margin: 4px 0 0;
  font-size: 28px;
  color: var(--mcp-text-primary);
  font-weight: 800;
  letter-spacing: -0.3px;
}

.name-row {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 10px;
}

.name-chips {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.hero-desc {
  margin: 0;
  color: var(--mcp-text-secondary);
  line-height: 1.6;
}

.chip {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.05);
  color: var(--mcp-text-primary);
  font-size: 12px;
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.chip.primary {
  background: rgba(0, 212, 255, 0.16);
  border-color: rgba(0, 212, 255, 0.3);
  color: #7ce8ff;
}

.chip.subtle {
  color: var(--mcp-text-secondary);
}

.meta-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 12px;
}

.meta-item {
  display: flex;
  gap: 10px;
  align-items: center;
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
  border-radius: 12px;
  padding: 10px 12px;
}

.meta-item.endpoint {
  grid-column: span 2;
}

.meta-label {
  margin: 0;
  color: var(--mcp-text-muted);
  font-size: 12px;
}

.meta-value {
  margin: 2px 0 0;
  color: var(--mcp-text-primary);
  font-weight: 700;
  font-size: 14px;
}

.meta-value.code {
  font-family: 'JetBrains Mono', 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
  font-weight: 600;
  letter-spacing: 0.2px;
}

.hero-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
  border-radius: 14px;
  padding: 14px;
}

.stat-block {
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid var(--mcp-border);
  border-radius: 12px;
  padding: 10px 12px;
}

.stat-label {
  margin: 0;
  color: var(--mcp-text-muted);
  font-size: 12px;
}

.stat-value {
  margin: 4px 0 0;
  color: var(--mcp-text-primary);
  font-weight: 700;
}

.action-buttons {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  margin-top: auto;
}

.action-buttons :deep(.el-button) {
  border-radius: 999px;
}

.doc-section {
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
  border-radius: 16px;
  padding: 18px;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.section-eyebrow {
  margin: 0;
  color: var(--mcp-text-muted);
  font-size: 12px;
  letter-spacing: 0.4px;
}

.section-title {
  margin: 6px 0 10px;
  color: var(--mcp-text-primary);
  font-size: 22px;
  font-weight: 800;
}

.section-desc {
  margin: 0 0 12px;
  color: var(--mcp-text-secondary);
  line-height: 1.7;
}

.markdown-body {
  margin-top: 12px;
  color: var(--mcp-text-primary);
  line-height: 1.7;
}

.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4 {
  margin: 16px 0 8px;
  color: var(--mcp-text-primary);
}

.markdown-body p {
  margin: 0 0 12px;
  color: var(--mcp-text-secondary);
}

.markdown-body pre {
  background: rgba(0, 0, 0, 0.35);
  border: 1px solid var(--mcp-border);
  border-radius: 12px;
  padding: 12px;
  color: var(--mcp-text-primary);
  overflow: auto;
  font-family: 'JetBrains Mono', 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
}

.markdown-body code {
  background: rgba(255, 255, 255, 0.06);
  padding: 2px 6px;
  border-radius: 6px;
  font-family: 'JetBrains Mono', 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
}

.markdown-body ul,
.markdown-body ol {
  padding-left: 20px;
  margin: 0 0 12px;
}

.markdown-body a {
  color: var(--mcp-accent-cyan);
}

.markdown-body blockquote {
  border-left: 3px solid var(--mcp-border-light);
  padding-left: 12px;
  margin: 0 0 12px;
  color: var(--mcp-text-secondary);
}

@media (max-width: 1024px) {
  .hero-card {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .hero-card,
  .doc-section {
    padding: 14px;
  }

  .meta-item.endpoint {
    grid-column: span 1;
  }
}
</style>
