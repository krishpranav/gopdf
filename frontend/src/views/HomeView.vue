<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-50 p-4">
    <h1 class="text-3xl font-bold text-gray-800 mb-6">PDF Text Extractor</h1>

    <FileUpload @upload-success="handleUploadSuccess" />

    <ExtractedText v-if="extractedText" :text="extractedText" :chatResponse="chatGPTResponse" />
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import FileUpload from "../components/FileUpload.vue";
import ExtractedText from "../components/ExtractedText.vue";

const extractedText = ref<string>("");
const chatGPTResponse = ref<string>("");

const handleUploadSuccess = (data: { extracted_text: string; chatgpt_response: string }) => {
  extractedText.value = data.extracted_text;
  chatGPTResponse.value = data.chatgpt_response;
};
</script>
