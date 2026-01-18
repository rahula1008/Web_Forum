import { Link } from "react-router";
import { useAuth } from "../../auth/useAuth";
import "./Header.css";

const UserLoggedInHeader = () => {
    const { user, logout } = useAuth();

    const handleLogout = async () => {
        await logout();
    };

    return (
        <header>
            <div className="left-section">
                <Link to="/" className="header-link">Home</Link>
            </div>
            <div className="right-section">
                <p>Hi {user && user.username}</p>
                <button className="header-action" type="button" onClick={handleLogout}>
                    Log Out
                </button>
            </div>
        </header>
    )
}

const UserNotLoggedInHeader = () => {
    return (
        <header>
            <div className='left-section'>
                <Link to="/" className="header-link">
                    PeerPrep
                </Link>
            </div>
            <div className='right-section'>
                <Link className="header-action" to='/login'>Log In</Link>
            </div>
        </header>
    );
}

export default function Header() {
    const { isAuthed } = useAuth();

    return (
        isAuthed ? <UserLoggedInHeader /> : <UserNotLoggedInHeader />
    )
}
