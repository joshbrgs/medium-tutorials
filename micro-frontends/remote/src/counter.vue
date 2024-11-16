<template>
    <button
      @click="increment"
      class="relative px-8 py-4 text-2xl font-bold text-black bg-white border-4 border-black shadow-[8px_8px_0_0_rgba(0,0,0,1)] hover:shadow-[12px_12px_0_0_rgba(0,0,0,1)] active:translate-x-2 active:translate-y-2 active:shadow-[4px_4px_0_0_rgba(0,0,0,1)] transition-transform duration-150"
    >
      Counter: {{ count }}
    </button>
</template>

<script>
import { ref } from "vue";

export default {
  setup() {
    const count = ref(0);

    const increment = () => {
      count.value++;
    };

    return {
      count,
      increment,
    };
  },
};
</script>

<style>
/* No additional styles needed; everything is styled with Tailwind CSS */
</style>

