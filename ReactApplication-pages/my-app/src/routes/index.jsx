import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import { Login } from '../pages/login'
import { Register } from '../pages/register'
import { Authenticate } from '../pages/authenticate'

export const AppRouter = () => {
	return (
		<Router>
			<Routes>
				<Route path="/login" exact element={<Login />} />
				<Route path="/register" exact element={<Register />} />
				<Route path="/authenticate" exact element={<Authenticate />} />
			</Routes>
		</Router>
	)
}
