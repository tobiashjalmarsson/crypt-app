import React from 'react';
import './App.scss';
import LoginRegistration     from './views/LoginRegistration';
import UploadComponent from './components/UploadComponent/UploadComponent';
import {
    BrowserRouter as Router,
    Routes,
    Route
} from 'react-router-dom'

function App() {
  return (
    <Router>
        <div className="App">
            <Routes>
                {/*Add startpage*/}
                <Route path="/" element={<LoginRegistration />}/>
                <Route path="/upload" element={<UploadComponent />} />
            </Routes>
        </div>
    </Router>
  );
}

export default App;
