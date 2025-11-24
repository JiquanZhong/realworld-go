<template>
  <div class="mcp-market">
    <div class="market-header">
      <h2>探索 MCP 服务</h2>
      <p class="subtitle">发现和使用强大的 Model Context Protocol 服务</p>
    </div>

    <div class="filter-section">
      <div class="filter-left">
        <el-select
          v-model="sortOption"
          placeholder="排序方式"
          style="width: 180px"
          @change="handleSortChange"
        >
          <el-option label="安装量最多" value="install_count_desc" />
          <el-option label="安装量最少" value="install_count_asc" />
          <el-option label="评分最高" value="rating_avg_desc" />
          <el-option label="评分最低" value="rating_avg_asc" />
          <el-option label="最新创建" value="created_at_desc" />
          <el-option label="最早创建" value="created_at_asc" />
        </el-select>

        <el-button @click="handleRefresh" :icon="Refresh">刷新</el-button>
      </div>

      <div class="filter-right">
        <span v-if="mcpStore.searchKeyword" class="search-info">
          搜索 "{{ mcpStore.searchKeyword }}"
        </span>
        <span class="result-count">共 {{ mcpStore.total }} 个服务</span>
      </div>
    </div>

    <el-divider />

    <div v-loading="mcpStore.loading" class="content-section">
      <div v-if="mcpStore.mcpList.length === 0 && !mcpStore.loading" class="empty-state">
        <el-empty description="暂无 MCP 服务">
          <el-button type="primary">提交服务</el-button>
        </el-empty>
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
        layout="prev, pager, next, jumper"
        @current-change="handlePageChange"
      />
    </div>

    <!-- MCP 详情对话框 -->
    <el-dialog
      v-model="detailDialogVisible"
      :title="selectedMcp?.name"
      width="800px"
      @close="handleDialogClose"
    >
      <div v-if="selectedMcp" class="mcp-detail">
        <div class="detail-header">
          <div class="icon-wrapper-large">
            <img
              v-if="selectedMcp.icon_url"
              :src="selectedMcp.icon_url"
              :alt="selectedMcp.name"
              class="mcp-icon-large"
            />
            <el-icon v-else :size="80" color="#909399"><Box /></el-icon>
          </div>
          <div class="detail-info">
            <h2>{{ selectedMcp.name }}</h2>
            <el-tag type="info">{{ selectedMcp.category || '未分类' }}</el-tag>
            <div class="detail-stats">
              <span><el-icon><Download /></el-icon> {{ selectedMcp.install_count || 0 }} 次安装</span>
              <span><el-icon><Star /></el-icon> {{ selectedMcp.rating_avg ? selectedMcp.rating_avg.toFixed(1) : '0.0' }} 评分</span>
            </div>
          </div>
        </div>

        <el-divider />

        <div class="detail-section">
          <h3>描述</h3>
          <p>{{ selectedMcp.description || '暂无描述' }}</p>
        </div>

        <div v-if="selectedMcp.endpoint" class="detail-section">
          <h3>服务端点</h3>
          <el-link :href="selectedMcp.endpoint" target="_blank" type="primary">
            {{ selectedMcp.endpoint }}
          </el-link>
        </div>

        <div v-if="selectedMcp.tags && selectedMcp.tags.length > 0" class="detail-section">
          <h3>标签</h3>
          <div class="tags-list">
            <el-tag
              v-for="tag in selectedMcp.tags"
              :key="tag.id"
              effect="plain"
              class="tag-item"
            >
              {{ tag.name }}
            </el-tag>
          </div>
        </div>

        <div v-if="selectedMcp.json_schema" class="detail-section">
          <h3>JSON Schema</h3>
          <el-input
            type="textarea"
            :value="selectedMcp.json_schema"
            :rows="8"
            readonly
          />
        </div>
      </div>

      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
        <el-button type="primary" @click="handleInstall">安装</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, shallowRef } from 'vue'
import { useMcpStore } from '@/stores/mcp'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import McpCard from '@/components/McpCard.vue'

const mcpStore = useMcpStore()
const sortOption = ref('install_count_desc')
const detailDialogVisible = ref(false)
const selectedMcp = shallowRef(null)

onMounted(() => {
  loadData()
})

const loadData = async () => {
  try {
    await mcpStore.fetchMcpList()
  } catch (error) {
    ElMessage.error('加载 MCP 列表失败')
  }
}

const handleSortChange = (value) => {
  // Split by last underscore to get field and order
  const lastUnderscoreIndex = value.lastIndexOf('_')
  const field = value.substring(0, lastUnderscoreIndex)
  const order = value.substring(lastUnderscoreIndex + 1)
  const isDesc = order === 'desc'
  mcpStore.changeSort(field, !isDesc)
}

const handleRefresh = () => {
  loadData()
}

const handlePageChange = (page) => {
  mcpStore.changePage(page)
}

const handleMcpClick = (mcp) => {
  selectedMcp.value = mcp
  detailDialogVisible.value = true
}

const handleDialogClose = () => {
  selectedMcp.value = null
}

const handleInstall = () => {
  ElMessage.info('安装功能开发中...')
}
</script>

<style scoped>
.mcp-market {
  max-width: 1400px;
  margin: 0 auto;
}

.market-header {
  margin-bottom: 32px;
}

.market-header h2 {
  margin: 0 0 8px 0;
  font-size: 32px;
  color: #303133;
}

.subtitle {
  margin: 0;
  font-size: 16px;
  color: #909399;
}

.filter-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.filter-left {
  display: flex;
  gap: 12px;
  align-items: center;
}

.filter-right {
  display: flex;
  gap: 12px;
  align-items: center;
}

.search-info {
  font-size: 14px;
  color: #606266;
  font-weight: 500;
}

.result-count {
  font-size: 14px;
  color: #909399;
}

.content-section {
  min-height: 400px;
}

.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.mcp-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 24px;
  margin-bottom: 32px;
}

.pagination-section {
  display: flex;
  justify-content: center;
  padding: 24px 0;
}

.mcp-detail {
  padding: 20px;
}

.detail-header {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

.icon-wrapper-large {
  flex-shrink: 0;
  width: 120px;
  height: 120px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f7fa;
  border-radius: 12px;
  overflow: hidden;
}

.mcp-icon-large {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.detail-info {
  flex: 1;
}

.detail-info h2 {
  margin: 0 0 12px 0;
  font-size: 24px;
  color: #303133;
}

.detail-stats {
  display: flex;
  gap: 20px;
  margin-top: 12px;
  font-size: 14px;
  color: #606266;
}

.detail-stats span {
  display: flex;
  align-items: center;
  gap: 4px;
}

.detail-section {
  margin-bottom: 24px;
}

.detail-section h3 {
  margin: 0 0 12px 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.detail-section p {
  margin: 0;
  font-size: 14px;
  color: #606266;
  line-height: 1.8;
}

.tags-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag-item {
  font-size: 13px;
}
</style>
