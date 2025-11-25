<template>
  <div class="mcp-market">
    <div class="market-header">
      <h1 class="page-title">探索MCP</h1>
      <p class="page-subtitle">聚合优质MCP资源，拓展模型智能边界</p>
    </div>

    <div class="search-bar">
      <el-input
        v-model="searchInput"
        placeholder="搜索"
        class="search-input"
        clearable
      >
        <template #prefix>
          <el-icon><Search /></el-icon>
        </template>
      </el-input>
    </div>

    <div class="filter-section">
      <div class="filter-row">
        <el-select v-model="category" placeholder="分类" class="filter-select">
          <el-option label="所有分类" value="" />
          <el-option label="生产力" value="productivity" />
          <el-option label="分析" value="analytics" />
          <el-option label="沟通" value="communication" />
        </el-select>

        <el-select v-model="rating" placeholder="评分" class="filter-select">
          <el-option label="所有评分" value="" />
          <el-option label="5 星" value="5" />
          <el-option label="4+星" value="4" />
          <el-option label="3+星" value="3" />
        </el-select>

        <el-select v-model="price" placeholder="价格" class="filter-select">
          <el-option label="所有价格" value="" />
          <el-option label="免费" value="free" />
          <el-option label="Paid" value="paid" />
        </el-select>

        <el-select
          v-model="sortOption"
          placeholder="Sort by"
          class="filter-select"
          @change="handleSortChange"
        >
          <el-option label="安装量最多" value="install_count_desc" />
          <el-option label="安装量最少" value="install_count_asc" />
          <el-option label="评分最高" value="rating_avg_desc" />
          <el-option label="评分最低" value="rating_avg_asc" />
          <el-option label="最新发布" value="created_at_desc" />
          <el-option label="最旧发布" value="created_at_asc" />
        </el-select>
      </div>
    </div>

    <div v-loading="mcpStore.loading" class="content-section">
      <div v-if="mcpStore.mcpList.length === 0 && !mcpStore.loading" class="empty-state">
        <el-empty description="No MCPs found" />
      </div>

      <div v-else class="mcp-grid">
        <mcp-card
          v-for="mcp in mcpStore.mcpList"
          :key="mcp.id"
          :mcp="mcp"
          @click="handleMcpClick"
        />
      </div>
    </div>

    <div v-if="mcpStore.total > mcpStore.pageSize" class="pagination-section">
      <el-pagination
        v-model:current-page="mcpStore.currentPage"
        :page-size="mcpStore.pageSize"
        :total="mcpStore.total"
        layout="prev, pager, next"
        background
        @current-change="handlePageChange"
      />
    </div>

    <PageFooter />
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useMcpStore } from '@/stores/mcp'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import McpCard from '@/components/McpCard.vue'
import PageFooter from '@/components/PageFooter.vue'

const router = useRouter()
const mcpStore = useMcpStore()
const sortOption = ref('install_count_desc')
const searchInput = ref('')
const category = ref('')
const rating = ref('')
const price = ref('')
let searchTimeout = null

onMounted(() => {
  loadData()
})

watch(searchInput, (newValue) => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    mcpStore.search(newValue)
  }, 500)
})

const loadData = async () => {
  try {
    await mcpStore.fetchMcpList()
  } catch (error) {
    ElMessage.error('Failed to load MCP list')
  }
}

const handleSortChange = (value) => {
  const lastUnderscoreIndex = value.lastIndexOf('_')
  const field = value.substring(0, lastUnderscoreIndex)
  const order = value.substring(lastUnderscoreIndex + 1)
  const isDesc = order === 'desc'
  mcpStore.changeSort(field, !isDesc)
}

const handlePageChange = (page) => {
  mcpStore.changePage(page)
}

const handleMcpClick = (mcp) => {
  router.push(`/mcps/${mcp.id}`)
}
</script>

<style scoped>
.mcp-market {
  max-width: 80vw;
  margin: 0 auto;
  padding-bottom: 40px;
}

.market-header {
  padding: 20px 16px;
  margin-bottom: 24px;
}

.page-title {
  margin: 0 0 12px;
  font-size: 36px;
  font-weight: 700;
  color: var(--mcp-text-primary);
  font-family: 'Space Grotesk', sans-serif;
}

