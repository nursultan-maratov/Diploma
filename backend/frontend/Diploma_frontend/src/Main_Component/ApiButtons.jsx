import { useState } from "react";
import "./ApiButtons.css";

export default function ApiButtons() {
    const [response, setResponse] = useState("");
    const playlistID = "2Bu59D3h2JWPPrsSiWqc7b";
    const accessToken = "Bearer BQDJZFQQ_LcMMXDYihyBukw9nrUs2ksHPIm4SqdxcM-v5XZkPLi9d9WpA-KH1jp8yKo1JM8B9yE6eozY5eUGW70eyDfcco_Uj9gKsQuwWC64FVbaeqhOX9FJaPq3RDrFYch8vuAYhFo";

    const TracksRequest = async (method) => {
        const url = `https://api.spotify.com/v1/playlists/${playlistID}/tracks`;
        let options = {
            method,
            headers: {
                "Authorization": accessToken,
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

        if (method === "POST") {
            const trackUri = prompt("Введите URI трека");
            if (!trackUri) {
                setResponse("Добавление отменено");
                return;
            }

            options.body = JSON.stringify({ uris: [trackUri] });

            try {
                const res = await fetch(url, options);
                if (!res.ok) throw new Error(`Ошибка: ${res.status}`);
                setResponse("Трек успешно добавлен!");
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
                <h1 className="title">API Запросы Spotify</h1>
                <div className="button-group">
                    <button onClick={() => TracksRequest("GET")} className="button get">GET (Треки плейлиста)</button>
                    <button onClick={() => TracksRequest("POST")} className="button post">POST (Добавить трек)</button>
                    <button onClick={() => TracksRequest("DELETE")} className="button delete">DELETE (Удалить трек)</button>
                </div>
                <pre className="response">{response || "Здесь будет ответ от сервера"}</pre>
            </div>
        </div>
    );
}
