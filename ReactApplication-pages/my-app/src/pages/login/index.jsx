import {Link} from 'react-router-dom'


import { useState } from 'react';
import { LayoutComponents, Tooltip } from '../../components/LayoutComponents';
import { ApiLogin } from '../../services/api';

export const Login = () => {
	const [email, setEmail] = useState("")
	const [password, setPassword] = useState("")
	const [message, setMessage] = useState("")

	async function SendLogin(email, password) {
		var response = (await ApiLogin(email, password))
		if (response.error) {
			setMessage(response.error)
	} else {
		setMessage("")
	}
}

	return (
		<LayoutComponents>
			<form className="login-form">
				<span className="login-form-title">Bem Vindo!</span>

					<div className="wrap-input">
						<input className={email !== "" ? 'has-val input' : 'input' }
							type="text"
							value={email}
							onChange={e => setEmail(e.target.value)}
						/>
						<span className="focus-input" data-placeholder="Email"></span>
						{message === "User not found" ? <Tooltip text={message} icon="bi bi-exclamation-triangle-fill"></Tooltip> : null}
				</div>

				<div className="wrap-input">
					<input className={password !== "" ? 'has-val input' : 'input'}
						type="password"
						value={password}
						onChange={e => setPassword(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Password"></span>
					{message === "Wrong password" ? <Tooltip text={message} icon="bi bi-exclamation-triangle-fill"></Tooltip> : null}
				</div>

				<div className="container-login-form-btn">
					<button type="button" className="login-form-btn" onClick={
						() => SendLogin(email, password)}>Login</button>
				</div>

				<div className="text-center">
					<span className="txt1">Não possui conta?</span>
					<Link className="txt2" to="/register">Criar conta.</Link>
				</div>
			</form>
		</LayoutComponents>
	)
}