.page-subtitle {
  margin: 0;
  font-size: 16px;
  color: var(--mcp-text-secondary);
  line-height: 1.5;
  max-width: 720px;
}

.search-bar {
  padding: 0 16px;
  margin-bottom: 20px;
}

.search-input {
  max-width: 600px;
}

.search-input :deep(.el-input__wrapper) {
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
  box-shadow: none;
}

.search-input :deep(.el-input__wrapper:hover) {
  border-color: var(--mcp-border-light);
}

.search-input :deep(.el-input__wrapper.is-focus) {
  border-color: var(--mcp-accent-cyan);
}

.search-input :deep(.el-input__inner) {
  color: var(--mcp-text-primary);
}

.search-input :deep(.el-input__inner::placeholder) {
  color: var(--mcp-text-muted);
}

.search-input :deep(.el-input__prefix) {
  color: var(--mcp-text-muted);
}

.filter-section {
  padding: 0 16px;
  margin-bottom: 32px;
}

.filter-row {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.filter-select {
  width: 160px;
  color: red;
}

.filter-select :deep(.el-select__wrapper) {
  background-color: var(--mcp-dark-bg-card) !important;
  color: white !important;
  font-size: 14px;
}

.filter-select :deep(.is-hovering) {
  background-color: var(--mcp-dark-bg-card) !important;
}

.filter-select :deep(.el-input__wrapper) {
  background: var(--mcp-dark-bg-card) !important;
  border: 1px solid var(--mcp-border);
  box-shadow: none;
}

.filter-select :deep(.el-input__wrapper:hover) {
  border-color: var(--mcp-border-light);
}

.filter-select :deep(.el-input__wrapper.is-focus) {
  border-color: var(--mcp-accent-cyan);
}

.filter-select :deep(.el-input__inner) {
  color: var(--mcp-text-primary) !important;
}

.filter-select :deep(.el-input__inner::placeholder) {
  color: var(--mcp-text-muted) !important;
}

.filter-select :deep(.el-select__placeholder) {
  color: var(--mcp-text-muted) !important;
}

.filter-select :deep(.el-input__suffix) {
  color: var(--mcp-text-muted);
}

.content-section {
  min-height: 400px;
  padding: 0 16px;
}

.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.mcp-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 16px;
  margin-bottom: 32px;
}

.pagination-section {
  display: flex;
  justify-content: center;
  padding: 32px 16px;
}

.pagination-section :deep(.el-pagination) {
  gap: 8px;
}

.pagination-section :deep(.el-pager li) {
  background: var(--mcp-dark-bg-card) !important;
  border-color: var(--mcp-border);
  color: var(--mcp-text-secondary);
}

.pagination-section :deep(.el-pager li:hover) {
  color: var(--mcp-accent-cyan) !important;
  border-color: var(--mcp-accent-cyan);
  background: var(--mcp-dark-bg-card-hover) !important;
}

.pagination-section :deep(.el-pager li.is-active) {
  background: var(--mcp-accent-cyan) !important;
  color: var(--mcp-dark-bg) !important;
  border-color: var(--mcp-accent-cyan);
}

.pagination-section :deep(.btn-prev),
.pagination-section :deep(.btn-next) {
  background: var(--mcp-dark-bg-card) !important;
  border-color: var(--mcp-border);
  color: var(--mcp-text-secondary);
}

.pagination-section :deep(.btn-prev:hover),
.pagination-section :deep(.btn-next:hover) {
  color: var(--mcp-accent-cyan) !important;
  border-color: var(--mcp-accent-cyan);
  background: var(--mcp-dark-bg-card-hover) !important;
}

.pagination-section :deep(.btn-prev:disabled),
.pagination-section :deep(.btn-next:disabled) {
  background: var(--mcp-dark-bg-card) !important;
  color: var(--mcp-text-disabled) !important;
  border-color: var(--mcp-border);
}

@media (max-width: 1024px) {
  .mcp-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .page-title {
    font-size: 28px;
  }

  .mcp-grid {
    grid-template-columns: 1fr;
  }

  .filter-row {
    flex-direction: column;
  }

  .filter-select {
    width: 100%;
  }

  .search-input {
    max-width: 100%;
  }
}
</style>
