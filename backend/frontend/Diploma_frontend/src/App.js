import {BrowserRouter as Router, Routes, Route, Navigate} from "react-router-dom";
import MainPage from "./MainComponent/MainPage";
import LoginPage from "./LoginPage/LoginPage";
import RegisterPage from "./RegisterPage/RegisterPage";
import ProfilePage from "./ProfilePage/ProfilePage";

export default function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Navigate to="/store" replace />} />
                <Route path="/store" element={<MainPage />} />
                <Route path="/auth" element={<LoginPage />} />
                <Route path="/register" element={<RegisterPage />} />
                <Route path="/profile" element={<ProfilePage />} />
            </Routes>
        </Router>
    );
}
