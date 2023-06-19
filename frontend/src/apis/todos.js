import axios from 'axios'

export const Axiostodos = async () => {
  axios.defaults.withCredentials = true;
  try {
    const response = await axios('http://localhost:8080/todos');
    return response.data;
  } catch (error) {
    throw error;
  }
};
