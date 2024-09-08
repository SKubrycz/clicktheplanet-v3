"use client";

import { useState } from "react";

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

export default function Register() {
  const [showPassword, setShowPassword] = useState<boolean>(false);
  const [showAgainPassword, setShowAgainPassword] = useState<boolean>(false);

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleClickShowAgainPassword = () => {
    setShowAgainPassword(!showAgainPassword);
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <main className="center-wrapper">
        <article className="form-panel">
          <h3 className="form-title">Register</h3>
          <div className="center-flex-col">
            <FormControl sx={{ width: "20em" }} variant="standard">
              <InputLabel htmlFor="register-login">Login</InputLabel>
              <Input id="register-login"></Input>
            </FormControl>
            <FormControl sx={{ width: "20em" }} variant="standard">
              <InputLabel htmlFor="register-email">Email</InputLabel>
              <Input id="register-email"></Input>
            </FormControl>
            <FormControl sx={{ width: "20em" }} variant="standard">
              <InputLabel htmlFor="register-password">Password</InputLabel>
              <Input
                id="register-password"
                type={showPassword ? "text" : "password"}
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
            <button type="submit" className="submit-button">
              Register
            </button>
          </div>
        </article>
      </main>
    </ThemeProvider>
  );
}
