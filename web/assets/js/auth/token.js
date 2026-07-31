export function saveAuthData(data) {
    localStorage.setItem('token', data.token);
    localStorage.setItem('superName', data.super_name);
    localStorage.setItem('superId', data.super_id);
}

export function getToken() {
    return localStorage.getItem('token');
}

export function getSuperName() {
    return localStorage.getItem('superName');
}

export function clearAuthData() {
    localStorage.removeItem('token');
    localStorage.removeItem('superName');
    localStorage.removeItem('superId');
}

export function logout(event) {
    if (event) {
        event.preventDefault();
    }
    clearAuthData();
    window.location.replace("/pages/super/login.html");
}