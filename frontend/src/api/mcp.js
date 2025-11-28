import request from './request'

// 获取 MCP 服务列表（分页）
export const getMcpList = (params) => {
  const tags = Array.isArray(params.tags) ? params.tags.join(',') : params.tags
  return request({
    url: '/mcps',
    method: 'get',
    params: {
      page: params.page || 1,
      page_size: params.pageSize || 20,
      by: params.sortBy || 'id',
      asc: params.asc !== false ? 'true' : 'false',
      search: params.search || '',
      tags: tags || ''
    }
  })
}

// 获取单个 MCP 服务详情
export const getMcpDetail = (id) => {
  return request({
    url: `/mcps/${id}`,
    method: 'get'
  })
}

// 收藏 MCP 服务
export const addMcpFavorite = (id) => {
  return request({
    url: `/mcps/${id}/favorite`,
    method: 'post'
  })
}

// 取消收藏 MCP 服务
export const removeMcpFavorite = (id) => {
  return request({
    url: `/mcps/${id}/favorite`,
    method: 'delete'
  })
}

// 获取 MCP 标签列表
export const getMcpTags = (params) => {
  return request({
    url: '/mcp-tags',
    method: 'get',
    params: {
      page: params?.page || 1,
      page_size: params?.pageSize || 100
    }
  })
}

// 注册新 MCP 服务（需要认证）
export const registerMcp = (data) => {
  return request({
    url: '/mcps',
    method: 'post',
    data
  })
}

// 删除 MCP 服务（管理员）
export const deleteMcp = (id) => {
  return request({
    url: `/mcps/${id}`,
    method: 'delete'
  })
}
