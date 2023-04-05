function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
    return false
}

$("#exit").click(function (e) {
    if(getCookie('user')) {
        document.cookie = "user=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    }

    if(getCookie('admin')) {
        document.cookie = "admin=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    }

    location.reload()
});