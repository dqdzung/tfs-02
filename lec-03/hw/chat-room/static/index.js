const name = document.querySelector(".name");
const message = document.querySelector("#message-input");
const nameInput = document.querySelector("#name-input");
const content = document.querySelector(".content");
const sendBtn = document.querySelector(".send-btn");
const chatForm = document.querySelector(".chat-form");
const counter = document.querySelector("#counter");

chatForm.onsubmit = (e) => {
	e.preventDefault();

	fetch("http://localhost:8080/chat", {
		method: "POST",
		headers: {
			Accept: "application/json",
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			name: nameInput.value,
			message: message.value,
		}),
	});

	message.value = "";
	renderMessage();
};

const renderMessage = async () => {
	content.innerHTML = "";
	const results = await getMessage();
	console.log(results);

	for (let result of results) {
		const html = `
    <div class="message-wrapper">
      <span class="name">${result.name}</span>
      <span class="message-content">${result.message}</span>
    </div>`;
		content.innerHTML += html;
	}
};

const getMessage = async () => {
	const res = await fetch("http://localhost:8080/message");
	return res.json();
};

message.onkeyup = (e) => {
	if (e.target.value.length) {
		counter.innerHTML = "160/160";
	}
	counter.innerHTML = `${160 - e.target.value.length}/160`;
};
