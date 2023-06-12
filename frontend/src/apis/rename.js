import axios from 'axios'

export const AxiosRename = async (name) => {
  const data = {name: name}

  try{
    await axios.put('http://localhost:8080/user/rename', data);
  }catch(error){
    if (error.response && error.response.data.message) {
      console.log(error.response.data.message);
    } else {
      console.log(error);
    }
  }
}
