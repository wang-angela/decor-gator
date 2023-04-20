import React from "react";
import { useNavigate, useLocation} from "react-router-dom";
import PostPage from "./PostPage"
import {useState, useEffect, useRef} from 'react';
import MasterPostContainer from './components/MasterPostContainer'
import PostDisplay from './components/PostDisplay'
import PostContainer from './components/MasterPostContainer'
import './BuyPage.css'

function BuyPage() {

    const {state} = useLocation();
    const email = state;
    
    const [postsToDisplay, updateDisplayedPosts] = useState<any[]>([])
    const [pageList, updatePageList] = useState<any[][]>([])
    const [page, setPage] = useState<number>(1)
    const [focusDisplayPost, updateFocusDisplayPost] = useState<any>(null)
    const allPostsRef = useRef<Array<any>>([])
    const postTypeRef = React.createRef<HTMLSelectElement>()
    const searchBarRef = React.createRef<HTMLInputElement>()

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
            let price = '$'+ String(element.price)
            let id = element._id
            let imageURL = element.imageURL
            let description = element.description

            let postObj = {id, title, furnitureType, posterUsername, price, imageURL, description}
            console.log(postObj)

            postArray.push(postObj)
        });
    }).then(() => {
        console.log("This is what I am logging")
        console.log(postArray)
        allPostsRef.current = postArray
        let pages = getPages(allPostsRef.current)
        updatePageList(pages)
        console.log(pages[0])
        console.log(pages[1])
        updateDisplayedPosts(pages[0])
    })

    }

    function getPages(inputPostArray:Array<any>) {
        let pageSize = 8
        let pageArray:Array<any> = []

        for (let i = 0; i < inputPostArray.length; i += pageSize) {
            let page = inputPostArray.slice(i, i + pageSize)
            pageArray.push(page)
        }

        return pageArray
    }

    function increasePage() {
        if (page === pageList.length) return
        setPage(previousPage => {
            return previousPage + 1
        })
    }

    function decreasePage() {
        if (page === 1) return
        setPage(previousPage => {
            return previousPage - 1
        })
    }

    function renderPage() {
        if (!pageList) return
        updateDisplayedPosts(pageList[page-1])
    }

    function filterBySearch() {
        let destinationCategory:any, matchedResults:any
        destinationCategory = postTypeRef.current?.value
        if (destinationCategory === '') {
            matchedResults = allPostsRef.current.filter((post) => {
                return post.title.toLowerCase().includes(searchBarRef.current?.value.toLowerCase())
            })
        } else {
            matchedResults = allPostsRef.current.filter((post) => {
                return post.title.toLowerCase().includes(searchBarRef.current?.value.toLowerCase()) && post.furnitureType === destinationCategory
            })
        }
        let newPages = getPages(matchedResults)
        updatePageList(newPages)
    }

    function newFocusPost(id:any, title:string, furnitureType:string, posterUsername:string, price:string, imageURL:string, description:string) {
        if (!id)
        updateFocusDisplayPost(null)
        else {
            let postObj = {id, title, furnitureType, posterUsername, price, imageURL, description}
            updateFocusDisplayPost(postObj)
        }
    }

    useEffect(() => {
        console.log("Rendering initial posts...")
        getAllPosts()
    }, [])

    useEffect(() => {
        renderPage()
    }, [page])

    useEffect(() => {
        if (page === 1)
            renderPage()
        else
            setPage(1)
    }, [pageList])

    useEffect(() => {

    }, [focusDisplayPost])

    const navigate = useNavigate();
    return (
        <div className='buy-container'>
            {focusDisplayPost ? <div className='overlay'>
                <PostDisplay id={focusDisplayPost.id} title={focusDisplayPost.title} furnitureType={focusDisplayPost.furnitureType} posterUsername={focusDisplayPost.posterUsername}
                price={focusDisplayPost.price} imageURL={focusDisplayPost.imageURL} description={focusDisplayPost.description} clickDisplayEvent={newFocusPost}/>
            </div> : <div className='underlay'>
            <button type="button" className="editUser-button" onClick={()=>navigate('/UserPage', { state: email})}>
            My Page
            </button>
            <button type="button" className="makePost-button" onClick={()=>navigate('/PostPage')}>
            + Post
            </button>
            <MasterPostContainer postContainers={postsToDisplay} clickDisplayEvent={newFocusPost}/>
            <div className='bottom-panel'>
            <select placeholder="Select Category" ref={postTypeRef} className='post-furniture-type-3'>   
                    <option className='default-value' value=''>All Categories</option>
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
            <input type="text" className="search-text-input" placeholder="Search Title" ref={searchBarRef}/>
            <button type="button" className="search-button" onClick={filterBySearch}>
            Search
            </button>
            <div className="change-page-container">
                <button type="button" className="previous-page-button" onClick={decreasePage}>
                «
                </button>
                <label className="page-label">
                Page {page} of {pageList.length}
                </label>
                <button type="button" className="next-page-button" onClick={increasePage}>
                »
                </button>
            </div>
            </div>
            </div>}
        </div>
    )
}

export default BuyPage