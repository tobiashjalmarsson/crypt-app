import {FC, useState} from 'react'
import './styles.scss'

const UploadComponent: FC = () => {

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

    return (
        <div className="mainContainer">
            <h2>Upload</h2>
            { isFilePicked &&
            <div id="informationContainer">
                <p>Name: {selectedFile.name}</p>
                <p>Type: {selectedFile.type}</p>
                <p>Size (Bytes): {selectedFile.size}</p> 
            </div>
            }
            <input type="file" name="file" onChange={handleFileChange}/>
            <button onClick={handleSubmit}>Upload</button> 
        </div>
    )
}

export default UploadComponent
