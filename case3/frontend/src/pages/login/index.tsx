import React, { useState } from 'react'
import Header from "../../components/header";
import axios from 'axios'
import { useNavigate } from "react-router-dom";
import {toast, ToastContainer} from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';

const Login: React.FC = () => {
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const navigate = useNavigate()

    const handleLogin = async(e: React.FormEvent) => {
        e.preventDefault()
        try {
            const baseUrl = import.meta.env.VITE_REACT_APP_API_BASE_URL
            const response = await axios.post(baseUrl + '/auth/login', {
                email: email,
                password: password
            })
            localStorage.setItem('token', response.data.data.token)
            navigate('/')
            // eslint-disable-next-line @typescript-eslint/no-unused-vars
        } catch (error) {
            toast.error("Login Failed check your email and password")
        }
    }

    return (
        <>
            <Header isAuthorized={false}/>
            <div className="is-flex is-justify-content-center mt-5">
               <form className="box container" onSubmit={handleLogin}>
                   <div className="field">
                       <label className="label">Email</label>
                       <div className="control">
                           <input className="input" type="email" onChange={(e) => setEmail(e.target.value)} placeholder="Email" required />
                       </div>
                   </div>

                   <div className="field">
                       <label className="label">Password</label>
                       <div className="control">
                           <input className="input" type="password" onChange={(e) => setPassword(e.target.value)} placeholder="Password" required />
                       </div>
                   </div>
                   <div className="field">
                       <div className="is-float-left">
                           Don't have an account? sign up <a href="/register" className=""> here</a>
                       </div>
                       <button className="button is-primary is-float-right" type="submit">Submit</button>
                   </div>
               </form>
            </div>
            <ToastContainer position="top-right" autoClose={5000} hideProgressBar={false} newestOnTop closeOnClick rtl={false} pauseOnFocusLoss draggable pauseOnHover />
        </>
    )
}

export default Login