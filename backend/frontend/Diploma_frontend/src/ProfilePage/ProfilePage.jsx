import React, { useEffect, useState } from "react";
import styles from"./ProfilePage.module.css";

const ProfilePage = () => {
    const [userData, setUserData] = useState(null);
    const [cartItems, setCartItems] = useState([]);

    useEffect(() => {
        // Запрос данных пользователя
        fetch("http://localhost:80/user-info?user_id=0")
            .then(response => response.json())
            .then(data => setUserData(data))
            .catch(error => console.error("Ошибка загрузки данных пользователя:", error));

        // Запрос корзины покупок
        fetch("http://localhost:80/list-order?user_id=0")
            .then(response => response.json())
            .then(data => setCartItems(data))
            .catch(error => console.error("Ошибка загрузки корзины:", error));
    }, []);

    return (
        <div className={styles.profileContainer}>
            <h1>Профиль пользователя</h1>

            {userData && (
                <div className={styles.userInfo}>
                    <p><strong>Имя:</strong> {userData.name}</p>
                    <p><strong>Email:</strong> {userData.email}</p>
                    <p><strong>Дата регистрации:</strong> {new Date(userData.created_at).toLocaleDateString()}</p>
                </div>
            )}

            <h2>Корзина покупок</h2>
            <div className={styles.cartGrid}>
                {cartItems.length > 0 ? (
                    cartItems.map(item => (
                        <div key={item.id} className={styles.productCard}>
                            <img src={item.img} alt={item.description} className={styles.productImage} />
                            <h3>Товар #{item.product_id}</h3>
                            <p>{item.description}</p>
                        </div>
                    ))
                ) : (
                    <p>Корзина пуста</p>
                )}
            </div>

            <button className={styles.backButton} onClick={() => window.location.href = "/"}>
                Вернуться на главную
            </button>
        </div>
    );
};

export default ProfilePage;
