import './PostDisplay.css'

type PostDisplayProps = {
    id: number
    title: string
    furnitureType: string
    posterUsername: string
    price: string
    imageURL: string
    description: string
    clickDisplayEvent: Function
}

export default function PostDisplay(props: PostDisplayProps) {
    return (
        <div className = 'post-window'>
            <div className = 'text-entries'>
                <label className='post-title'>{props.title}</label>
                <label className='post-furniture-type'>{props.furnitureType}</label>
                <label className='dollar-sign'>$</label>
                <label className='post-price'>{props.price}</label>
                <p className='post-description'>{props.description}</p>
                <p className='post-owner'>Posted by {props.posterUsername}</p>

                <button type='button' onClick={() => {props.clickDisplayEvent(null)}} className='post-submit-button'>Back</button>
            </div>

            <form className = 'image-renderer'>
                <img className='image-display' src={props.imageURL} /> 
            </form>     
        </div>
    )
}