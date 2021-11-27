import React, {FC} from 'react'
import './styles.scss'

interface LoginProps {
    login: boolean,
    setLogin: (value: boolean) => void
}

const RegisterComponent: FC<LoginProps> = ({login, setLogin}) => {

    const handleClick = (e: any) => {
        e.preventDefault()
        setLogin(true)
    }

    return(
        <div>
            <form>
                <div>
                    <h2>Register</h2>
                    <p>Already have an account?<button id="changeLogin" onClick={handleClick}>Login</button></p>
                </div>
                <input type="text" placeholder="username" />
                <input type="text" placeholder="email" />
                <input type="text" placeholder="password" />
                <input type="text" placeholder="repeat password" />
                <button type="submit">Register</button>
            </form>
        </div>
    )
}

export default RegisterComponent
