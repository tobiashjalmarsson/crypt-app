import  {FC} from 'react'
import './styles.scss'
import {Link} from 'react-router-dom'

const NavBarComponent: FC = () => {

    return (
        <div className="headerContainer">
            <Link className="mainLink" to="/">Start</Link>
            <Link className="mainLink" to="/login">Login</Link>
            <Link className="mainLink" to="/upload">Upload</Link>
            <Link className="mainLink" to="/">Groups</Link>
            <Link className="mainLink" to="/">Settings</Link>
        </div>
    )
}

export default NavBarComponent
