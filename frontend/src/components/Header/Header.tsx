import { Link } from "react-router";
import { useAuth } from "../../auth/useAuth";
import "./Header.css";

type HeaderProps = {
  typeOfPage?: string;
  centerText?: string;
};

const defaultTypeOfPage = "Home Page";


const UserLoggedInHeader = ({typeOfPage = defaultTypeOfPage, centerText}: HeaderProps) => {
    const { user, logout } = useAuth();

    const handleLogout = async () => {
        await logout();
    };

    return (
        <header>
            <div className="left-section">
                <Link to="/" className="header-link">Home</Link>
            </div>
            <div className="header">
                <p className="kicker">{typeOfPage}</p>
                <h1 className="title">{centerText || ""}</h1>
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

const UserNotLoggedInHeader = ({typeOfPage = defaultTypeOfPage, centerText}: HeaderProps) => {
    return (
        <header>
            <div className='left-section'>
                <Link to="/" className="header-link">
                    Home
                </Link>
            </div>
            <div className="header">
                <p className="kicker">{typeOfPage}</p>
                <h1 className="title">{centerText || ""}</h1>
            </div>
            <div className='right-section'>
                <Link className="header-action" to='/login'>Log In</Link>
            </div>
        </header>
    );
}

export default function Header({typeOfPage = defaultTypeOfPage, centerText}: HeaderProps) {
    const { isAuthed } = useAuth();

    return (
        isAuthed 
            ? <UserLoggedInHeader typeOfPage={typeOfPage} centerText={centerText} /> 
            : <UserNotLoggedInHeader typeOfPage={typeOfPage} centerText={centerText}/>
    )
}
