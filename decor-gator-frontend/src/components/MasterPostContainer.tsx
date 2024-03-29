import './MasterPostContainer.css'
import PostContainer from './PostContainer'

type MasterPostContainerProps = {
    postContainers: {
        id: number
        title: string
        furnitureType: string
        posterUsername: string
        price: string
        imageURL: string
        description: string
    }[]
    clickDisplayEvent: Function
}

export default function MasterPostContainer(props: MasterPostContainerProps) {
    return (
        <div className = 'master-container'>
                {props.postContainers?.map(post => {
                    return <PostContainer key={post.id} title={post.title} furnitureType={post.furnitureType}
                        posterUsername={post.posterUsername} price={post.price} id={post.id} imageURL={post.imageURL} description={post.description}
                        clickDisplayEvent = {props.clickDisplayEvent}/>
                })}
        </div>
    )
}