import axios from "axios";

const API_URL = "http://localhost:8080";

export const uploadPDF = async (file: File): Promise<{ extracted_text: string; chatgpt_response: string }> => {
    const formData = new FormData();
    formData.append("file", file);

    const response = await axios.post<{ extracted_text: string; chatgpt_response: string }>(
        `${API_URL}/upload`,
        formData,
        { headers: { "Content-Type": "multipart/form-data" } }
    );

    return response.data;
};
