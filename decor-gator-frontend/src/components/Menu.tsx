import React, {createRef} from 'react'
import { useNavigate } from "react-router-dom";
import './Menu.css'
import bcrypt from 'bcryptjs'

export default function Menu({onClick} : {onClick:React.MouseEventHandler<HTMLButtonElement>}) {
    
    const loginEmailRef = React.createRef<HTMLInputElement>()
    const loginPasswordRef = React.createRef<HTMLInputElement>()

    const signupFirstNameRef = React.createRef<HTMLInputElement>()
    const signupLastNameRef = React.createRef<HTMLInputElement>()
    const signupEmailRef = React.createRef<HTMLInputElement>()
    const signupPasswordRef = React.createRef<HTMLInputElement>()

    const navigate = useNavigate();
    const goToBuyPage = () => { //check if user is authenticated
        navigate('/BuyPage');
    }

    function handleLogin() {

        if (loginEmailRef.current && loginPasswordRef.current) {
            if (!loginEmailRef.current.value || !loginPasswordRef.current.value) {
                alert("Please enter credentials.")
                return
                }
            fetch('http://localhost:8080/user/'+loginEmailRef.current.value).then((res) => {
                return res.json()
            }).then((response) => {
                console.log(response)
                let status = 'ERROR'
                if (loginPasswordRef.current) {
                    bcrypt.compare(loginPasswordRef.current.value, response.password).then((result) => {
                        console.log(result)
                        if (result) {
                            status = 'SUCCESS'
                            alert("Login successful!")
                            localStorage.setItem('userData', JSON.stringify(response)) // store userData in local storage for whenever we need info about the current user
                            console.log(JSON.parse(localStorage.getItem('userData') || ""))
                            goToBuyPage()
                        }
                        else {
                            status = 'INVALID PASSWORD'
                            alert("Invalid password")
                        }
                    }).catch((err) => {
                        console.log(err)
                    })
                }

            }).catch((err) => {
                console.log(err)
            })
        }
    }

    function handleSignup() {
        if (signupFirstNameRef.current && signupLastNameRef.current && signupEmailRef.current && signupPasswordRef.current) {
            if (!signupFirstNameRef.current.value || !signupLastNameRef.current.value || !signupEmailRef.current.value || !signupPasswordRef.current.value) {
                alert("Please enter all fields.")
                return
            }
            let username = signupFirstNameRef.current.value + signupLastNameRef.current.value
            let email = signupEmailRef.current.value
            let password = signupPasswordRef.current.value
            let signupObj = {username, email, password}

            fetch('http://localhost:8080/user/'+email).then((res) => {
                return res.json()
            }).then((response) => {
                console.log(response)
                if (response != 'User does not exist')
                    return false
                else
                    return true
            }).then((isAvailableEmail) => {
                if (!isAvailableEmail)
                    alert("Email already registered.")
                else {
                    fetch('http://localhost:8080/user', {
                        method: "POST",
                        headers: {'content-type': 'application/json'},
                        body:JSON.stringify(signupObj)
                    }).then((response)=>{
                        console.log(response.json())
                        alert("User successfully created!")
                    }).catch((err) => {
                    console.log(err)
                    })
                }
                
            })
            console.log(signupObj)
        }
    }

    function handlePasswordEmail() {
        if (!loginEmailRef.current?.value) {
            alert("Enter an email address first!")
        } else {
            fetch('http://localhost:8080/emails/'+loginEmailRef.current.value, {
            method: "PUT",
            headers: {'content-type': 'application/json'},
            body:JSON.stringify(loginEmailRef.current.value)
            }).then((response)=>{
            console.log(response)
            alert("Password reset email successfully handled.")
            }).catch((err) => {
            console.log(err)
        })
        }
    }

    return (
  
        <div className = 'menu'>
            
            <div className = 'menu-back-signup'>
                <h2 className = 'signup-title'>Sign Up</h2>

                <div className = 'signup-name-headers'>
                    <label className="first-name-text">First Name</label>
                    <label className="last-name-text">Last Name</label>
                </div>

                <div className = 'signup-name-inputs'>
                    <input ref={signupFirstNameRef} type='text' placeholder='First Name' className="signup-box-small"/>
                    <input ref={signupLastNameRef} type='text' placeholder='Last Name' className="signup-box-small"/>
                </div>
                
                <label className="signup-text">Email</label>
                <input ref={signupEmailRef} type='text' placeholder='Email' className="login-box"/>

                <label className="signup-text">Password</label>
                <input ref={signupPasswordRef} type='password' placeholder='Password' className="login-box"/>

                <button type="submit" className="signup-button" onClick = {handleSignup}>
                    SIGN UP
                </button>

                <button className = 'flip-button' onClick = {onClick}>
                    Already have an account? LOG IN
                </button>
            </div>

            <div className = 'menu-front-login'>
                <h2 className = 'login-title'>Login</h2>

                <label className="login-text">Email</label>
                <input ref={loginEmailRef} type='text' placeholder='Email' className="login-box"/>

                <label className="login-text">Password</label>
                <input ref={loginPasswordRef} type='password' placeholder='Password' className="login-box"/>
                
                <button type="submit" onClick = {handleLogin} className="login-button"> 
                    SIGN IN
                </button>

                <button type="submit" className="forgot-password-link" onClick = {handlePasswordEmail}>
                    Forgot password?
                </button>

                <button className = 'flip-button' onClick = {onClick}>
                    Don't have an account? SIGN UP
                </button>
            </div>
            

        </div>
    )
};