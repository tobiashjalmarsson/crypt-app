import {FC, useState} from 'react'
import './styles.scss'
import { getBase64, encrypt, decrypt } from '../../utils/encryption'

const UploadComponent: FC = () => {
    const [file64, setFile64] = useState<any>("")
    const [passkey, setPasskey] = useState<string>("")
    const [selectedFile, setSelectedFile] = useState({
        name: "",
        size: 0,
        type: ""
    })
    const [isFilePicked, setIsFilePicked] = useState(false)
    
    const handleSubmit = (e : any) => {
        e.preventDefault()
        console.log("Submitting file")
    }

    const handleFileChange = async (e : any) => {
        try {
        console.log("original file is")
        console.log(e.target.files[0])
        setSelectedFile(e.target.files[0])
        setIsFilePicked(true)
        console.log("base64 file is")
        const newfile = await getBase64(e.target.files[0]) as string
        console.log(newfile)
        console.log("type: ", typeof newfile )
        setFile64(newfile)
        
        console.log("Check so it remains the same")
        const encryptedfile = encrypt(newfile, "hej")
        const decryptedfile = decrypt(encryptedfile, "hej")
        console.log(decryptedfile === newfile)
        } catch (err){
            console.log("ERR OCCURED")
        }
    }

    const updatePasskey = (e : any) => {
        e.preventDefault()
        setPasskey(e.target.value)
    }

    const removeFile = (e : any) => {
        e.preventDefault()
        setSelectedFile({
            name: "",
            size: 0,
            type: "" 
        })
        setIsFilePicked(false)
    }

    return (
        <div className="mainContainer">
            <h2>Upload</h2>
            { isFilePicked &&
            <div id="informationContainer">
                <p>Name: {selectedFile.name}</p>
                <p>Type: {selectedFile.type}</p>
                <p>Size (Bytes): {selectedFile.size}</p> 
                <button onClick={removeFile}>Remove</button>
            </div>
            }
            <input id="fileInput" type="file" name="file" onChange={handleFileChange}/>
            <input id="passkeyInput" type="text" name="passkey" onChange={updatePasskey} value={passkey} placeholder="Passkey"/>
            <button onClick={handleSubmit}>Upload</button> 
        </div>
    )
}

export default UploadComponent
