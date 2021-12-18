import React, {FC, useState} from 'react'
import {APICallLogin} from './api'
import './styles.scss'
import '../../index.scss'

interface LoginProps {
    login: boolean,
    setLogin: (value: boolean) => void
}

const LoginComponent: FC<LoginProps> = ({login, setLogin}) => {

    const [user, setUser] = useState({email: "", password: ""}) 
    const [error, setError] = useState(false)

    const handleClick = (e: any) => {
        e.preventDefault()
        setLogin(false)
    }

    const handleSubmit = async (e: any) => {
        e.preventDefault()
        const result = await APICallLogin(user.email, user.password)
        setError(!result)
    }

    const handleChange = (e: any) : void => {
        e.preventDefault()
        setUser({
            ...user,
            [e.target.name] : e.target.value,
        })
    }
    return(
            <form className="loginForm" onSubmit={handleSubmit}>
                <div>
                    <h2>Login</h2>
                    <p id="navigateP" onClick={handleClick}>Don't Have an account? Register!</p>
                </div>
                <input name="email" type="text" placeholder="Email" onChange={handleChange} value={user.email}/>
                <input name="password" type="text" placeholder="Password" onChange={handleChange} value={user.password}/>
                <button className="mainButton" type="submit">Log In</button>
                {error && <p>Invalid Information</p>}
            </form>
    )
}

export default LoginComponent
