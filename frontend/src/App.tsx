import React from 'react';
import './App.scss';
import LoginRegistration     from './views/LoginRegistration';
import NavBarComponent from './components/NavBarComponent/NavBarComponent';
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
            <NavBarComponent />
            <div className="mainContainer">
            <Routes>
                <Route path="/login" element={<LoginRegistration />}/>
                <Route path="/upload" element={<UploadComponent />} />
            </Routes>
            </div>
        </div>
    </Router>
  );
}

export default App;
