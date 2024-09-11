import React, {useEffect, useState} from 'react'
import Header from "../../components/header";

const Home: React.FC = () => {
    const [isAuthorized, setIsAuthorized] = useState(false)

    useEffect(() => {
        const existToken = localStorage.getItem('token')
        if (existToken) {
            setIsAuthorized(true)
        }else {
            setIsAuthorized(false)
        }
    }, []);

    return (
        <>
            <Header isAuthorized={isAuthorized}/>
            <div className="is-flex is-justify-content-center">
                <h1 className="is-size-1">Welcome to Store</h1>
            </div>
        </>
    )
}

export default Home