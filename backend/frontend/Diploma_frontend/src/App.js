import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MainPage from "./MainComponent/MainPage";
import LoginPage from "./LoginPage/LoginPage";
import RegisterPage from "./RegisterPage/RegisterPage";

export default function App() {
    return (
        <Router>
            <Routes>
                <Route path="/store" element={<MainPage />} />
                <Route path="/auth" element={<LoginPage />} />
                <Route path="/register" element={<RegisterPage />} />
            </Routes>
        </Router>
    );
}
