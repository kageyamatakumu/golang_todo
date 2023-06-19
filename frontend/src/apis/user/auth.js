import axios from 'axios'

export const AxiosAuth = async (email, password) => {
  const data = { email: email, password: password }

  try {
    await axios.post('http://localhost:8080/login', data);
  } catch (error) {
    if (error.response && error.response.data.message) {
      console.log(error.response.data.message);
    } else {
      console.log(error);
    }
  }
}