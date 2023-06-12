import React from 'react'
import { AxiosAuth } from '../apis/auth';
import { AxiosSingUP } from '../apis/singup';
import { useState } from 'react'
import { useNavigate } from 'react-router-dom';

export const Auth = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [name, setName] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const navigate = useNavigate();

  const submitAuthHandler = async (e) => {
    e.preventDefault()
    AxiosAuth(email, password)
      .then(() => navigate('/todosPage'))
  }

  const submitSingUpHandler = async (e) => {
    e.preventDefault()
    AxiosSingUP(email, name, password)
      .then(() => AxiosAuth(email, password))
        .then(() => navigate('/todosPage'))
  }

  return (
    <>
      <div>{isLogin ? 'Auth' : 'SignUp'}</div>
      {isLogin ?
        <form onSubmit={submitAuthHandler}>
          <div>
            <input
              name="email"
              type="email"
              autoFocus
              placeholder='Email address'
              onChange={(e) => setEmail(e.target.value)}
              value={email}
            />
          </div>
          <div>
            <input
              name="password"
              type="password"
              placeholder='Password'
              onChange={(e) => setPassword(e.target.value)}
              value={password}
            />
          </div>
          <div>
            <button
              disabled={!email || !password}
              type="submit"
            >LOGIN</button>
          </div>
        </form>
      :
        <form onSubmit={submitSingUpHandler}>
          <div>
            <input
              name="email"
              type="email"
              autoFocus
              placeholder='Email address'
              onChange={(e) => setEmail(e.target.value)}
              value={email}
            />
          </div>
          <div>
            <input
              name="password"
              type="password"
              placeholder='Password'
              onChange={(e) => setPassword(e.target.value)}
              value={password}
            />
          </div>
          <div>
            <input
              name="name"
              type="text"
              placeholder='Name'
              onChange={(e) => setName(e.target.value)}
              value={name}
            />
          </div>
          <div>
            <button
              disabled={!email || !password || !name}
              type="submit"
            >SignUp</button>
          </div>
        </form>
      }
      <div>
        <button onClick={() => setIsLogin(!isLogin)}>
          {isLogin ? 'ToSingUp' : 'ToLogin'}
        </button>
      </div>
    </>
  )
}
