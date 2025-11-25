<template>
  <div class="login-container">
    <!-- 返回按钮 -->
    <el-button
      class="back-button"
      circle
      size="large"
      @click="goBack"
    >
      <el-icon :size="20"><ArrowLeft /></el-icon>
    </el-button>

    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <h2>用户登录</h2>
          <p class="subtitle">欢迎回到 MCP 服务广场</p>
        </div>
      </template>

      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        label-width="0"
        size="large"
        @submit.prevent="handleLogin"
      >
        <el-form-item prop="email">
          <el-input
            v-model="loginForm.email"
            placeholder="请输入邮箱"
            prefix-icon="User"
            clearable
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            show-password
            @keyup.enter="handleLogin"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            class="login-button"
            @click="handleLogin"
          >
            {{ loading ? '登录中...' : '登 录' }}
          </el-button>
        </el-form-item>
      </el-form>

      <div class="footer-links">
        <span class="tip">还没有账号？</span>
        <el-link type="primary" @click="goToRegister">立即注册</el-link>
      </div>

      <div class="footer-links" style="border-top: none; padding-top: 10px;">
        <el-link type="info" @click="goBack">返回 MCP 广场</el-link>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useFormValidation } from '@/composables/useFormValidation'

const router = useRouter()
const userStore = useUserStore()

// 表单引用
const loginFormRef = ref(null)

// 加载状态
const loading = ref(false)

// 表单数据
const loginForm = reactive({
  email: '',
  password: ''
})

// 使用共享的表单验证规则
const { emailRules, passwordRules } = useFormValidation()

// 表单验证规则
const loginRules = {
  email: emailRules,
  password: passwordRules
}

// 处理登录
const handleLogin = async () => {
  if (!loginFormRef.value) return

  try {
    // 表单验证
    await loginFormRef.value.validate()

    loading.value = true

    // 调用登录接口
    const success = await userStore.login(loginForm.email, loginForm.password)

    if (success) {
      // 登录成功，跳转到首页
      router.push('/')
    }
  } catch (error) {
    // 表单验证失败
    if (error !== false) {
      console.error('Login validation error:', error)
    }
  } finally {
    loading.value = false
  }
}

// 跳转到注册页面
const goToRegister = () => {
  router.push('/register')
}

// 返回到 MCP 广场
const goBack = () => {
  router.push('/mcp-market')
}
</script>

<style scoped>
.login-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.back-button {
  position: absolute;
  top: 10vh;
  left: 30px;
  background: rgba(255, 255, 255, 0.9);
  border: none;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.15);
  transition: all 0.3s;
}

.back-button:hover {
  background: #fff;
  transform: translateX(-3px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}

.login-card {
  width: 100%;
  max-width: 420px;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.card-header {
  text-align: center;
}

.card-header h2 {
  margin: 0 0 8px 0;
  font-size: 28px;
  font-weight: 600;
  color: #303133;
}

.subtitle {
  margin: 0;
  font-size: 14px;
  color: #909399;
}

.el-form {
  margin-top: 20px;
}

.login-button {
  width: 100%;
  margin-top: 10px;
  font-size: 16px;
  height: 44px;
  border-radius: 6px;
}

.footer-links {
  text-align: center;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #ebeef5;
}

.tip {
  color: #909399;
  font-size: 14px;
  margin-right: 8px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-card {
    max-width: 100%;
  }

  .card-header h2 {
    font-size: 24px;
  }
}
</style>
