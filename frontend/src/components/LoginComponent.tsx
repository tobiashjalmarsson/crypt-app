import React, {FC} from 'react'
import './styles.scss'


const LoginComponent: FC = () => {
    return(
        <div>
            <form className="mainForm">
                <input type="text"/>
                <input type="text"/>
                <button type="submit">Log In</button>
            </form>
        </div>
    )
}

export default LoginComponent
