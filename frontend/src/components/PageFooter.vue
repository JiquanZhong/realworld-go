<template>
  <footer class="page-footer">
    <div class="footer-content">
      <nav class="footer-links">
        <a href="#" class="footer-link">关于</a>
        <a href="#" class="footer-link">联系我们</a>
        <a href="#" class="footer-link">隐私政策</a>
        <a href="#" class="footer-link">服务条款</a>
      </nav>
      <p class="footer-copyright">©2025 MCP hub 版权所有.</p>
    </div>
  </footer>
</template>

<script setup>
</script>

<style scoped>
.page-footer {
  background: var(--mcp-dark-bg-elevated);
  border-top: 1px solid var(--mcp-border);
  padding: 32px 40px;
  margin-top: auto;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.footer-links {
  display: flex;
  gap: 32px;
  flex-wrap: wrap;
  justify-content: center;
}

.footer-link {
  color: var(--mcp-text-secondary);
  text-decoration: none;
  font-size: 14px;
  transition: color 0.3s;
}

.footer-link:hover {
  color: var(--mcp-accent-cyan);
}

.footer-copyright {
  color: var(--mcp-text-muted);
  font-size: 13px;
  margin: 0;
}

@media (max-width: 768px) {
  .page-footer {
    padding: 24px 16px;
  }

  .footer-links {
    gap: 16px;
  }
}
</style>
