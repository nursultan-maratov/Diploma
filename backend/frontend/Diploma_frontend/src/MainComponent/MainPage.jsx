import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./MainPage.module.css";

export default function MainPage() {
    const [products, setProducts] = useState([]);
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            setIsAuthenticated(true);
        }
        fetchProducts();
    }, []);

    const fetchProducts = async () => {
        try {
            const response = await fetch("http://localhost:80/list-product");
            if (!response.ok) {
                throw new Error("Ошибка загрузки товаров");
            }
            const data = await response.json();
            setProducts(data);
        } catch (error) {
            console.error("Ошибка:", error);
        }
    };

    const handleBuy = async (productId) => {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                alert("Вы не авторизованы!");
                return;
            }

            const response = await fetch("http://localhost:80/buy-product", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify({
                    product_id: productId,
                    user_id: 0,
                }),
            });

            if (!response.ok) {
                throw new Error("Ошибка при покупке товара");
            }

            alert("Товар успешно куплен!");
        } catch (error) {
            console.error("Ошибка покупки:", error);
            alert("Не удалось купить товар.");
        }
    };

    return (
        <div className={styles.storeContainer}>
            <div className={styles.header}>
                <h1 className={styles.storeTitle}>Магазин Одежды</h1>
                {isAuthenticated ? (
                    <button className={styles.profileButton} onClick={() => navigate("/profile")}>
                        Профиль
                    </button>
                ) : (
                    <button className={styles.registerButton} onClick={() => navigate("/auth")}>
                        Войти
                    </button>
                )}
            </div>
            <div className={styles.productGrid}>
                {products.length > 0 ? (
                    products.map((product) => (
                        <div key={product.id} className={styles.productCard}>
                            <div className={styles.product_image_container}>
                                <img src={product.img || "default_image.jpg"} alt={product.name} className={styles.productImage} />
                            </div>
                            <h2 className={styles.productName}>{product.name}</h2>
                            <p className={styles.productDescription}>{product.description}</p>
                            <p className={styles.productPrice}>{product.price} ₸</p>
                            <button className={styles.buyButton} onClick={() => handleBuy(product.id)}>
                                Купить
                            </button>
                        </div>
                    ))
                ) : (
                    <p>Загрузка товаров...</p>
                )}
            </div>
        </div>
    );
}
