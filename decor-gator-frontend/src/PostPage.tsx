import React from "react";
import { useNavigate } from "react-router-dom";

export default function PostPage() {

    const postTitleRef = React.createRef<HTMLInputElement>()
    const postTypeRef = React.createRef<HTMLInputElement>()

    const navigate = useNavigate();
    const goToBuyPage = () => { //check if user is authenticated
        navigate('/BuyPage');
    }

    function uploadPost() {
        var Title, FurnitureType, UserPosted, postObj
        if (postTitleRef.current && postTypeRef.current) {
            Title = postTitleRef.current.value
            FurnitureType = postTypeRef.current.value
            console.log(JSON.parse(localStorage.getItem('userData') || ""))
            let a = JSON.parse(localStorage.getItem('userData') || "")
            UserPosted = JSON.parse(localStorage.getItem('userData') || "").email
            postObj = {Title, FurnitureType, UserPosted}
        }
        if (!Title || !FurnitureType) {
            alert("Please enter all fields.")
        } else {
            fetch('http://localhost:8080/posts', {
            method: "POST",
            headers: {'content-type': 'application/json'},
            body:JSON.stringify(postObj)
        }).then((response)=>{
            console.log(response)
            alert("Post successfully created!")
            goToBuyPage()
        }).catch((err) => {
            console.log(err)
        })
        }

    }

    return (
        <div className = 'post-editor'>
            <input ref={postTitleRef} type='text' placeholder='Title' className='post-title' />
            <input ref={postTypeRef} type='text' placeholder='Furniture Type' className='post-furniture-type' />
            <button type='button' onClick={uploadPost} className='post-submit-button'>Submit Post</button>
        </div>
    )
}