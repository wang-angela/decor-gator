import React from "react";
import { useNavigate } from "react-router-dom";
import Header from "./Header"

//We might need class instead of function; In that case, find out how to use navigate in class component
function BuyPage() {
    
    //Build more structure
    const navigate = useNavigate();
    return (
        <div>
            <Header />cd 
            <button type="button" className="makePost-button" onClick={()=>navigate('/PostPage')}>
                + Post
            </button>
        </div>
    )
}

export default BuyPage