var tempButtonFunc;

function onSignIn(googleUser) {
    // Useful data for your client-side scripts:
    var profile = googleUser.getBasicProfile();
    console.log("ID: " + profile.getId()); // Don't send this directly to your server!
    console.log('Full Name: ' + profile.getName());
    console.log('Given Name: ' + profile.getGivenName());
    console.log('Family Name: ' + profile.getFamilyName());
    console.log("Email: " + profile.getEmail());

    // The ID token you need to pass to your backend:
    var id_token = googleUser.getAuthResponse().id_token;
    console.log("ID Token: " + id_token);

    if(profile.getEmail() != "whdrjs0@gmail.com")
    {
        signOut();
        alert("you are not admin");
    } else {
        var child = document.getElementById('nav-login').children[0];
        child.innerHTML = "logout";
        tempButtonFunc = child.onclick;
        child.onclick = function() { signOut() };
    }
};

function signOut() {
    var auth2 = gapi.auth2.getAuthInstance();
    auth2.signOut().then(function () {
        var child = document.getElementById('nav-login').children[0];
        child.onclick = tempButtonFunc;
        child.innerHTML = "admin";
        console.log('User signed out.');
    });

  }