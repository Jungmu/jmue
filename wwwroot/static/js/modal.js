var loginModal = document.getElementById("login-modal");

function openLoginModal() {
    loginModal.style.display = "block";
}

function closeLoginModal() {
    loginModal.style.display = "none";
}

window.onclick = function(event) {
    if(event.target == loginModal) {
        loginModal.style.display = "none";
    }
}