<template>
  <div>
    <form method="post">
      <div id="calculator">
        <div class="data-input">
          <label for="base">Base salary</label>
          <input
            type="text"
            name="base-salary"
            id="base"
            v-model="baseSal"
            pattern="[0-9]+.[0-9]{2}"
            placeholder="5000.00"
          />
        </div>
        <div class="data-input">
          <label for="dependents">Dependents</label>
          <input
            type="text"
            name="dependents"
            id="dependents"
            v-model="dependents"
            pattern="[0-9]+"
            placeholder="2"
          />
        </div>
        <div class="data-input">
          <label for="others">Other discounts</label>
          <input
            type="text"
            name="other-discounts"
            id="others"
            v-model="others"
            pattern="[0-9]+.[0-9]{2}"
            placeholder="45.22"
          />
        </div>
        <button type="submit" @click.prevent="calculate">Calculate</button>
      </div>
    </form>
    <!-- <hr /> -->
    <div id="amt-section">
      <h2 id="total-amt">Total amount</h2>
      <transition name="fade-slide-in" mode="out-in">
        <total-amount :value="result" v-if="result != null" :key="Date.now()" />
      </transition>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import Card from './Card.vue';
import TotalAmount from './TotalAmount.vue';
export default defineComponent({
  setup() {
    const baseSal = ref(0);
    const dependents = ref(0);
    const others = ref(0);

    const calculate = async () => {
      const body = JSON.stringify({
        salary: +baseSal.value,
        dependents: +dependents.value,
        others: +others.value,
      });
      const response = await fetch('http://localhost:8080/salary', {
        method: 'POST',
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
        },
        body,
      });
      const newSal = await response.json();
      console.log(newSal);
      result.value = newSal;
    };
    const result = ref<number | null>(null);

    return {
      baseSal,
      dependents,
      others,
      result,
      calculate,
    };
  },
  components: { Card, TotalAmount },
});
</script>

<style scoped>
label {
  font-weight: 600;
  font-size: 1.1rem;
  padding: 0.6475rem 0.4rem;
}
#calculator {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-start;
  height: 440px;
  padding: 18px 18px 2px 0px;
}

.data-input > label {
  display: block;
  text-align: left;
  font-size: 1.25rem;
  margin-bottom: 0.25rem;
}

.data-input > input[type='text'] {
  padding: 0.5rem 0.9rem;
  margin-left: 0.9ch;
  width: 300px;
  height: 1.75rem;
  border: 0;
  border-bottom: 4px solid #055c9ddd;
  border-radius: 0.175rem 0.175rem 0 0;
  background-color: #e7e7e7;
  /* color: #e8eef1; */
  font-weight: 900;
}

#calculator > button[type='submit'] {
  margin-top: 3rem;
  margin-left: 0.8ch;
  padding: 1rem;
  border: 0;
  border-radius: 0.175rem;
  background-color: #057dcd;
  font-weight: 900;
  font-size: 1rem;
  color: #f0f0f0;
  letter-spacing: 0.425ch;
  text-transform: uppercase;
  font-family: 'Inter', Helvetica, sans-serif;
  transition: all 0.3s cubic-bezier(0.165, 0.84, 0.44, 1);
}

#calculator > button[type='submit']:hover {
  background-color: #057dcddd;
  color: #e7e7e7;
}

#amt-section {
  text-align: left;
  margin: 0;
  margin-top: 3.75rem;
  margin-bottom: 2rem;
  margin-left: 6px;
}

hr {
  max-width: 560px;

  border-color: #055c9d22;
}

#total-amt {
  font-size: 1.35rem;
  font-weight: bold;
  letter-spacing: 0.0375ch;
}

.fade-slide-in-enter-active {
  transition: all 0.25s cubic-bezier(0.075, 0.82, 0.165, 1);
}

.fade-slide-in-leave-active {
  transition: all 0.3s cubic-bezier(0.075, 0.82, 0.165, 1);
}

.fade-slide-in-enter-from,
.fade-slide-in-leave-to {
  transform: translateX(-20px);
  opacity: 0;
}
</style>
