document.getElementById("submit").addEventListener('click', function(event) {
    event.preventDefault();

    var email = document.getElementById('typeEmailX').value;
    var password = document.getElementById('typePasswordX').value;

    var data = {
        email: email,
        password: password
    };

    fetch('http://localhost:8080/SignIn', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })

    window.location.href = "/index/index.html";
});


