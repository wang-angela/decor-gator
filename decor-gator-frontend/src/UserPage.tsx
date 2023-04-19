import React from "react";
import { useState, useEffect } from 'react'
import { useNavigate, useLocation } from "react-router-dom";

export default function UserPage() {

    const {state} = useLocation();
    const { userEmail } = state;

    function retrieveUser() {
        
    }

    return (
        <div className = 'user-editor'>
            Email: 
        </div>
    )
}
