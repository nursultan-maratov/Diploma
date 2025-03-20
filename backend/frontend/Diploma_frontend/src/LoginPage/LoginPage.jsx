import styles from "./LoginPage.module.css";
import {useState} from "react";
import {useNavigate} from "react-router-dom";

export default function LoginPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError("");

        try {
            const response = await fetch("http://localhost:80/auth", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({email, password}),
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.message || "Ошибка авторизации");
            }

            localStorage.setItem("token", data.token);
            alert("Вы успешно вошли!");
            navigate("/store");
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
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
                    <label>Пароль</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    {error && <p className="error-message">{error}</p>}
                    <button type="submit" className={styles.loginButton} disabled={loading}>
                        {loading ? "Вход..." : "Войти"}
                    </button>
                </form>
                <div className={styles.buttonContainer}>
                    <button className={styles.registerLink} onClick={() => navigate("/register")}>Регистрация</button>
                    <button className={styles.backButton} onClick={() => navigate("/store")}>Назад на главную</button>
                </div>
            </div>
        </div>
    );
}
