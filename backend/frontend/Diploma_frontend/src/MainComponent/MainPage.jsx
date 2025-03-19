import { useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./MainPage.module.css";
import Adidas_Jacket from "../imageSet/Adidas_Jacket.jpg";
import Levi_Jeans from "../imageSet/Leivs_Jeans.jpg";
import Nike_Tshirt from "../imageSet/Nike_Tshirt.jpeg";
import North_Face_shapka from "../imageSet/North_Face_shapka.jpg";
import Puma_crossfits from "../imageSet/Puma_crossfeats.jpg";

const products = [
    { id: 1, name: "Футболка Nike", price: 12500, image: Nike_Tshirt },
    { id: 2, name: "Джинсы Levi's", price: 25000, image: Levi_Jeans },
    { id: 3, name: "Куртка Adidas", price: 35000, image: Adidas_Jacket },
    { id: 4, name: "Кроссовки Puma", price: 30000, image: Puma_crossfits },
    { id: 5, name: "Шапка North Face", price: 15000, image: North_Face_shapka }
];

export default function MainPage() {
    const [cart, setCart] = useState([]);
    const navigate = useNavigate();

    const handleBuy = (product) => {
        setCart([...cart, product]);
        alert(`${product.name} добавлен в корзину!`);
    };

    return (
        <div className={styles.storeContainer}>
            <div className={styles.header}>
                <h1 className={styles.storeTitle}>Магазин Одежды</h1>
                <button className={styles.registerButton} onClick={() => navigate("/auth")}>Войти</button>
            </div>
            <div className={styles.productGrid}>
            {products.map((product) => (
                    <div key={product.id} className={styles.productCard}>
                        <div className={styles.product_image_container}>
                            <img src={product.image} alt={product.name} className={styles.productImage}/>
                        </div>
                        <h2 className={styles.productName}>{product.name}</h2>
                        <p className={styles.productPrice}>{product.price} ₸</p>
                        <button onClick={() => handleBuy(product)} className={styles.buyButton}>
                            Купить
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
}