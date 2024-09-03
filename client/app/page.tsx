import { ReactNode } from "react";
import Navbar from "@/components/Navbar/Navbar";

import "./home.scss";
import Link from "next/link";

export default function Home() {
  return (
    <>
      <Navbar></Navbar>
      <main className="main-home">
        <h2>Click the planet</h2>
        <article>Simple clicker game</article>
        <article>
          <h4>Play now:</h4>
          <button className="home-register-button">
            <Link href="/register">Register</Link>
          </button>
          <p>
            Already have an account? <Link href="/login">Login here</Link>
          </p>
        </article>
      </main>
    </>
  );
}
