<template>
  <div v-if="text" class="bg-gray-100 p-6 rounded-lg shadow-lg w-full max-w-2xl">
    <h2 class="text-xl font-semibold mb-4">Extracted Text</h2>
    <pre class="whitespace-pre-wrap text-gray-700">{{ text }}</pre>

    <h2 class="text-xl font-semibold mt-6 mb-4">ChatGPT Response</h2>
    <pre class="whitespace-pre-wrap text-green-700">{{ chatResponse }}</pre>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  text: string;
  chatResponse: string;
}>();
</script>
