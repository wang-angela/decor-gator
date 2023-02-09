import React from 'react';
import './FlipMenu.css'
import Menu from './Menu'
import {CSSTransition} from 'react-transition-group'
import {useState} from 'react';

export default function FlipMenu() {
    const [showLogin, setShowLogin] = useState(true)
    return (
        <div className = 'flip-menu-container'>
            <CSSTransition
                in = {showLogin}
                timeout = {500}
                classNames = 'flip'
            >
                <Menu onClick = {() => {
                    setShowLogin((state) => !state)
                }} />
            </CSSTransition>
        </div>
    )
}