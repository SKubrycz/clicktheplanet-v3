import Link from "next/link";
import "./Navbar.scss";

export default function Navbar() {
  return (
    <nav className="main-navbar">
      <div className="main-navbar-left">
        <Link href="/">LOGO</Link>
      </div>
      <div className="main-navbar-right">
        <Link href="/login">Login</Link>
        <Link href="/register">Register</Link>
      </div>
    </nav>
  );
}
