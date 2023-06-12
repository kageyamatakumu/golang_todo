import React from 'react'
import { Axioslogout } from '../apis/logout'
import { useNavigate } from 'react-router-dom'

export const Todo = () => {
  const navigate = useNavigate()

  const submitLogoutHandler = async (e) => {
    e.preventDefault()
    Axioslogout().then(() => navigate('/'))
  }

  const submitToRenamePage = () => {
    navigate('/renamePage')
  }
  return(
    <>
      <div>Todoリスト</div>
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
