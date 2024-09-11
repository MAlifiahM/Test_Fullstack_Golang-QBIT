import React from 'react'
import axios from "axios";

interface ProductCardProps {
    id : number
    name: string
    image: string
    description: string
    price: number
    category: string
    stock: number
    isAuthorized: boolean
    showSuccessNotification: (message: string) => void
    showErrorNotification: (message: string) => void
}

const ProductCard: React.FC<ProductCardProps> = ({ id, name, description, category, image, price, stock, isAuthorized, showSuccessNotification, showErrorNotification}) => {
    const [quantity, setQuantity] = React.useState<number>(1)

    const handleAddToCart = async () => {
        if (!isAuthorized) {
            showErrorNotification('Please login first')
        } else {
            try {
                const baseUrl = import.meta.env.VITE_REACT_APP_API_BASE_URL
                await axios.post(baseUrl + '/cart', {
                        product_id: id,
                        quantity: quantity
                    },
                    {
                        headers : {
                            Authorization: `Bearer ${localStorage.getItem('token')}`
                        }
                    })
                showSuccessNotification(`Successfully added ${name} to cart`)
                setQuantity(1)
                // eslint-disable-next-line @typescript-eslint/no-unused-vars
            } catch (error) {
                showErrorNotification(`Failed to add ${name} to cart`)
            }
        }
    }

    const handleBuyNow = async() => {
        if (!isAuthorized) {
            showErrorNotification('Please login first')
        } else {
            try {
                const baseUrl = import.meta.env.VITE_REACT_APP_API_BASE_URL
                await axios.post(baseUrl + '/order', {
                        product_id: id,
                        quantity: quantity
                    },
                    {
                        headers : {
                            Authorization: `Bearer ${localStorage.getItem('token')}`
                        }
                    })
                showSuccessNotification(`Successfully purchased ${name}`)
                setQuantity(1)
                // eslint-disable-next-line @typescript-eslint/no-unused-vars
            } catch (error) {
                showErrorNotification(`Failed purchased ${name}`)
            }
        }
    }

    return (
        <>
            <div className="card container is-flex-grow-1">
                <div className="card-image">
                    <figure className="image is-square m-3">
                        <img src={image} alt={`product-${name}`} className=""/>
                    </figure>
                </div>
                <div className="card-content">
                    <div className="title is-size-5">{name}</div>
                    <div className="subtitle is-size-6 mt-1">
                        <div className="is-size-7 my-2 tag is-info">{category}</div>
                        <div>IDR. {price}</div>
                        <div>Current Stock: {stock}</div>
                    </div>
                    <p className="">{description}</p>
                </div>
                <div className="card-footer is-align-items-center">
                    <div className="card-footer-item">
                        <button className="button is-info is-outlined is-size-7" onClick={handleAddToCart}>Add to cart
                        </button>
                    </div>
                    <div className="card-footer-item">
                        <input
                            type="number"
                            className="input is-small has-text-centered"
                            value={quantity}
                            onChange={(e) => setQuantity(Number(e.target.value))}
                            min={1}
                            max={stock}
                            style={{width: '60px'}}
                        />
                    </div>
                    <div className="card-footer-item">
                        <button className="button is-success is-outlined is-size-7" onClick={handleBuyNow}>Buy now
                        </button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default ProductCard