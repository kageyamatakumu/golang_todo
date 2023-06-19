import { Routes, Route } from 'react-router-dom'
import { useEffect } from 'react';
import './App.css';
import { Auth } from './containers/Auth'
import { Todo } from './containers/Todo';
import { Rename } from './containers/Rename';
import { Nomatch } from './containers/Nomatch';
import axios from 'axios';

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get('http://localhost:8080/csrf')
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])
  return (
    <>
    <Routes>
      <Route path="/" element={<Auth />}/>
      <Route path="/todosPage" element={<Todo />}/>
      <Route path="/renamePage" element={<Rename />}/>
      <Route path="*" element={<Nomatch/>} />
    </Routes>
    </>

  );
}

export default App;
