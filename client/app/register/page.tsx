"use client";

import { FormEvent, useState } from "react";

import Link from "next/link";

import {
  FormControl,
  Input,
  InputLabel,
  InputAdornment,
  IconButton,
  ThemeProvider,
} from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";

import { defaultTheme } from "../../assets/defaultTheme";
import { useRouter } from "next/navigation";

export default function Register() {
  const [login, setLogin] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [passwordAgain, setPasswordAgain] = useState<string>("");

  const [showPassword, setShowPassword] = useState<boolean>(false);
  const [showAgainPassword, setShowAgainPassword] = useState<boolean>(false);

  const router = useRouter();

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleClickShowAgainPassword = () => {
    setShowAgainPassword(!showAgainPassword);
  };

  const handleRegister = async (e: FormEvent) => {
    e.preventDefault();

    await fetch("http://localhost:8000/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        login: login,
        email: email,
        password: password,
        againPassword: passwordAgain,
      }),
      mode: "cors",
      credentials: "include",
    })
      .then((req) => {
        if (!req.ok) {
          console.log(req);
        } else {
          console.log(req);
          router.push("/");
        }
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <main className="center-wrapper">
        <article className="form-panel">
          <h3 className="form-title">Register</h3>
          <div className="center-flex-col">
            <form className="center-flex-col" method="POST">
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-login">Login</InputLabel>
                <Input
                  id="register-login"
                  autoComplete="username"
                  onChange={(e) => setLogin(e.target.value)}
                ></Input>
              </FormControl>
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-email">Email</InputLabel>
                <Input
                  id="register-email"
                  autoComplete="email"
                  onChange={(e) => setEmail(e.target.value)}
                ></Input>
              </FormControl>
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-password">Password</InputLabel>
                <Input
                  id="register-password"
                  type={showPassword ? "text" : "password"}
                  onChange={(e) => setPassword(e.target.value)}
                  endAdornment={
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="toggle password visibility"
                        onClick={handleClickShowPassword}
                      >
                        {showPassword ? <VisibilityOff /> : <Visibility />}
                      </IconButton>
                    </InputAdornment>
                  }
                />
              </FormControl>
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-again-password">
                  Password again
                </InputLabel>
                <Input
                  id="register-again-password"
                  type={showAgainPassword ? "text" : "password"}
                  onChange={(e) => {
                    setPasswordAgain(e.target.value);
                  }}
                  endAdornment={
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="toggle password visibility"
                        onClick={handleClickShowAgainPassword}
                      >
                        {showAgainPassword ? <VisibilityOff /> : <Visibility />}
                      </IconButton>
                    </InputAdornment>
                  }
                />
              </FormControl>
              <p className="form-prompt">
                Already have an existing account?{" "}
                <Link href="/login">Login here</Link>
              </p>
              <button
                type="submit"
                onClick={(e) => handleRegister(e)}
                className="submit-button"
              >
                Register
              </button>
            </form>
          </div>
        </article>
      </main>
    </ThemeProvider>
  );
}
