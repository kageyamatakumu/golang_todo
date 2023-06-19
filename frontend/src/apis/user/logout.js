import axios from "axios"

export const Axioslogout = async () => {
  try{
    await axios.post("http://localhost:8080/logout")
  }catch(error){
    if (error.response && error.response.data.message) {
      console.log(error.response.data.message);
    } else {
      console.log(error);
    }
  }
}
