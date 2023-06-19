import axios from "axios";

export const AxiosSingUP = async (email, name, password) => {
  const data = {email: email, name: name, password: password}

  try{
    await axios.post('http://localhost:8080/signup', data);
  } catch(error) {
    if (error.response && error.response.data.message) {
      console.log(error.response.data.message);
    } else {
      console.log(error);
    }
  }
}