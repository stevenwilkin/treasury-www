function formatPrice(f) {
  return f.toFixed(2);
}

function initPrices(prices) {
  var template = $('#template-price').html();

  for(var price in prices) {
    var item = template
                 .replace(/__NAME__/, price)
                 .replace(/__ID__/, 'price-' + price)
                 .replace(/__PRICE__/, formatPrice(prices[price]));
    $('.prices').append(item);
  }
}

function updatePrices(prices) {
  var normalColour = 'rgb(178, 178, 178)',
      flashColour = '#FFF';

  for(var price in prices) {
    var item = $('.price-' + price),
        value = formatPrice(prices[price]);

    if(item.text() == value) {
      continue;
    }

    item
      .text(value)
      .css('color', flashColour);

    setTimeout(function() {
      item.css('color', normalColour);
    }, 1000);
  }
}

function handlePrices(prices) {
  if($('.prices').children().length) {
    updatePrices(prices);
  } else {
    initPrices(prices);
  }
}

$(function() {
  var ws = new WebSocket('ws://0.0.0.0:8080/ws');

  ws.onopen = function() {
    console.log('> onopen');
  }

  ws.onclose = function() {
    console.log('> onclose');
  }

  ws.onerror = function(error) {
    console.log('> onerror');
    console.log(error);
  }

  ws.onmessage = function(message) {
    var json = JSON.parse(message.data);

    if(json.prices) {
      handlePrices(json.prices);
    }
  }
});
