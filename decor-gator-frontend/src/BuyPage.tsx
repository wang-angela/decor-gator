import React from "react";
import { useNavigate } from "react-router-dom";
import PostPage from "./PostPage"
import {useState, useEffect} from 'react';
import MasterPostContainer from './components/MasterPostContainer'
import PostContainer from './components/MasterPostContainer'

function BuyPage() {
    
    const [postsToDisplay, updateDisplayedPosts] = useState<any[]>([])

    function getAllPosts() {
        
        let postArray:Array<any> = []

        fetch('http://localhost:8080/posts').then((res) => {
            return res.json()
        }).then((response) => {
        console.log("I should be here first")
        console.log(response)
        response.forEach((element: any) => {

            let title = element.title
            let furnitureType = element.furnitureType
            let posterUsername = element.userPosted
            let price = '$20'
            let id = element.id

            let postObj = {id, title, furnitureType, posterUsername, price}
            console.log(postObj)

            postArray.push(postObj)
        });
    }).then(() => {
        console.log("This is what I am logging")
        console.log(postArray)
        updateDisplayedPosts(postArray)
    })

    }

    useEffect(() => {
        getAllPosts()
    }, [])

    const navigate = useNavigate();
    return (
        <div className='buy-container'>
            <MasterPostContainer postContainers={postsToDisplay}/>
            <button type="button" className="makePost-button" onClick={()=>navigate('/PostPage')}>
            + Post
            </button>
            <button type="button" className="viewPosts-button" onClick={getAllPosts}>
            List Posts
            </button>
        </div>
        
    )
}

export default BuyPage