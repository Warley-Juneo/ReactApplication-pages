import axios from "axios"

let urlServer = 'http://localhost:8081'

export const ApiPost = (name, email, password) => {
	const update = {
		username: name.name,
		email: email.email,
		password: password.password,
	};

	const options = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: (update),
	};

	let response = axios.post(urlServer, options)
	.then(encoded => encoded.json())
	.then(response => response)
	.catch(function (error) {
		console.log(error);
	});
	console.log(response)
	return(response)
}
