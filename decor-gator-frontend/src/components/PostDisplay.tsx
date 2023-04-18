type PostDisplayProps = {
    id: any
    title: string
    furnitureType: string
    posterUsername: string
    price: string
}

export default function PostDisplay(props: PostDisplayProps) {
    return (
        <div className = 'post-window'>
            <div className = 'text-entries'>
                <label className='post-title'>{props.title}</label>
                <label className='post-furniture-type'>{props.furnitureType}</label>
                <label className='dollar-sign'>$</label>
                <label className='post-price'>{props.price}</label>
                <p className='post-description'>Test</p>

                <button type='button' className='post-submit-button'>Create Post</button>
            </div>

            <form className = 'image-renderer'>
                <img className='image-display' src={'https://cdn.discordapp.com/attachments/726320415827427360/1090294211645018192/image.png'} /> 
            </form>     
        </div>
    )
}