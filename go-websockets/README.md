- Rodar com `go run main.go`

# Websockets Mensagem

- Iniciar uma aba no navegador e no DevTools do Chrome e digitar
	```javascript
	let socket = new WebSocket("ws://localhost:3000/ws");
	socket.onmessage = (event) => { console.log("received from the server: ", event.data) };
	```
- Iniciar outra aba no navegador e no DevTools do Chrome e digitar a mesma coisa
- Usar `socket.send("Hello")` em uma das abas para enviar uma mensagem
- Verificar na outra aba

# Websockets Feed

- Iniciar uma aba no navegador e no DevTools do Chrome e digitar
	```javascript
	let socket = new WebSocket("ws://localhost:3000/feed");
	socket.onmessage = (event) => { console.log("received from the server: ", event.data) };
	```
- Esperar ~1 segundo e verificar mensagens