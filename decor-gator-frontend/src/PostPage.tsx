import React from "react";
import { useState, useEffect } from 'react'
import { useNavigate } from "react-router-dom";
import './PostPage.css'

export default function PostPage() {

    const postTitleRef = React.createRef<HTMLInputElement>()
    const postTypeRef = React.createRef<HTMLSelectElement>()
    const postPriceRef = React.createRef<HTMLInputElement>()
    const postDescriptionRef = React.createRef<HTMLInputElement>()
    const imageUploadRef = React.createRef<HTMLInputElement>()

    const [image, setImage] = useState<File | null>()
    const [imagePreview, setPreview] = useState<string | null>()

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

    useEffect(() => {
        if (image) {
            const fileReader = new FileReader()
            fileReader.onloadend = () => {
                setPreview(fileReader.result as string)
            }
            fileReader.readAsDataURL(image)
        } else {
            setPreview(null)
        }
    }, [image])

    return (
        <div className = 'post-editor'>
            <div className = 'text-entries'>
                <input ref={postTitleRef} type='text' placeholder='Title' className='post-title' />
                <select ref={postTypeRef} className='post-furniture-type'>   
                    <option value='' selected disabled></option>
                    <option value='Chair'>Chair</option>
                    <option value='Sofa'>Sofa</option>
                    <option value='Table'>Table</option>
                    <option value='Desk'>Desk</option>
                    <option value='Appliance'>Appliance</option>
                    <option value='Bedding'>Bedding</option>
                    <option value='Decoration'>Decoration</option>
                    <option value='Storage'>Storage</option>
                    <option value='Lighting'>Lighting</option>
                    <option value='Other'>Other</option>
                </select>
                <input ref={postPriceRef} type='text' placeholder='Price' className='post-price'/>
                <input ref={postDescriptionRef} type='text' placeholder='Description' className='post-description'/>

                <button type='button' onClick={uploadPost} className='post-submit-button'>Submit Post</button>
            </div>

            <form className = 'image-renderer'>
                {imagePreview ? <img onClick= {() => {setImage(null)}} className='image-display' src={imagePreview} /> :
                <button onClick={(event) => {
                    event.preventDefault()
                    imageUploadRef.current?.click()
                }} className='file-upload-display'>Upload Image</button>
                }
                
                <input onChange={(event) => {
                    let file = null
                    if (event.target.files)
                        file = event.target.files[0]
                    setImage(file)

                }} ref={imageUploadRef} type='file' accept='image/*' className='file-upload'/>
            </form>     
        </div>
    )
}