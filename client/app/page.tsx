"use client";

import { useEffect, useState } from "react";

import Link from "next/link";
import { Alert, Snackbar } from "@mui/material";

import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { SetGlobalError } from "@/lib/game/globalErrorSlice";

export default function Home() {
  const globalErrorData = useAppSelector((state) => state.globalError);
  const dispatch = useAppDispatch();

  const [open, setOpen] = useState<boolean>(false);

  const fetchHome = async () => {
    try {
      const res = await fetch("http://localhost:8000/", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        mode: "cors",
        credentials: "include",
      });
      const data = await res.json();
      console.log(data);
      if (!res.ok) {
        throw res;
      }
    } catch (err: unknown) {
      if (err instanceof Response) {
        console.error(err);
        dispatch(
          SetGlobalError({
            message: String(err.statusText),
            statusCode: err.status,
          })
        );
        setOpen(true);
      }
    }
  };

  useEffect(() => {
    fetchHome();
  }, []);

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <>
      <main className="main-home side-anim center-wrapper">
        <Snackbar open={open} autoHideDuration={5000} onClose={handleClose}>
          <Alert
            onClose={handleClose}
            severity="error"
            variant="filled"
            sx={{ width: "100%" }}
          >
            <b>{globalErrorData.message}</b> - Status code:{" "}
            {globalErrorData.statusCode}
          </Alert>
        </Snackbar>
        <h2 className="main-title">Click the planet</h2>
        <article className="main-subtitle">Simple clicker game</article>
        <article className="home-article">
          <h4>Start your clicking journey right now!</h4>
          <Link href="/register">
            <button className="submit-button home-register-button">
              Register
            </button>
          </Link>
          <p className="home-login-info">
            Already have an account? <Link href="/login">Login here</Link>
          </p>
        </article>
      </main>
    </>
  );
}
