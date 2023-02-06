import React from 'react'
import './styles.css'
import { useNavigate } from 'react-router-dom';

const LoginField = () => {
    const navigate = useNavigate();

    const goToDashboard = () => {
        navigate ('/redirect');
    };

  return <form className='box'>
        <h2 className="title">Login</h2>

        <label className="label">Username</label>
        
        <input type='text' placeholder='Username' className="login_box"/>
        
        <label className="label">Password</label>
        <input type='text' placeholder='Password' className="login_box"/>
        
        <button id="loginButton" type="submit" className="login_submit" onClick={goToDashboard}>
            SIGN IN
        </button>
    </form>;
}

export default LoginField

//rafce shortcut