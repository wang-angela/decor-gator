import React from 'react'
import './styles.css'
import { useNavigate } from 'react-router-dom';

const LoginField = () => {
    const navigate = useNavigate();

    const goToDashboard = () => {
        navigate ('/redirect');
    };

  return <form className='input'>
        Username:&nbsp;
        <input type='input' placeholder='Username' className="Input-box"/>
        
        Password:&nbsp;
        <input type='input' placeholder='Password' className="Input-box"/>
        
        <button id="loginButton" type="submit" className="Input-submit" onClick={goToDashboard}>
            Log in
        </button>
    </form>;
}

export default LoginField

//rafce shortcut