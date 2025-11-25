<template>
  <div class="register-container">
    <!-- 返回按钮 -->
    <el-button
      class="back-button"
      circle
      size="large"
      @click="goBack"
    >
      <el-icon :size="20"><ArrowLeft /></el-icon>
    </el-button>

    <el-card class="register-card">
      <template #header>
        <div class="card-header">
          <h2>用户注册</h2>
          <p class="subtitle">创建您的账户</p>
        </div>
      </template>

      <el-form
        ref="registerFormRef"
        :model="registerForm"
        :rules="rules"
        label-width="0"
        size="large"
        @submit.prevent="handleRegister"
      >
        <el-form-item prop="name">
          <el-input
            v-model="registerForm.name"
            placeholder="请输入用户名"
            prefix-icon="User"
            clearable
          />
        </el-form-item>

        <el-form-item prop="email">
          <el-input
            v-model="registerForm.email"
            placeholder="请输入邮箱"
            prefix-icon="User"
            clearable
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="registerForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            show-password
          />
        </el-form-item>

        <el-form-item prop="confirmPassword">
          <el-input
            v-model="registerForm.confirmPassword"
            type="password"
            placeholder="请再次输入密码"
            prefix-icon="Lock"
            show-password
          />
        </el-form-item>

        <el-form-item prop="age">
          <el-input
            v-model.number="registerForm.age"
            placeholder="请输入年龄"
            prefix-icon="User"
            clearable
            @keyup.enter="handleRegister"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            class="register-button"
            @click="handleRegister"
          >
            {{ loading ? '注册中...' : '注 册' }}
          </el-button>
        </el-form-item>
      </el-form>

      <div class="footer-links">
        <span class="tip">已有账户？</span>
        <el-link type="primary" @click="goToLogin">立即登录</el-link>
      </div>

      <div class="footer-links" style="border-top: none; padding-top: 10px;">
        <el-link type="info" @click="goBack">返回 MCP 广场</el-link>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeft, User, Message, Lock } from '@element-plus/icons-vue'
import { register } from '@/api/auth'
import { useFormValidation } from '@/composables/useFormValidation'

const router = useRouter()
const registerFormRef = ref(null)
const loading = ref(false)

// 表单数据
const registerForm = reactive({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
  age: null
})

// 使用共享的表单验证规则
const { emailRules, passwordRules, nameRules, ageRules, confirmPasswordRules } = useFormValidation()

// 表单验证规则
const rules = computed(() => ({
  name: nameRules,
  email: emailRules,
  password: passwordRules,
  confirmPassword: confirmPasswordRules(registerForm.password),
  age: ageRules
}))

// 返回到 MCP 广场
const goBack = () => {
  router.push('/mcp-market')
}

// 跳转到登录页
const goToLogin = () => {
  router.push('/login')
}

// 处理注册
const handleRegister = async () => {
  if (!registerFormRef.value) return

  try {
    // 验证表单
    await registerFormRef.value.validate()

    loading.value = true

    // 准备注册数据（不包含 confirmPassword）
    const registerData = {
      name: registerForm.name,
      email: registerForm.email,
      password: registerForm.password,
      age: registerForm.age
    }

    // 调用注册 API
    // 注意：request.js 的响应拦截器已经处理了 response.data.data
    // 所以这里直接得到的是用户对象，而不是完整的 {code, data, message}
    const userData = await register(registerData)

    // 如果成功返回了用户数据
    if (userData) {
      ElMessage.success('注册成功！请登录您的账户')

      // 延迟跳转到登录页
      setTimeout(() => {
        router.push('/login')
      }, 1500)
    }
  } catch (error) {
    console.error('注册失败:', error)
    // 注意：
    // 1. 如果是表单验证错误（error === false），ElementPlus 会自动显示
    // 2. 如果是 API 错误，request.js 的响应拦截器已经显示了错误消息
    // 所以这里不需要再显示错误消息
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.register-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  /* background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); */
  padding: 20px;
}

.register-card {
  width: 100%;
  max-width: 420px;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
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

.register-button {
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
  .register-card {
    max-width: 100%;
  }

  .register-header h2 {
    font-size: 24px;
  }
}
</style>
