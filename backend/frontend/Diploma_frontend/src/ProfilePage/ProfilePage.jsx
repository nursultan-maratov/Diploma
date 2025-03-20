import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./ProfilePage.module.css";

export default function ProfilePage() {
    const [user, setUser] = useState(null);
    const [cart, setCart] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        fetchUserData();
        fetchCart();
    }, []);

    const fetchUserData = async () => {
        try {
            const response = await fetch("http://localhost:80/user-info");
            if (!response.ok) {
                throw new Error("Ошибка загрузки данных пользователя");
            }
            const data = await response.json();
            setUser(data);
        } catch (error) {
            console.error("Ошибка:", error);
        }
    };

    const fetchCart = async () => {
        try {
            const response = await fetch("http://localhost:80/user-cart");
            if (!response.ok) {
                throw new Error("Ошибка загрузки корзины");
            }
            const data = await response.json();
            setCart(data);
        } catch (error) {
            console.error("Ошибка:", error);
        }
    };

    return (
        <div className={styles.profileContainer}>
            <h1>Профиль</h1>

            {user ? (
                <div className={styles.userInfo}>
                    <p><strong>Имя:</strong> {user.name}</p>
                    <p><strong>Email:</strong> {user.email}</p>
                </div>
            ) : (
                <p>Загрузка данных пользователя...</p>
            )}

            <h2>Корзина покупок</h2>
            <div className={styles.cartGrid}>
                {cart.length > 0 ? (
                    cart.map((product) => (
                        <div key={product.id} className={styles.productCard}>
                            <img src={product.image || "default_image.jpg"} alt={product.name} className={styles.productImage} />
                            <h3>{product.name}</h3>
                            <p>{product.description}</p>
                            <p><strong>Цена:</strong> {product.price} ₸</p>
                        </div>
                    ))
                ) : (
                    <p>Корзина пуста</p>
                )}
            </div>

            <button onClick={() => navigate("/store")} className={styles.backButton}>
                Вернуться в магазин
            </button>
        </div>
    );
}
