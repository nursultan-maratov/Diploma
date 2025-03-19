import styles from "./LoginPage.module.css";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function LoginPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("Email:", email, "Password:", password);
    };

    return (
        <div className={styles.loginContainer}>
            <div className={styles.loginBox}>
                <h1 className={styles.loginTitle}>Войти</h1>
                <form className={styles.loginForm} onSubmit={handleSubmit}>
                    <label>Email</label>
                    <input
                        type="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <label>Password</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <button type="submit" className={styles.loginButton}>Войти</button>
                </form>
                <div className={styles.buttonContainer}>
                    <button className={styles.registerLink} onClick={() => navigate("/register")}>Регистрация</button>
                    <button className={styles.backButton} onClick={() => navigate("/store")}>Назад на главную</button>
                </div>
            </div>
        </div>
    );
}
