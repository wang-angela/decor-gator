import React from "react";
import { useNavigate } from "react-router-dom";
import PostPage from "./PostPage"

function listPosts() {
    fetch('http://localhost:8080/posts').then((res) => {
        return res.json()
    }).then((response) => {
        console.log(response)
        response.forEach((element: any) => {
            console.log(element.userPosted)
            console.log(element.furnitureType)
            console.log(element.title)
        });
    })
}

function BuyPage() {
    
    
    const navigate = useNavigate();
    return (
        <div className='buy-container'>
            <button type="button" className="makePost-button" onClick={()=>navigate('/PostPage')}>
            + Post
            </button>
            <button type="button" className="viewPosts-button" onClick={listPosts}>
            List Posts
            </button>
        </div>
        
    )
}

export default BuyPage