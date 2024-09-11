import React, {useEffect, useState} from 'react';
import axios from "axios";
import Header from "../../components/header";
import {toast, ToastContainer} from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';
import {useNavigate} from "react-router-dom";

interface CartItem {
    id: number;
    name: string;
    image: string;
    quantity: number;
    total: number;
    product_id: number;
}

const Cart: React.FC = () => {
    const [isAuthorized, setIsAuthorized] = useState(false);
    const [carts, setCarts] = useState<CartItem[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [currentItem, setCurrentItem] = useState<CartItem | null>(null);
    const [newQuantity, setNewQuantity] = useState(0);
    const navigate = useNavigate();

    // eslint-disable-next-line react-hooks/exhaustive-deps
    const checkAuthorization = () => {
        const existToken = localStorage.getItem('token');
        setIsAuthorized(!!existToken);
        if (!existToken) {
            navigate('/login');
        }
    };

    const fetchProducts = async () => {
        try {
            const url = import.meta.env.VITE_REACT_APP_API_BASE_URL;
            const response = await axios.get(url + '/cart', {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            });
            if (response.data.data !== null) {
                setCarts(response.data.data);
            }
            // eslint-disable-next-line @typescript-eslint/no-unused-vars
        } catch (error) {
            // console.error('Error fetching products:', error);
        }
    };

    useEffect(() => {
        checkAuthorization();
        fetchProducts();
    }, [checkAuthorization]);

    const handleEditCart = (item: CartItem) => {
        setCurrentItem(item);
        setNewQuantity(item.quantity);
        setIsModalOpen(true);
    };

    const handleSaveChanges = async () => {
        if (currentItem) {
            try {
                const url = `${import.meta.env.VITE_REACT_APP_API_BASE_URL}/cart/${currentItem.id}`;
                await axios.patch(url, { product_id: currentItem.product_id, quantity: newQuantity }, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem('token')}`
                    }
                });
                setCarts(prevCarts => prevCarts.map(cart =>
                    cart.id === currentItem.id ? { ...cart, quantity: newQuantity } : cart
                ));
                setIsModalOpen(false);
                toast.success("Cart item updated successfully");
                fetchProducts().catch((error) => console.error('Error fetching products:', error));
            } catch (error) {
                console.error('Error updating cart item:', error);
                toast.error("Failed to update cart item");
            }
        }
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        setCurrentItem(null);
    };

    const handleDeleteCart = async (id: number) => {
        try {
            const url = `${import.meta.env.VITE_REACT_APP_API_BASE_URL}/cart/${id}`;
            await axios.delete(url, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            });
            toast.success("Cart item deleted successfully");
            fetchProducts().catch((error) => console.error('Error fetching products:', error)); // Fetch the cart items again to update the state
        } catch (error) {
            console.error('Error deleting cart item:', error);
            toast.error("Failed to delete cart item");
        }
    };

    const handleBuyNow = async (id: number, product_id: number, quantity: number) => {
        try {
            const url = `${import.meta.env.VITE_REACT_APP_API_BASE_URL}`;
            await axios.delete(url + `/cart/${id}`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            })
            await axios.post(url + '/order', { product_id, quantity }, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            });
            toast.success("Product added to transaction successfully");
            fetchProducts().catch((error) => console.error('Error fetching products:', error));
        } catch (error) {
            console.error('Error deleting cart item:', error);
            toast.error("Failed to purchase cart item");
        }
    };

    return (
        <>
            <Header isAuthorized={isAuthorized}/>
            <div>
                <h1 className="is-size-1 has-text-centered my-5">Product List</h1>
                <div>
                    {isModalOpen && currentItem && (
                        <div className={"modal is-active"}>
                            <div className="modal-background"></div>
                            <div className="modal-card">
                                <header className="modal-card-head">
                                    <p className="modal-card-title">Edit Cart</p>
                                    <button className="delete" aria-label="close" onClick={handleCloseModal}></button>
                                </header>
                                <section className="modal-card-body">
                                    <div className="field">
                                        <label className="label">Quantity</label>
                                        <div className="control">
                                            <input className="input" type="number" value={newQuantity} onChange={(e) => setNewQuantity(parseInt(e.target.value, 10))}/>
                                        </div>
                                    </div>
                                </section>
                                <footer className="modal-card-foot">
                                    <button className="button is-success mr-3" onClick={handleSaveChanges}>Save changes</button>
                                    <button className="button" onClick={handleCloseModal}>Cancel</button>
                                </footer>
                            </div>
                        </div>
                    )}
                    {carts.length > 0 ? (
                        carts.map(cart => (
                            <div className="box" key={cart.id}>
                                <article className="media">
                                    <div className="media-left">
                                        <figure className="image is-64x64">
                                            <img src={cart.image} alt={`product-${cart.name}`}/>
                                        </figure>
                                    </div>
                                    <div className="media-content">
                                        <div className="content">
                                            <p>
                                                <strong>{cart.name}</strong><br/>
                                                Quantity: {cart.quantity}<br />
                                                Total Price : IDR. {cart.total}
                                            </p>
                                        </div>
                                        <div className="level-right is-mobile">
                                            <div>
                                                <button className="button is-info is-small" onClick={() => handleEditCart(cart)}>Edit</button>
                                            </div>
                                            <div>
                                                <button className="button is-danger is-small" onClick={() => handleDeleteCart(cart.id)}>Delete</button>
                                            </div>
                                            <div>
                                                <button className="button is-success is-small" onClick={() => handleBuyNow(cart.id, cart.product_id, cart.quantity)}>Buy Now</button>
                                            </div>
                                        </div>
                                    </div>
                                </article>
                            </div>
                        ))
                    ) : (
                        <div className="is-flex is-justify-content-center">
                            <h1 className="is-size-1">Cart is empty</h1>
                        </div>
                    )}
                </div>
                <ToastContainer position="top-right" autoClose={5000} hideProgressBar={false} newestOnTop closeOnClick rtl={false} pauseOnFocusLoss draggable pauseOnHover />
            </div>
        </>
    );
};

export default Cart;