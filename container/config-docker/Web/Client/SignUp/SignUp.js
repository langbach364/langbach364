document.getElementById('submit').addEventListener('click', function(event) {
    event.preventDefault();

    var email = document.getElementById('typeEmailX').value;
    var password = document.getElementById('typePasswordX').value;
    var confirm = document.getElementById('typeConfirmPasswordX').value;

    var data = {
        email: email,
        password: password
    };

    if(password != confirm) {
        alert('Mật khẩu xác nhận không khớp. Vui lòng thử lại.');
    }

    fetch('http://localhost:8080/SignUp',  {
        mode:"cors",
            method:'POST',
            headers: {
            'Content-Type' : 'application/json'
        },
        body: JSON.stringify(data),
    })

    window.location.href = "/Client/login/login.html";
});
