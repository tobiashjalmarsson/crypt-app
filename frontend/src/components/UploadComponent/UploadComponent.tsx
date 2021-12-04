import {FC, useState} from 'react'
import './styles.scss'

const UploadComponent: FC = () => {
    const [passkey, setPasskey] = useState("")
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

    const handleFileChange = (e : any) => {
        setSelectedFile(e.target.files[0])
        setIsFilePicked(true)
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
