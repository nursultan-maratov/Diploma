import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MainPage from "./MainComponent/MainPage";
import LoginPage from "./LoginPage/LoginPage";

export default function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<MainPage />} />
                <Route path="/login" element={<LoginPage />} />
            </Routes>
        </Router>
    );
}
