import { Link } from "react-router";
import { useAuth } from "../../auth/useAuth";
import "./Header.css";

const UserLoggedInHeader = () => {
    const { user } = useAuth();
    return (
        <header>
            <div className="left-section">
                <Link to="/" className="header-link">Home</Link>
            </div>
            <div className="right-section">
                <p>Hi {user && user.username}</p>
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
                <Link to='/login'>Login</Link>
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
