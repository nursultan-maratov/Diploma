import styles from "./RegisterPage.module.css";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function RegisterPage() {
    const [formData, setFormData] = useState({
        firstName: "",
        lastName: "",
        email: "",
        password: "",
        confirmPassword: "",
    });

    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (formData.password !== formData.confirmPassword) {
            setError("Пароли не совпадают!");
            return;
        }

        setLoading(true);
        setError("");

        try {
            const response = await fetch("http://localhost:80/user/create", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    first_name: formData.firstName,
                    last_name: formData.lastName,
                    email: formData.email,
                    password: formData.password,
                }),
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.message || "Ошибка регистрации");
            }

            alert("Регистрация успешна! Перенаправление на страницу входа...");
            navigate("/auth");
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
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
                    <button type="submit" className={styles.registerButton} disabled={loading}>
                        {loading ? "Регистрация..." : "Зарегистрироваться"}
                    </button>
                </form>
                <button className={styles.backButton} onClick={() => navigate("/auth")}>
                    Назад
                </button>
            </div>
        </div>
    );
}
