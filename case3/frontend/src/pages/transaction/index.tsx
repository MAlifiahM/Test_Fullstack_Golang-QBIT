import React, {useEffect, useState} from 'react'
import Header from "../../components/header";
import axios from "axios";
import {useNavigate} from "react-router-dom";

interface TrxItem {
    id: number;
    name: string;
    image: string;
    quantity: number;
    total: number;
    product_id: number;
}

const Transaction: React.FC = () => {
    const [isAuthorized, setIsAuthorized] = useState(false)
    const [transactionsData, setTransactionsData] = useState<TrxItem[]>([])
    const navigate = useNavigate();

    // eslint-disable-next-line react-hooks/exhaustive-deps
    const checkAuthorization = () => {
        const existToken = localStorage.getItem('token');
        setIsAuthorized(!!existToken);
        if (!existToken) {
            navigate('/login');
        }
    };

    const fetchTransactions = async () => {
        try {
            const url = import.meta.env.VITE_REACT_APP_API_BASE_URL;
            const response = await axios.get(url + '/order', {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            });

            if (response.data.data !== null) {
                setTransactionsData(response.data.data);
            }
            // eslint-disable-next-line @typescript-eslint/no-unused-vars
        } catch (error) {
            // console.error('Error fetching products:', error);
        }
    };

    useEffect(() => {
        checkAuthorization();
        fetchTransactions()
    }, [checkAuthorization]);
    return (
        <>
            <Header isAuthorized={isAuthorized}/>
            <div>
                <h1 className="is-size-1 has-text-centered my-5">Transaction List</h1>
                <div>
                    {transactionsData.length > 0 ? (
                        transactionsData.map(trx => (
                            <div className="box" key={trx.id}>
                                <article className="media">
                                    <div className="media-left">
                                        <figure className="image is-64x64">
                                            <img src={trx.image} alt={`product-${trx.name}`}/>
                                        </figure>
                                    </div>
                                    <div className="media-content">
                                        <div className="content">
                                            <p>
                                                <strong>{trx.name}</strong><br/>
                                                Quantity: {trx.quantity}<br />
                                                Total Price : IDR. {trx.total}
                                            </p>
                                        </div>
                                    </div>
                                </article>
                            </div>
                        ))
                    ) : (
                        <div className="is-flex is-justify-content-center">
                            <h1 className="is-size-1">Transaction is empty</h1>
                        </div>
                    )}
                </div>
            </div>
        </>
    )
}

export default Transaction