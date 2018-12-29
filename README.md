
### 1. Just Deploy the same on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### 2. Paste PORT to Heroku

![](images/Bot5.png)

Go to heroku dashboard, go to "Setting" -> "Config Variables".

- Add "Config Vars"
- Name -> "PORT"
- Value use  `80`.


### Use

```javascript
const addScript = (src) => {
  var elem = document.createElement("script");
  elem.src = src;
  document.head.appendChild(elem);
}

const getData = (data) => {
  conosle.log(data)
}

addScript("https://APP_ADDRESS.herokuapp.com/data?url=http://habr.ru&callback=getData")

```
