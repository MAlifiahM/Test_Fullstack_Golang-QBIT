import React, {useEffect, useState} from 'react'
import Header from "../../components/header";
import axios from "axios";
import ProductCard from "../../components/card_product";
import {toast, ToastContainer} from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';


const Product: React.FC = () => {
    const [isAuthorized, setIsAuthorized] = useState(false)
    const [products, setProducts] = useState([])
    const [refetchProducts, setRefetchProducts] = useState(false)

    useEffect(() => {
        const checkAuthorization = () => {
            const existToken = localStorage.getItem('token');
            if (existToken) {
                setIsAuthorized(true);
            } else {
                setIsAuthorized(false);
            }
        };

        const fetchProducts = async () => {
            try {
                const url = import.meta.env.VITE_REACT_APP_API_BASE_URL;
                const response = await axios.get(url + '/product');
                setProducts(response.data.data);
            } catch (error) {
                console.error('Error fetching products:', error);
            }
        };

        checkAuthorization();
        fetchProducts().catch(
            (error) => console.error('Error fetching products:', error),
        );
    }, []);

    const showSuccessNotification = (message: string) => {
        toast.success(message)
        setRefetchProducts(true)
    }

    const showErrorNotification = (message: string) => {
        toast.error(message)
    }

    useEffect(() => {
        const fetchProducts = async () => {
            try {
                const url = import.meta.env.VITE_REACT_APP_API_BASE_URL;
                const response = await axios.get(url + '/product');
                setProducts(response.data.data);
            } catch (error) {
                console.error('Error fetching products:', error);
            }
        };

        fetchProducts().catch((error) => console.error('Error fetching products:', error));
        setRefetchProducts(false);
    }, [refetchProducts]);

    return (
        <>
            <Header isAuthorized={isAuthorized}/>
            <div className="">
                <h1 className="is-size-1 has-text-centered my-5">Product List</h1>
                <div className="columns is-multiline">
                    {
                        products.length > 0 ?
                            products.map((product: any) => (
                                    <div className="column is-3 is-flex is-flex-direction-column mx-2" key={product.id}>
                                        <ProductCard
                                            id={product.id}
                                            price={product.price}
                                            name={product.name}
                                            description={product.description}
                                            category={product.category}
                                            image={product.image}
                                            stock={product.stock}
                                            isAuthorized={isAuthorized}
                                            showSuccessNotification={showSuccessNotification}
                                            showErrorNotification={showErrorNotification}
                                        />
                                    </div>
                            ))
                            :
                            <p>No products found</p>
                            }
                </div>
                <ToastContainer position="top-right" autoClose={5000} hideProgressBar={false} newestOnTop closeOnClick rtl={false} pauseOnFocusLoss draggable pauseOnHover />
            </div>
        </>
    )
}

export default Product