import React, {FC, useState} from 'react'
import './styles.scss'
import {APICallRegister} from './api'

interface LoginProps {
    login: boolean,
    setLogin: (value: boolean) => void
}

const RegError = {
    MatchPassword : "Not matching passwords",
    EnterPassword : "Fill in password",
    Username: "Username already exists",
    EnterUsername: "Please enter a username",
    None: "None"
}

const RegisterComponent: FC<LoginProps> = ({login, setLogin}) => {
    const [error, setError] = useState(RegError.None)
    const [user, setUser] = useState({
        username: "",
        password: "",
        passwordRepeat: ""
    })

    const handleClick = (e: any) => {
        e.preventDefault()
        setLogin(true)
    }

    const handleSubmit = async (e: any) => {
        e.preventDefault()
        
        if (user.password === "") {
            setError(RegError.EnterPassword)     
        } else if (user.username === "") {
            setError(RegError.EnterUsername)
        } else if (user.password !== user.passwordRepeat) {
            setError(RegError.MatchPassword)
        } else {
            setError(RegError.None)
            const success = await APICallRegister(user.username, user.password)
            if (!success){
                setError(RegError.Username)
            }
        }
    }

    const handleChange = (e: any) => {
        e.preventDefault()
        setUser({
            ...user,
            [e.target.name] : e.target.value,
        })
    }

    return(
        <div>
            <form onSubmit={handleSubmit}>
                <div>
                    <h2>Register</h2>
                    <p>Already have an account?<button id="changeLogin" onClick={handleClick}>Login</button></p>
                </div>
                <input type="text" placeholder="username" onChange={handleChange} value={user.username} name="username" />
                <input type="text" placeholder="password" onChange={handleChange} value={user.password} name="password" />
                <input type="text" placeholder="repeat password" onChange={handleChange} value={user.passwordRepeat} name="passwordRepeat" />
                <button type="submit">Register</button>
                {(error === RegError.MatchPassword) && <p>{RegError.MatchPassword}</p>}
                {(error === RegError.EnterUsername) && <p>{RegError.EnterUsername}</p>}
                {(error === RegError.EnterPassword) && <p>{RegError.EnterPassword}</p>}
                {(error === RegError.Username) && <p>{RegError.Username}</p>}
            </form>
        </div>
    )
}

export default RegisterComponent
