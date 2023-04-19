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
        <div className = 'post-overlay'>
            <div className = 'post-window-2'>
            <div className = 'text-entries-2'>
                <div className = 'display-header'>
                <label className='post-title-2'>{props.title}</label>
                <label className='post-furniture-type-2'>{props.furnitureType}</label>
                <label className='post-price-2'>{props.price}</label>
                </div>
                <label className='post-owner-prefix'>Posted by </label>
                <label className='post-owner'>{props.posterUsername}</label>
                <p className='post-description-2'>{props.description}</p>

                <button type='button' onClick={() => {props.clickDisplayEvent(null)}} className='post-submit-button-2'>← Back</button>
            </div>

            <form className = 'image-renderer-2'>
                <img className='image-display-2' src={props.imageURL} /> 
            </form>     
        </div>
        </div>
    )
}