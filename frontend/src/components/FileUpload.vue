<template>
  <div class="bg-white p-6 rounded-lg shadow-lg w-full max-w-lg">
    <h2 class="text-2xl font-semibold mb-4">Upload a PDF</h2>

    <input
        type="file"
        @change="handleFileUpload"
        accept=".pdf"
        class="mb-4 block w-full border border-gray-300 rounded-lg p-2 focus:outline-none"
    />

    <button
        @click="uploadFile"
        :disabled="!file"
        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-lg w-full disabled:bg-gray-400">
      Upload
    </button>

    <p v-if="error" class="text-red-500 mt-2">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits } from "vue";
import { uploadPDF } from "../services/api";

const file = ref<File | null>(null);
const error = ref<string>("");
const emit = defineEmits<{
  (e: "upload-success", data: { extracted_text: string; chatgpt_response: string }): void;
}>();

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  file.value = target.files?.[0] ?? null;
};

const uploadFile = async () => {
  if (!file.value) return;

  try {
    error.value = "";
    const response = await uploadPDF(file.value);
    emit("upload-success", response);
  } catch (err) {
    error.value = "Error uploading file. Please try again.";
  }
};
</script>
