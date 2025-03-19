import styles from "./RegisterPage.module.css";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function RegisterPage() {
    const [formData, setFormData] = useState({
        firstName: "",
        lastName: "",
        birthDate: "",
        email: "",
        password: "",
        confirmPassword: "",
    });

    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (formData.password !== formData.confirmPassword) {
            setError("Пароли не совпадают!");
            return;
        }
        console.log("Регистрация успешна!", formData);
        setError("");
        // Здесь можно добавить запрос к серверу для регистрации
    };

    return (
        <div className={styles.registerContainer}>
            <div className={styles.registerBox}>
                <h1 className={styles.registerTitle}>Регистрация</h1>
                <form className={styles.registerForm} onSubmit={handleSubmit}>
                    <input
                        type="text"
                        name="firstName"
                        placeholder="Имя"
                        value={formData.firstName}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="text"
                        name="lastName"
                        placeholder="Фамилия"
                        value={formData.lastName}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="date"
                        name="birthDate"
                        placeholder="Год рождения"
                        value={formData.birthDate}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="email"
                        name="email"
                        placeholder="Email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="password"
                        name="password"
                        placeholder="Пароль"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="password"
                        name="confirmPassword"
                        placeholder="Подтверждение пароля"
                        value={formData.confirmPassword}
                        onChange={handleChange}
                        required
                    />
                    {error && <p className="error-message">{error}</p>}
                    <button type="submit" className={styles.registerButton}>Зарегистрироваться</button>
                </form>
                <button className={styles.backButton} onClick={() => navigate("/login")}>
                    Назад
                </button>
            </div>
        </div>
    );
}
