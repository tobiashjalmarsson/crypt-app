import React, {FC} from 'react'

const LoginComponent: FC = () => {
    return(
        <div>
            <form>
                <input type="text"/>
                <input type="text"/>
                <button type="submit">Log In</button>
            </form>
        </div>
    )
}

export default LoginComponent
