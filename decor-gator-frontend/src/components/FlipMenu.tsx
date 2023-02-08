import React from 'react';
import './FlipMenu.css'
import Menu from './Menu'
import {CSSTransition} from 'react-transition-group'
import {useState} from 'react';

export default function FlipMenu() {
    const [showFront, setShowFront] = useState(true)
    return (
        <div className = 'flip-menu-container'>
            <CSSTransition
                in = {showFront}
                timeout = {500}
                classNames = 'flip'
            >
                <Menu onClick = {() => {
                    setShowFront((state) => !state)
                }} />
            </CSSTransition>
        </div>
    )
}