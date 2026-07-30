import { getToken, clearAuthData } from "./token.js";

async function isLogged() {

    const token = getToken();

    if (!token) {
        window.location.replace("/login.html");
        return;
    }

    const url = '/api/v1/auth/super/islogged';

    try {

        const response = await fetch(url, {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token}`
            }
        });

        if (!response.ok) {
            clearAuthData();
            window.location.replace("/login.html");
            return;
        }

        const data = await response.json();

        if (!data.valid) {
            clearAuthData();
            window.location.replace("/login.html");
            return;
        }

        return true;

    } catch (err) {
        console.error(err);
    }
}

async function loadHome() {

    const token = getToken();

    if (!token) {
        return true;
    }

    const url = '/api/v1/auth/super/islogged';

    try {

        const response = await fetch(url, {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token}`
            }
        });

        const data = await response.json();

        if (data.valid) {
            window.location.replace("/home.html");
            return;
        }

        return true;

    } catch (err) {
        console.error(err);
    }
}

const page = document.body.dataset.page;

if(page){
    console.log(page);
    switch (page) {
        case "login":
            loadHome();
            break;
    }
} else {
    isLogged();
}