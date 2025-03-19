import { useState } from "react";
import "./MainPage.module.css";
import Adidas_Jacket from "../imageSet/Adidas_Jacket.jpg";
import Levi_Jeans from "../imageSet/Leivs_Jeans.jpg";
import Nike_Tshirt from "../imageSet/Nike_Tshirt.jpeg";
import North_Face_shapka from "../imageSet/North_Face_shapka.jpg";
import Puma_crossfits from "../imageSet/Puma_crossfeats.jpg";

const products = [
    { id: 1, name: "Футболка Nike", price: 2500, image: Nike_Tshirt },
    { id: 2, name: "Джинсы Levi's", price: 5000, image: Levi_Jeans },
    { id: 3, name: "Куртка Adidas", price: 7000, image: Adidas_Jacket },
    { id: 4, name: "Кроссовки Puma", price: 6000, image: Puma_crossfits },
    { id: 5, name: "Шапка North Face", price: 3000, image: North_Face_shapka }
];

export default function MainPage() {
    const [cart, setCart] = useState([]);

    const handleBuy = (product) => {
        setCart([...cart, product]);
        alert(`${product.name} добавлен в корзину!`);
    };

    return (
        <div className="store-container">
            <div className="header">
                <button className="register-button">Регистрация</button>
                <h1 className="store-title">Магазин Одежды</h1>
            </div>
            <div className="product-grid horizontal center">
                {products.map((product) => (
                    <div key={product.id} className="product-card">
                        <img src={product.image} alt={product.name} className="product-image" />
                        <h2 className="product-name">{product.name}</h2>
                        <p className="product-price">{product.price} ₽</p>
                        <button onClick={() => handleBuy(product)} className="buy-button">
                            Купить
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
}