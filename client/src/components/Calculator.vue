<template>
  <div>
    <Card>
      <template v-slot:content>
        <form action="" method="post">
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
          </div>
          <button type="submit" @click.prevent="calculate">Calculate</button>
        </form>
      </template>
    </Card>
    <hr />
    <total-amount :value="result" id="amt-section" v-if="result != null" />
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
  justify-content: space-around;
  align-items: flex-start;
  height: 290px;
  width: 500px;
  padding: 18px 18px 2px 18px;
}

.data-input {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  box-sizing: border-box;
  text-align: left;
}

#calculator ~ button[type='submit'] {
  margin: 1rem 0;
  width: 500px;
  padding: 1.2rem;
  border: 0;
  border-radius: 0.175rem;
  background-color: #057dcd;
  font-weight: 900;
  font-size: 1rem;
  color: #f0f0f0;
  letter-spacing: 0.5ch;
  text-transform: uppercase;
  font-family: 'Inter', Helvetica, sans-serif;
}

.data-input > input[type='text'] {
  padding: 0.5rem 0.9rem;
  border: 0;
  border-bottom: 4px solid #055c9ddd;
  border-radius: 0.175rem 0.175rem 0 0;
  background-color: #055c9daa;
  /* color: #e8eef1; */
  font-weight: 900;
}

#amt-section {
  text-align: left;
  /* margin-top: 4rem; */
}

hr {
  max-width: 560px;
  margin: 0;
  margin-top: 4rem;
  margin-bottom: 3rem;
  border-color: #055c9d22;
}
</style>
