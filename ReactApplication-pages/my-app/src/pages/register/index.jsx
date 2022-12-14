import { useState } from 'react';
import {Link} from 'react-router-dom'

import { LayoutComponents } from "../../components/LayoutComponents"
import { ApiRegister } from '../../services/api';


export const Register = () => {
	const [name, setName] = useState("")
	const [email, setEmail] = useState("")
	const [password, setPassword] = useState("")

	return (
		<LayoutComponents>
			<form className="login-form">
				<span className="login-form-title">Criar Conta</span>

				<div className="wrap-input">
					<input className={name !== "" ? 'has-val input' : 'input' }
						type="text"
						value={name}
						onChange={e => setName(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Nome"></span>
				</div>

				<div className="wrap-input">
					<input className={email !== "" ? 'has-val input' : 'input' }
						type="email"
						value={email}
						onChange={e => setEmail(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Email"></span>
				</div>

				<div className="wrap-input">
					<input className={password !== "" ? 'has-val input' : 'input' }
						type="password"
						value={password}
						onChange={e => setPassword(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Password"></span>
				</div>

				<div className="container-login-form-btn">
					<button type="button" className="login-form-btn" onClick={ () => {
						ApiRegister(name, email, password)
					}}>Cadastrar</button>
				</div>

				<div className="text-center">
					<span className="txt1">Já possui conta?</span>
					<Link className="txt2" to="/login">Logar.</Link>
				</div>
			</form>
		</LayoutComponents>
	)
}
