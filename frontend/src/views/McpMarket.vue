<template>
  <div class="mcp-market">
    <div class="market-header">
      <div class="header-left">
        <h1 class="page-title">探索MCP</h1>
        <p class="page-subtitle">聚合优质MCP资源，拓展模型智能边界</p>
      </div>

      <div class="header-controls">
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

        <el-select
          v-model="sortOption"
          placeholder="排序"
          class="sort-select"
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

    <div v-loading="mcpStore.loading" class="content-layout">
      <aside class="tag-sidebar">
        <div class="sidebar-header">
          <h3>标签</h3>
          <p>选择一个标签探索相关 MCP</p>
        </div>
        <div class="tag-list">
          <button
            class="tag-item"
            :class="{ active: !selectedTag }"
            type="button"
            @click="handleTagSelect(null)"
          >
            全部
          </button>
          <button
            v-for="tag in mcpStore.tags"
            :key="tag.id"
            class="tag-item"
            :class="{ active: selectedTag === tag.id }"
            type="button"
            @click="handleTagSelect(tag.id)"
          >
            {{ tag.name }}
          </button>
        </div>
      </aside>

      <div class="market-content">
        <div v-if="filteredMcpList.length === 0 && !mcpStore.loading" class="empty-state">
          <el-empty description="没找到符合条件的MCP" />
        </div>

        <div v-else class="mcp-grid">
          <mcp-card
            v-for="mcp in filteredMcpList"
            :key="mcp.id"
            :mcp="mcp"
            @click="handleMcpClick"
          />
        </div>
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
import { ref, onMounted, watch, computed } from 'vue'
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
const selectedTag = ref(null)
let searchTimeout = null

const filteredMcpList = computed(() => {
  if (!selectedTag.value) return mcpStore.mcpList
  return mcpStore.mcpList.filter((mcp) =>
    (mcp.tags || []).some((tag) => tag.id === selectedTag.value)
  )
})

onMounted(() => {
  loadData()
  loadTags()
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

const loadTags = async () => {
  try {
    await mcpStore.fetchTags()
  } catch (error) {
    console.error('Failed to load tags', error)
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

const handleTagSelect = (tagId) => {
  if (selectedTag.value === tagId) {
    selectedTag.value = null
    return
  }
  selectedTag.value = tagId
}
</script>

<style scoped>
.mcp-market {
  max-width: 80vw;
  margin: 0 auto;
  padding-bottom: 40px;
}

.market-header {
  padding: 0px 16px;
  margin-bottom: 24px;
  display: flex;
  justify-content: space-between;
  gap: 24px;
  align-items: flex-start;
  flex-wrap: wrap;
}

.header-left {
  flex: 1;
  min-width: 260px;
}

.header-controls {
  display: flex;
  gap: 12px;
  align-items: center;
  justify-content: flex-end;
  min-width: 320px;
  flex-wrap: wrap;
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

.search-input {
  width: 280px;
  min-width: 200px;
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

.sort-select {
  width: 180px;
}

.sort-select :deep(.el-input__wrapper) {
  background: var(--mcp-dark-bg-card) !important;
  border: 1px solid var(--mcp-border);
  box-shadow: none;
}

.sort-select :deep(.el-input__wrapper:hover) {
  border-color: var(--mcp-border-light);
}

.sort-select :deep(.el-input__wrapper.is-focus) {
  border-color: var(--mcp-accent-cyan);
}

.sort-select :deep(.el-select__wrapper) {
  background: var(--mcp-dark-bg-card) !important;
  border: 1px solid var(--mcp-border);
  box-shadow: none;
}

.sort-select :deep(.el-select__wrapper:hover) {
  border-color: var(--mcp-border-light);
}

.sort-select :deep(.el-select__wrapper.is-hovering),
.sort-select :deep(.el-select__wrapper.is-focused) {
  border-color: var(--mcp-accent-cyan);
}

.sort-select :deep(.el-input__inner) {
  color: var(--mcp-text-primary) !important;
}

.sort-select :deep(.el-input__suffix) {
  color: var(--mcp-text-muted);
}

.content-layout {
  display: flex;
  gap: 24px;
  padding: 0 16px;
  min-height: 400px;
}

.tag-sidebar {
  flex: 0 0 20%;
  max-width: 240px;
  background: var(--mcp-dark-bg-card);
  border: 1px solid var(--mcp-border);
  border-radius: var(--radius-md);
  padding: 16px;
  height: fit-content;
  position: sticky;
  top: 100px;
}

.sidebar-header {
  margin-bottom: 16px;
}

.sidebar-header h3 {
  margin: 0 0 6px;
  font-size: 16px;
  color: var(--mcp-text-primary);
}

.sidebar-header p {
  margin: 0;
  color: var(--mcp-text-secondary);
  font-size: 13px;
}

.tag-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.tag-item {
  background: transparent;
  border: 1px solid var(--mcp-border);
  border-radius: var(--radius-sm);
  padding: 8px 12px;
  text-align: left;
  color: var(--mcp-text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 13px;
}

.tag-item:hover {
  border-color: var(--mcp-border-light);
  color: var(--mcp-text-primary);
}

.tag-item.active {
  border-color: var(--mcp-accent-cyan);
  color: var(--mcp-accent-cyan);
  background: rgba(0, 212, 255, 0.08);
}

.market-content {
  flex: 1;
}

.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.mcp-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 20px;
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

@media (max-width: 992px) {
  .content-layout {
    flex-direction: column;
  }

  .tag-sidebar {
    position: static;
    width: 100%;
    max-width: none;
  }
}

@media (max-width: 1024px) {
  .mcp-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .market-header {
    flex-direction: column;
  }

  .header-controls {
    width: 100%;
    justify-content: flex-start;
  }

  .search-input {
    flex: 1;
    width: 100%;
  }

  .sort-select {
    width: 100%;
  }

  .page-title {
    font-size: 28px;
  }

  .mcp-grid {
    grid-template-columns: 1fr;
  }

  .tag-sidebar {
    display: none;
  }
}
</style>
