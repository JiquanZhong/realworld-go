/**
 * 表单验证规则组合函数
 * Shared form validation rules for authentication forms
 */

export function useFormValidation() {
  // 邮箱验证规则
  const emailRules = [
    {
      required: true,
      message: '请输入邮箱地址',
      trigger: 'blur'
    },
    {
      type: 'email',
      message: '请输入正确的邮箱地址',
      trigger: ['blur', 'change']
    }
  ]

  // 密码验证规则
  const passwordRules = [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    },
    {
      min: 6,
      message: '密码长度不能小于6位',
      trigger: ['blur', 'change']
    }
  ]

  // 用户名验证规则
  const nameRules = [
    {
      required: true,
      message: '请输入用户名',
      trigger: 'blur'
    },
    {
      min: 2,
      max: 50,
      message: '用户名长度应在2-50个字符之间',
      trigger: ['blur', 'change']
    }
  ]

  // 年龄验证规则
  const ageRules = [
    {
      required: true,
      message: '请输入年龄',
      trigger: 'blur'
    },
    {
      type: 'number',
      message: '年龄必须是数字',
      trigger: ['blur', 'change']
    },
    {
      validator: (rule, value, callback) => {
        if (value < 0 || value > 130) {
          callback(new Error('年龄必须在0-130之间'))
        } else {
          callback()
        }
      },
      trigger: ['blur', 'change']
    }
  ]

  // 确认密码验证规则（需要传入原密码值）
  const confirmPasswordRules = (password) => [
    {
      required: true,
      message: '请再次输入密码',
      trigger: 'blur'
    },
    {
      validator: (rule, value, callback) => {
        if (value !== password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: ['blur', 'change']
    }
  ]

  return {
    emailRules,
    passwordRules,
    nameRules,
    ageRules,
    confirmPasswordRules
  }
}
