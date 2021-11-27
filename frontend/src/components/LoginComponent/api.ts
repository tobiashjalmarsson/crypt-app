import axios from 'axios'

const urls = {
    getUsers : "http://localhost:8080/users",
    login: "http://localhost:8080/login"
}

interface User {
    email: string,
    id: number,
    token: string
}

export const APICallLogin = async (email: string, password: string) : Promise<boolean>  => {
    try {
        const response = await axios.post(urls.login, {
            email: email,
            password: password
        })
        console.log("Status is: ", response.status)
        if (response.status === 202) {
            // TODO Add middleware
            // TODO Add register
            const user : User = response.data
            sessionStorage.setItem('user', JSON.stringify(user))
            return true
        } else {
            return false
        }

    } catch (err) {
        console.log("Error occured while login in")
        return false
    }
}

export const APICallRegister = (username: string, email: string, password: string) : boolean => {
    console.log("test")
    return true
}
