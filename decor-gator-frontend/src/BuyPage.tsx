import React from "react";
import { useNavigate } from "react-router-dom";
import PostPage from "./PostPage"
import {useState, useEffect, useRef} from 'react';
import MasterPostContainer from './components/MasterPostContainer'
import PostContainer from './components/MasterPostContainer'
import './BuyPage.css'

function BuyPage() {
    
    const [postsToDisplay, updateDisplayedPosts] = useState<any[]>([])
    const [pageList, updatePageList] = useState<any[][]>([])
    const [page, setPage] = useState<number>(1)
    const allPostsRef = useRef<Array<any>>([])
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
            let price = '$20'
            let id = element.id

            let postObj = {id, title, furnitureType, posterUsername, price}
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
        let pageSize = 12
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
        let matchedResults = allPostsRef.current.filter((post) => {
            return post.title.toLowerCase().includes(searchBarRef.current?.value.toLowerCase())
        })
        let newPages = getPages(matchedResults)
        updatePageList(newPages)
    }

    useEffect(() => {
        console.log("Rendering initial posts...")
        getAllPosts()
    }, [])

    useEffect(() => {
        console.log("AHHHH!!")
        renderPage()
    }, [page])

    useEffect(() => {
        console.log(pageList)
        if (page === 1)
            renderPage()
        else
            setPage(1)
    }, [pageList])

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
            <button type="button" className="previous-page-button" onClick={decreasePage}>
            Previous Page
            </button>
            <button type="button" className="next-page-button" onClick={increasePage}>
            Next Page
            </button>
            <input type="text" className="search-text-input" placeholder="Search" ref={searchBarRef}/>
            <button type="button" className="search-button" onClick={filterBySearch}>
            Search
            </button>
            <label className="page-label">
            Page {page} of {pageList.length}
            </label>
        </div>
        
    )
}

export default BuyPage