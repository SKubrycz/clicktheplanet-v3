import { ReactNode } from "react";
import Navbar from "@/components/Navbar/Navbar";

import "./home.scss";
import Link from "next/link";

export default function Home() {
  return (
    <>
      {/* <Navbar></Navbar> */}
      <main className="main-home side-anim">
        <h2 className="main-title">Click the planet</h2>
        <article className="main-subtitle">Simple clicker game</article>
        <article className="home-article">
          <h4>Start your clicking journey right now!</h4>
          <Link href="/register">
            <button className="home-register-button">Register</button>
          </Link>
          <p className="home-login-info">
            Already have an account? <Link href="/login">Login here</Link>
          </p>
        </article>
      </main>
    </>
  );
}
