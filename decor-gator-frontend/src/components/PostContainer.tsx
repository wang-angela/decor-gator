import './PostContainer.css'

type PostContainerProps = {
    id: any
    title: string
    furnitureType: string
    posterUsername: string
    price: string
}

export default function PostContainer(props: PostContainerProps) {
    return (
        <div className = 'container'>  
            <label>
                {props.title}: posted by {props.posterUsername} for a price of {props.price} with category {props.furnitureType} and id {props.id}
            </label>
        </div>
    )
}