<template>
  <div class="lottery-container">
    <!-- 抽奖轮盘区域 -->
    <div>
      <div id="lottery-wheel" :style="{ transform: `rotate(${rotationAngle}deg)` }">
        <div v-for="(good, index) in goods" :key="good.id" class="prize" :style="{ transform: `rotate(${index * angle}deg)` }">
          {{ good.name }}
        </div>
      </div>
      <button id="start-button" @click="startLottery" :disabled="isRotating">开始抽奖</button>
      <div id="result">{{ result }}</div>
    </div>
    <!-- 商品列表区域 -->
    <div class="goods-list-container">
      <h2>商品列表</h2>
      <ul class="goods-list">
        <li v-for="(good, index) in availableGoods" :key="good.id">
          <span class="good-name">{{ good.name }}</span>
          <span class="good-stock">剩余库存: {{ good.number }}</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
const API_URL = 'http://localhost:8084';
const goods = ref([]);
const result = ref('');
const angle = ref(0);
const rotationAngle = ref(0);
const isRotating = ref(false);
const availableGoods = ref([]);

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

// 获取可用商品信息（库存大于0）
const getAvailableGoods = async () => {
  try {
    const response = await fetch(`${API_URL}/goodsforlottery`);
    if (!response.ok) {
      throw new Error('获取可用商品信息失败');
    }
    const availableGoodsData = await response.json();
    availableGoods.value = availableGoodsData.map(item => {
      const fullGood = goods.value.find(good => good.id === item.id);
      return {
        id: item.id,
        name: fullGood ? fullGood.name : '未知商品',
        number: item.number
      };
    });
  } catch (error) {
    console.error(error);
  }
};

// 开始抽奖
const startLottery = async () => {
  if (isRotating.value) return;
  isRotating.value = true;
  try {
    const response = await fetch(`${API_URL}/lottery`);
    if (!response.ok) {
      throw new Error('抽奖失败');
    }
    const data = await response.json();
    const luckyId = data.lucky_id;
    const luckyGood = goods.value.find(good => good.id === luckyId);
    if (luckyGood) {
      const index = goods.value.indexOf(luckyGood);
      const finalAngle = 3600 + index * angle.value; // 旋转多圈后停在中奖区域
      rotationAngle.value = finalAngle;
      setTimeout(() => {
        result.value = `恭喜你，抽中了 ${luckyGood.name}!`;
        isRotating.value = false;
        rotationAngle.value = 0;
        // 重新获取可用商品信息以更新库存显示
        getAvailableGoods();
      }, 5000); // 动画时长 5 秒
    } else {
      result.value = '未找到中奖商品信息';
      isRotating.value = false;
    }
  } catch (error) {
    console.error(error);
    result.value = '抽奖失败，请稍后重试';
    isRotating.value = false;
  }
};

onMounted(() => {
  getGoods();
  getAvailableGoods();
});
</script>

<style scoped>
/* 这里可以添加组件级别的样式 */
</style>