
import './styles.css';

export const LayoutComponents = (props) => {
	return (
		<div className="container">
			<div className="container-login">
				<div className="wrap-login">
					{props.children}
				</div>
			</div>
		</div>
	)
}

export const Tooltip = (props) => {
	return (
		<div className='tooltip'>
			<i className={"tooltip__icon " + props.icon}></i>
			<span className='tooltip__message'>{props.text}</span>
		</div>
	)
}
