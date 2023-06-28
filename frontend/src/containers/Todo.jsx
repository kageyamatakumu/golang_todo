import React, { useEffect, useState } from 'react'
import { Axioslogout } from '../apis/user/logout'
import { useNavigate } from 'react-router-dom'
import { Axiostodos } from '../apis/todos'

export const Todo = () => {
  const navigate = useNavigate()
  const [todos, setTodos] = useState([]);

  const submitLogoutHandler = async (e) => {
    e.preventDefault()
    Axioslogout().then(() => navigate('/'))
  }

  const submitToRenamePage = () => {
    navigate('/renamePage')
  }

  useEffect(() => {
    const data = async () => {
      try {
        const result = await Axiostodos();
        setTodos(result)
      } catch (error) {
        console.log(error);
      }
    };
    data();
  },[])

  return(
    <>
      <div>Todoリスト</div>
      <div>
        {todos.map((todo, index) => {
          return (
            <div key={todo.id}>
              <div>
                {todo.title}
              </div>
              <div>
                {todo.description}
              </div>
            </div>
          )
        })}
      </div>
      <div>
        <button onClick={submitLogoutHandler}>
          ログアウト
        </button>
      </div>
      <div>
        <button onClick={submitToRenamePage}>
          名前を変えたい
        </button>
      </div>
    </>
  )
}
