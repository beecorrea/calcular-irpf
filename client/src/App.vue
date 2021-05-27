<template>
  <div>
    <Header></Header>
    <Calculator />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Calculator from './components/Calculator.vue';
import Header from './components/Header.vue';
export default defineComponent({
  name: 'App',
  components: {
    Calculator,
    Header,
  },
});
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;900&display=swap');
#app {
  font-family: 'Inter', Helvetica, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  margin-left: 250px;
  margin-right: 100px;
}
</style>
