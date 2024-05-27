<template>
  <el-container class="captcha-box">
    <el-button type="primary" class="send-btn" :size="props.size" :disabled="!canSend" @click="sendMsg" plain>
      {{ btnText }}
    </el-button>
  </el-container>
</template>


<script setup>
// 导入必要模块
import { ref } from "vue";
import { validateEmail, validateMobile } from "@/utils/validate";
import { ElMessage } from "element-plus";
import { httpPost } from "@/utils/http";

// 定义props
const props = defineProps({
  receiver: String,
  size: String,
});

// 定义响应式变量
const btnText = ref('发送验证码')
const canSend = ref(true)

// 发送短信验证码
const sendMsg = () => {
  if (!validateMobile(props.receiver) && !validateEmail(props.receiver)) {
    return showMessageError("请输入合法的手机号/邮箱地址")
  }

  if (!canSend.value) {
    return
  }

  canSend.value = false
  httpPost('/api/sms/code', { receiver: props.receiver }).then(() => {
    ElMessage.success('验证码发送成功')
    let time = 120
    btnText.value = time
    const handler = setInterval(() => {
      time = time - 1
      if (time <= 0) {
        clearInterval(handler)
        btnText.value = '重新发送'
        canSend.value = true
      } else {
        btnText.value = time
      }
    }, 1000)
  }).catch(e => {
    canSend.value = true
    showMessageError('验证码发送失败：' + e.message)
  })
}
</script>

.captcha-box {
  .send-btn {
    width: 100%;
  }
}
