import React from 'react'
import './Menu.css'

export default function Menu({onClick} : {onClick:React.MouseEventHandler<HTMLButtonElement>}) {
    return (
        <div className = 'menu'>
            
            <div className = 'menu-back-signup'>
                Signup
                <button className = 'flip-button' onClick = {onClick}>
                    Take me to log in screen!
                </button>
            </div>

            <div className = 'menu-front-login'>
                <h2 className = 'login-title'>Login</h2>

                <label className="login-text">Username</label>
                <input type='text' placeholder='Username' className="login_box"/>

                <label className="login-text">Password</label>
                <input type='text' placeholder='Password' className="login_box"/>

                <button type="submit" className="login-button">
                    SIGN IN
                </button>

                <button className = 'flip-button' onClick = {onClick}>
                    Take me to sign up page!
                </button>
            </div>
            

        </div>
    )
}