import { useState } from "react";
import "./ApiButtons.css";

export default function ApiButtons() {
    const [response, setResponse] = useState("");

    const TracksRequest = async (method) => {
        const url = `http://localhost:80/user/create`;
        let options = {
            method,
            headers: {
                "Content-Type": "application/json"
            }
        };

        if (method === "GET") {
            try {
                const res = await fetch(url, options);
                if (!res.ok) throw new Error(`Ошибка: ${res.status}`);
                const data = await res.json();

                const tracksInfo = data.items.map(item => ({
                    name: item.track.name,
                    duration: `${Math.floor(item.track.duration_ms / 60000)}:${String(Math.floor((item.track.duration_ms % 60000) / 1000)).padStart(2, "0")} мин`,
                    added_at: new Date(item.added_at).toLocaleString("ru-RU"),
                    uri: item.track.uri // Нужно для DELETE
                }));

                setResponse(JSON.stringify(tracksInfo, null, 2));
                return;
            } catch (error) {
                setResponse(`Ошибка запроса: ${error.message}`);
            }
        }

        if (method === "POST") {const first_name = prompt("Введите имя:");
            const last_name = prompt("Введите фамилию:");
            const email = prompt("Введите email:");
            const phone = prompt("Введите телефон:");
            const address = prompt("Введите адрес:");
            const status = prompt("Введите статус (например, active):");

            if (!first_name || !last_name || !email || !phone || !address || !status) {
                setResponse("Добавление отменено (не все поля заполнены)");
                return;
            }

            options.body = JSON.stringify({
                first_name,
                last_name,
                email,
                phone,
                address,
                status
            });


            try {
                const res = await fetch(url, options);
                if (!res.ok) throw new Error(`Ошибка: ${res.status}`);

                const data = await res.json(); // получаем число-ответ
                setResponse(`✅ Пользователь успешно добавлен! ID: ${data}`);
            } catch (error) {
                setResponse(`Ошибка запроса: ${error.message}`);
            }
        }

        if (method === "DELETE") {
            const trackUri = prompt("Введите URI трека");
            if (!trackUri) {
                setResponse("Удаление отменено");
                return;
            }

            options.body = JSON.stringify({
                tracks: [{ uri: trackUri }]
            });

            try {
                const res = await fetch(url, options);
                if (!res.ok) throw new Error(`Ошибка: ${res.status}`);
                setResponse("Трек успешно удален!");
            } catch (error) {
                setResponse(`Ошибка запроса: ${error.message}`);
            }
        }
    };

    return (
        <div className="container">
            <div className="card">
                <h1 className="title">API Запросы</h1>
                <div className="button-group">
                    <button onClick={() => TracksRequest("GET")} className="button get">GET</button>
                    <button onClick={() => TracksRequest("POST")} className="button post">POST</button>
                    <button onClick={() => TracksRequest("DELETE")} className="button delete">DELETE</button>
                </div>
                <pre className="response">{response || "Здесь будет ответ от сервера"}</pre>
            </div>
        </div>
    );
}
