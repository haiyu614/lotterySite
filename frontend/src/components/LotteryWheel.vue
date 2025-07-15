<template>
  <div>
    <div id="lottery-wheel">
      <div v-for="(good, index) in goods" :key="good.id" class="prize" :style="{ transform: `rotate(${index * angle}deg)` }">
        {{ good.name }}
      </div>
    </div>
    <button id="start-button" @click="startLottery">开始抽奖</button>
    <div id="result">{{ result }}</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
const API_URL = 'http://localhost:8084';
const goods = ref([]);
const result = ref('');
const angle = ref(0);

// 获取所有商品信息
const getGoods = async () => {
  try {
    const response = await fetch(`${API_URL}/goods`);
    if (!response.ok) {
      throw new Error('获取商品信息失败');
    }
    goods.value = await response.json();
    angle.value = 360 / goods.value.length;
  } catch (error) {
    console.error(error);
  }
};

// 开始抽奖
const startLottery = async () => {
  try {
    const response = await fetch(`${API_URL}/lottery`);
    if (!response.ok) {
      throw new Error('抽奖失败');
    }
    const data = await response.json();
    const luckyId = data.lucky_id;
    const luckyGood = goods.value.find(good => good.id === luckyId);
    if (luckyGood) {
      result.value = `恭喜你，抽中了 ${luckyGood.name}!`;
    } else {
      result.value = '未找到中奖商品信息';
    }
  } catch (error) {
    console.error(error);
    result.value = '抽奖失败，请稍后重试';
  }
};

onMounted(() => {
  getGoods();
});
</script>

<style scoped>
/* 这里可以添加组件级别的样式 */
</style>