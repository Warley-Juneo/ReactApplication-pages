import {Link} from 'react-router-dom'
// import jpIMG from './imagem'

import { useState } from 'react';
import { LayoutComponents } from '../../components/LayoutComponents';
import { ApiAuthenticate } from '../../services/api';

export const Authenticate = () => {
	const [email, setEmail] = useState("")
	const [password, setPassword] = useState("")
	const [token, setToken] = useState("")

	return (
		<LayoutComponents>
			<form className="login-form">
				<span className="login-form-title">Insert Token</span>

				<div className="wrap-input">
					<input className={email !== "" ? 'has-val input' : 'input' }
						type="text"
						value={email}
						onChange={e => setEmail(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Email"></span>
				</div>

				<div className="wrap-input">
					<input className={password !== "" ? 'has-val input' : 'input'}
						type="password"
						value={password}
						onChange={e => setPassword(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Password"></span>
				</div>

				<div className="wrap-input">
					<input className={token !== "" ? 'has-val input' : 'input' }
						type="text"
						value={token}
						onChange={e => setToken(e.target.value)}
					/>
					<span className="focus-input" data-placeholder="Token"></span>
				</div>

				<div className="container-login-form-btn">
					<button type="button" className="login-form-btn" onClick={ () => {
						ApiAuthenticate(email, password, token)
					}}>Authenticate</button>

				</div>

				<div className="text-center">
					<span className="txt1">Não possui conta?</span>
					<Link className="txt2" to="/register">Criar conta.</Link>
				</div>
			</form>
		</LayoutComponents>
	)
}
