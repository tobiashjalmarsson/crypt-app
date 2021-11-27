import React, {FC, useState} from 'react'
import LoginComponent from '../components/LoginComponent/LoginComponent'
import RegisterComponent from '../components/LoginComponent/RegisterComponent'


const LoginRegistration: FC = () => {
    // Determines if we display the login or register component
    const [login, setLogin] = useState(true)

    return(
        <>
            {login ?
            <LoginComponent login={login} setLogin={setLogin} />
            :
            <RegisterComponent login={login} setLogin={setLogin}/>
            }
        </>
    )
}

export default LoginRegistration
