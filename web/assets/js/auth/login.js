import { saveToken, getToken, removeToken } from "./token.js";

async function login(event) {
    event.preventDefault();
    const superLogin = document.getElementById("super_login").value;
    const superPassword = document.getElementById("super_password").value;

    const url = '/api/v1/auth/super/login';

    const body = {
        super_login: superLogin,
        super_password: superPassword
    };

    try {

        const response = await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body)
        });

        const data = await response.json();

        if (!response.ok) {
            const alert = document.getElementById("login_alert");

            alert.textContent = data.error;
            alert.classList.remove("d-none");

            return;
        }

        // console.log(data);
        saveToken(data.token);
        window.location.replace("/home.html");

    } catch (err) {
        console.error(err);
    }
}

const form = document.getElementById("login_form");

form.addEventListener("submit", login);

export function logout(event) {
    if (event) {
        event.preventDefault();
    }

    removeToken();
    window.location.replace("/login.html");
}