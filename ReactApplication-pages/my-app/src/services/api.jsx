// import axios from "axios"

const urlRegister = 'http://localhost:8080/register'
const urlLogin = 'http://localhost:8080/login'
const urlAuthenticate = 'http://localhost:8080/authenticate'

export const ApiRegister = async (name, email, password) => {
	const update = {
		username: name,
		email: email,
		password: password,
	};

	const options = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(update)
	};
	let response = await fetch(urlRegister, options)
	.then(encoded => encoded.json())
	.then(response => response)
	console.log(response)
	return (response)
}

export const ApiLogin = async (email, password) => {
	const update = {
		email: email,
		password: password,
	};

	const options = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(update)
	};
	let response = await fetch(urlLogin, options)
	.then(encoded => encoded.json())
	.then(response => response)
	console.log(response)
	return response
}

export const ApiAuthenticate = async (email, password, token) => {
	const update = {
		email: email,
		password: password,
		token: token,
	};

	const options = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(update)
	};
	let response = await fetch(urlAuthenticate, options)
	.then(encoded => encoded.json())
	.then(response => response)
	return (response)
}

