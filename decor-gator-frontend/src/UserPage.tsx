import React from "react";
import { useState, useEffect } from 'react'
import { useNavigate, useLocation } from "react-router-dom";
import "./UserPage.css"
import bcrypt from 'bcryptjs'

export default function UserPage() {

    const {state} = useLocation();
    let [email, reloadEmail] = useState(JSON.parse(localStorage.getItem('userData') || "").email);
    const [showChangeEmail, setShowChangeEmail] = useState(false);
    const [showChangePassword, setShowChangePassword] = useState(false);

    const newEmailRef = React.createRef<HTMLInputElement>()
    const newPasswordRef = React.createRef<HTMLInputElement>()
    const enteredPasswordRef1 = React.createRef<HTMLInputElement>()
    const enteredPasswordRef2 = React.createRef<HTMLInputElement>()

    const updateChangeEmail = () => setShowChangeEmail(!showChangeEmail);

    const updateChangePassword = () => setShowChangePassword(!showChangePassword);

    function updateEmail() {
        if (newEmailRef.current && enteredPasswordRef1.current) {
            if (!newEmailRef.current.value || !enteredPasswordRef1.current.value){
                alert("Please enter all credentials.")
                return
            }
            
            let newEmail = newEmailRef.current.value
            let enteredPassword = enteredPasswordRef1.current.value
            fetch('http://localhost:8080/user/'+email).then((res) => {
                return res.json()
            }).then((response) => {
                console.log(response)
                if (response == 'User does not exist')
                    return false
                else   
                    return response
            }).then((user) => {
                if (!user)
                    console.log("Something went wrong! User does not exist")
                else
                {
                    let status = 'ERROR'
                    if (enteredPasswordRef1.current) {
                        bcrypt.compare(enteredPasswordRef1.current.value, user.password).then((result) => {
                            console.log("Password correct?", result)
                            if (result) {
                                status = 'SUCCESS'
                                console.log("password correct!")
                                
                                let currentusername = user.username
                                fetch('http://localhost:8080/user/'+email, {
                                    method: "PUT",
                                    headers: {'content-type': 'application/json'},
                                    body: JSON.stringify({username: currentusername, email: newEmail, password: enteredPassword})
                                }).then((response)=>{
                                    console.log(response.json())
                                    alert("Email successfully updated!")
                                    reloadEmail(newEmail)
                                }).catch((err) => {
                                    console.log(err)
                                })
                            }
                            else {
                                status = 'INVALID PASSWORD'
                                alert("Invalid password")
                            }
                        })
                    }
                    
                }
            })
        }
        
    }

    function updatePassword() {
        if (newPasswordRef.current && enteredPasswordRef2.current) {
            if (!newPasswordRef.current.value || !enteredPasswordRef2.current.value){
                alert("Please enter all credentials.")
                return
            }
            
            let newPassword = newPasswordRef.current.value
            //let enteredPassword = enteredPasswordRef.current.value
            let currentPassword = null;

            fetch('http://localhost:8080/user/'+email).then((res) => {
                return res.json()
            }).then((response) => {
                console.log(response)
                if (response == 'User does not exist')
                    return false
                else   
                    return response
                    
            }).then((user) => {
                if (!user)
                    console.log("Something went wrong! User does not exist")
                else
                {
                    let status = 'ERROR'
                    if (enteredPasswordRef2.current) {
                        currentPassword = user.password;
                        bcrypt.compare(enteredPasswordRef2.current.value, currentPassword).then((result) => {
                            console.log("Password correct?", result)
                            if (result) {
                                status = 'SUCCESS'
                                console.log("password correct!")
                                
                                console.log("Changing password...")
                                
                                let currentusername = user.username
                                let currentemail = user.email
                                fetch('http://localhost:8080/user/'+email, {
                                    method: "PUT",
                                    headers: {'content-type': 'application/json'},
                                    body: JSON.stringify({username: currentusername, email: currentemail, password: newPassword})
                                }).then((response)=>{
                                    console.log(response.json())
                                    alert("Password successfully updated!")
                                }).catch((err) => {
                                    console.log(err)
                                })
                            }
                            else {
                                status = 'INVALID PASSWORD'
                                alert("Invalid password")
                            }
                        })
                    }
                }
            })
        }
    }

    return (
        <div className = 'post-editor'>
            <div className = 'edit-user'>
                <button type="button" className="makePost-button" onClick={updateChangeEmail}>
                Change Email
                </button>
                <div className='text-entries'>
                    Current Email: {email}
                </div>
                <div>
                    {showChangeEmail &&
                    <div>
                        <input ref={newEmailRef} type='text' placeholder='New Email' className="search-text-input"/>
                        <input ref={enteredPasswordRef1} type='text' placeholder='Password' className="search-text-input"/>
                        <button type="button" className="makePost-button" onClick={updateEmail}>
                            Enter
                        </button>
                    </div>
                    }
                </div>
                
            </div>
            <div className = 'edit-user'>
                <button type="button" className="makePost-button" onClick={updateChangePassword}>
                    Change Password
                </button>
                <div>
                    {showChangePassword &&
                    <div>
                        <input ref={enteredPasswordRef2} type='text' placeholder='Current Password' className="search-text-input"/>
                        <input ref={newPasswordRef} type='text' placeholder='New Password' className="search-text-input"/>
                        <button type="button" className="makePost-button" onClick={updatePassword}>
                            Enter
                        </button>
                    </div>
                    }
                </div>
            </div>
            
        </div>
    )
}
