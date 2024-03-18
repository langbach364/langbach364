window.onload = function () {
  anime({
    targets: ".icon_introduction",
    translateX: 120,
    rotate: 540,
    duration: 2000,
  });
  anime({
    targets: ".introduction",
    translateX: 120,
    rotate: 0,
    duration: 2000,
  });
  anime({
    targets: ".header_right",
    translateX: -100,
    rotate: 0,
    duration: 2000,
  });
  anime({
    targets: ".card, #name_product",
    translateY: -200,
    delay: anime.stagger(100),
  });
  document.querySelectorAll(".card").forEach(function (card) {
    card.addEventListener("mouseover", function () {
      anime({
        targets: card,
        translateX: 0,
        scale: 1.2,
      });
    });

    card.addEventListener("mouseout", function () {
      anime({
        targets: card,
        translateX: 0,
        scale: 1.0,
      });
    });
  });

  document.getElementById('click_login').addEventListener('click', function (e) {
    e.preventDefault();
    location.href = "/Client/login/login.html";
  })


  var dataWatch = {
    start:0,
    limit:6
  };

  var dataBag = {
    start:6,
    limit:12
  };

  fetch('http://localhost:8080/View', {
    mode: "cors",
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(dataWatch),
  })
    .then(response => response.json())
    .then(data => {
      var productElements = data.map((product) => {
        return (
          `<div class="card" style="width: 14rem;">
        <img src="./imgae/anh-dong-ho-dep-51-1024x576.jpg" class="card-img-top" alt="Product" />
        <div class="card-body">
            <h5 class="card-title">${product.Type}</h5>
            <p class="card-text">${product.Name}</p>
            <a id="btn" href="#" class="btn btn-primary">Nhấn để mua</a>
        </div>
      </div>`
        );
      }).join('');

      document.querySelector('.block_card_watch').insertAdjacentHTML('beforeend', productElements);

      anime({
        targets: ".card, #name_product",
        translateY: -200,
        delay: anime.stagger(100),
      });
      document.querySelectorAll(".card").forEach(function (card) {
        card.addEventListener("mouseover", function () {
          anime({
            targets: card,
            translateX: 0,
            scale: 1.2,
          });
        });

        card.addEventListener("mouseout", function () {
          anime({
            targets: card,
            translateX: 0,
            scale: 1.0,
          });
        });
      });
    })
    .catch((error) => {
      console.error('Error:', error);
    });

    fetch('http://localhost:8080/View', {
      mode: "cors",
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(dataBag),
    })
      .then(response => response.json())
      .then(data => {
        var productElements = data.map((product) => {
          return (
            `<div class="card" style="width: 14rem;">
          <img src="./imgae/Tui xach.jpg" class="card-img-top" alt="Product" />
          <div class="card-body">
              <h5 class="card-title">${product.Type}</h5>
              <p class="card-text">${product.Name}</p>
              <a id="btn" href="#" class="btn btn-primary">Nhấn để mua</a>
          </div>
        </div>`
          );
        }).join('');
  
        document.querySelector('.block_card_bag').insertAdjacentHTML('beforeend', productElements);
  
        anime({
          targets: ".card, #name_product",
          translateY: -200,
          delay: anime.stagger(100),
        });
        document.querySelectorAll(".card").forEach(function (card) {
          card.addEventListener("mouseover", function () {
            anime({
              targets: card,
              translateX: 0,
              scale: 1.2,
            });
          });
  
          card.addEventListener("mouseout", function () {
            anime({
              targets: card,
              translateX: 0,
              scale: 1.0,
            });
          });
        });
      })
      .catch((error) => {
        console.error('Error:', error);
      });
};

