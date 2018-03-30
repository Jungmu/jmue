let loginModal = document.getElementById("login-modal");

function openLoginModal() {
    loginModal.style.display = "block";
}

function closeLoginModal() {
    loginModal.style.display = "none";
}

let alertModal = document.getElementById("alert");

function openAlert(text) {
    alertModal.style.display = "block";
    document.getElementById("alert-text").innerHTML=text;
}

function closeAlert() {
    alertModal.style.display = "none";
}

window.onclick = (event) => {
    if(event.target == loginModal) {
        loginModal.style.display = "none";
        
    }
    if(event.target == alertModal) {
        alertModal.style.display = "none";
    }
}