import React from 'react'
import './Menu.css'

export default function Menu({onClick} : {onClick:React.MouseEventHandler<HTMLButtonElement>}) {
    return (
        <div className = 'menu'>
            
            <div className = 'menu-back-signup'>
                Signup
                <button className = 'flip-button-temp' onClick = {onClick}>hi</button>
            </div>
            <div className = 'menu-front-login'>
                Front
                <button className = 'flip-button-temp' onClick = {onClick}>hi</button>
            </div>
            

        </div>
    )
}