import React from 'react'
import { AxiosRename } from '../apis/user/rename'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

export const Rename = () => {
  const [name, setName] = useState('')
  const navigate = useNavigate()

  const submitRenameHandler = async (e) => {
    e.preventDefault()
    AxiosRename(name)
      .then(() => navigate('/todosPage'))
  }

  return (
    <>
      <form onSubmit={submitRenameHandler}>
        <div>
          <input
            name="name"
            type="text"
            autoFocus
            placeholder='rename'
            onChange={(e) => setName(e.target.value)}
            value={name}
          />
        </div>
        <div>
          <button
            type="submit"
          >ReName</button>
        </div>
      </form>
    </>
  )
}
