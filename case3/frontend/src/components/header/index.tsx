import React from 'react'
import Logo from "../../assets/logo.svg"

interface HeaderProps {
    isAuthorized: boolean
}

const Header: React.FC<HeaderProps> = ({ isAuthorized }) => {
    const [isActive, setIsActive] = React.useState(false)

    const handleBurgerClick = () => {
        setIsActive(!isActive)
    }

    const handleLogout = () => {
        localStorage.removeItem('token')
        window.location.reload()
    }
    return (
        <>
            <nav className="navbar px-5" role="navigation">
                <div className="navbar-brand is-align-items-center mr-5">
                    <a href="/" className="">
                        <div className="is-flex is-align-items-center">
                            <img src={Logo} alt="home-logo" className="image is-48x48 mr-2"/>
                            <span className="has-text-primary">Store</span>
                        </div>
                    </a>
                    <a role="button"
                       className={`navbar-burger ${isActive ? 'is-active' : ''}`}
                       aria-label="menu"
                       aria-expanded={isActive}
                       data-target="navbarMenu"
                       onClick={handleBurgerClick}
                    >
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                    </a>
                </div>
                <div className={`navbar-menu ${isActive ? 'is-active' : ''}`} id="navbarMenu">
                    <div className="navbar-start">
                        <a href="/" className="navbar-item">Home</a>
                        <a href="/product" className="navbar-item">Product</a>
                    </div>
                    {
                        isAuthorized === false ?
                            <div className="navbar-end">
                                <div className="navbar-item">
                                    <a className="button is-primary is-outlined" href="/login"> Login / Register</a>
                                </div>
                            </div>
                            :
                            <div className="navbar-end">
                                <div className="navbar-item">
                                    <a className="button is-success is-outlined" href="/transaction">Transaction</a>
                                </div>
                                <div className="navbar-item">
                                    <a className="button is-info is-outlined" href="/cart">Cart</a>
                                </div>
                                <div className="navbar-item">
                                    <div className="button is-danger is-outlined" onClick={handleLogout}>Logout</div>
                                </div>
                            </div>
                    }
                </div>

            </nav>
        </>
    )
}

export default Header