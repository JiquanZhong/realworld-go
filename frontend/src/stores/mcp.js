import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getMcpList, getMcpTags } from '@/api/mcp'

export const useMcpStore = defineStore('mcp', () => {
  // 状态
  const mcpList = ref([])
  const total = ref(0)
  const loading = ref(false)
  const tags = ref([])

  // 当前筛选条件
  const currentPage = ref(1)
  const pageSize = ref(20)
  const sortBy = ref('install_count')
  const sortAsc = ref(false)
  const selectedCategory = ref('')
  const selectedTags = ref([])
  const searchKeyword = ref('')

  // 获取 MCP 列表
  const fetchMcpList = async () => {
    loading.value = true
    try {
      const data = await getMcpList({
        page: currentPage.value,
        pageSize: pageSize.value,
        sortBy: sortBy.value,
        asc: sortAsc.value,
        search: searchKeyword.value
      })
      mcpList.value = data.list || []
      total.value = data.total || 0
    } catch (error) {
      console.error('获取 MCP 列表失败:', error)
      mcpList.value = []
      total.value = 0
    } finally {
      loading.value = false
    }
  }

  // 获取标签列表
  const fetchTags = async () => {
    try {
      const data = await getMcpTags({ pageSize: 100 })
      tags.value = data.list || []
    } catch (error) {
      console.error('获取标签列表失败:', error)
      tags.value = []
    }
  }

  // 改变页码
  const changePage = (page) => {
    currentPage.value = page
    fetchMcpList()
  }

  // 改变排序
  const changeSort = (field, asc = false) => {
    sortBy.value = field
    sortAsc.value = asc
    currentPage.value = 1
    fetchMcpList()
  }

  // 搜索
  const search = (keyword) => {
    searchKeyword.value = keyword
    currentPage.value = 1
    fetchMcpList()
  }

  // 重置筛选条件
  const resetFilters = () => {
    currentPage.value = 1
    sortBy.value = 'install_count'
    sortAsc.value = false
    selectedCategory.value = ''
    selectedTags.value = []
    searchKeyword.value = ''
    fetchMcpList()
  }

  return {
    // 状态
    mcpList,
    total,
    loading,
    tags,
    currentPage,
    pageSize,
    sortBy,
    sortAsc,
    selectedCategory,
    selectedTags,
    searchKeyword,

    // 方法
    fetchMcpList,
    fetchTags,
    changePage,
    changeSort,
    search,
    resetFilters
  }
})
