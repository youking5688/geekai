<template>
  <div>正在处理微信登录...</div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { httpGet } from '@/utils/http'
import { setUserToken } from '@/store/session'
const router = useRouter()

// 获取 URL 参数
const query = new URLSearchParams(window.location.search)
const code = query.get('code')
const state = query.get('state')

// 发送请求到后端以处理微信登录
httpGet('/api/user/weChatCallback', { code, state })
  .then((res) => {
    setUserToken(res.data)
    if (isMobile()) {
      router.push('/mobile')
    } else {
      router.push('/chat')
    }
    router.push('/')
  })
  .catch((e) => {
    console.error('Error:', e)
    router.push('/')
    // alert('微信登录失败，请重试')
  })
</script>
