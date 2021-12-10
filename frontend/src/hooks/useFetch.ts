// Lägg till custom hooks
import axios from 'axios'
import {useEffect, useState} from 'react'
// For fetching a site on load
export const useFetch = async (url: string) => {
    const [response, setResponse] = useState<any | null>(null) 
    const [error, setError] = useState<any | null>(null)

    useEffect(() => {
        const doFetch = async () => {
            const res = await useGet(url)
            setResponse(res.response)
            setError(res.error)
        }
        doFetch()
    })

    return {response, error}
}

export const useGet = async (url: string) => {
    const [response, setResponse] = useState<any | null>(null)
    const [error, setError] = useState<any | null>(null)

    try {
        const res = await axios.get(url)
        const json = await res.data.json()
        setResponse(json)
    } catch (err) {
        setError(err)
    }
    return {response, error}
}

export const usePost = async (url: string, data: object) => {
    const [response, setResponse] = useState<any | null>(null)
    const [error, setError] = useState<any | null>(null)

    try {
        const res = await axios.post(url, data)
        const json = await res.data.json()
        setResponse(json)
    } catch (err) {
        setError(err)
    }
    return {response, error}
}
