import React, {useState} from 'react'
import Header from "../../components/header";
import {useNavigate} from "react-router-dom";
import axios from "axios";
import {toast, ToastContainer} from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';


const Register: React.FC = () => {
    const [registEmail, setRegistEmail] = useState<string>('')
    const [registPassword, setRegistPassword] = useState<string>('')
    const [registUsername, setRegistUsername] = useState<string>('')
    const [registConfirmPassword, setRegistConfirmPassword] = useState<string>('')
    const navigate = useNavigate()

    const handleRegister = async(e: React.FormEvent) => {
        e.preventDefault()

        if (registPassword !== registConfirmPassword) {
            toast.error('Password not match')
            return
        }

        const baseUrl = import.meta.env.VITE_REACT_APP_API_BASE_URL
        const response = await axios.post(baseUrl + '/auth/register', {
            email: registEmail,
            password: registPassword,
            username: registUsername
        })
        if (response.status == 200) {
           toast.success("Register Success")
            setTimeout(() => {
                navigate("/login")
            }, 1000)
        } else {
            toast.error("Register Failed " + response.data.message)
        }
    }

    return (
        <>
            <Header isAuthorized={false}/>
            <div className="is-flex is-justify-content-center mt-5">
                <form className="box container" onSubmit={handleRegister}>
                    <div className="field">
                        <label className="label">Username</label>
                        <div className="control">
                            <input className="input" type="username" onChange={(e) => setRegistUsername(e.target.value)}
                                   placeholder="Username" required/>
                        </div>
                    </div>

                    <div className="field">
                        <label className="label">Email</label>
                        <div className="control">
                            <input className="input" type="email" onChange={(e) => setRegistEmail(e.target.value)}
                                   placeholder="Email" required/>
                        </div>
                    </div>

                    <div className="field">
                        <label className="label">Password</label>
                        <div className="control">
                            <input className="input" type="password" onChange={(e) => setRegistPassword(e.target.value)}
                                   placeholder="Password" required/>
                        </div>
                    </div>

                    <div className="field">
                        <label className="label">Confirm Password</label>
                        <div className="control">
                            <input className="input" type="password" onChange={(e) => setRegistConfirmPassword(e.target.value)}
                                   placeholder="Password" required/>
                        </div>
                    </div>

                    <div className="field">
                        <button className="button is-primary is-float-right" type="submit">Sign Up</button>
                    </div>
                </form>
            </div>
            <ToastContainer position="top-right" autoClose={5000} hideProgressBar={false} newestOnTop closeOnClick rtl={false} pauseOnFocusLoss draggable pauseOnHover />
        </>
    )
}

export default Register