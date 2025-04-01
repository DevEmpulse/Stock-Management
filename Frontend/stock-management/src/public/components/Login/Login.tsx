"use client";
import { loginUser } from '@/features/auth/authSlice';
import { AppDispatch, RootState } from '@/store';
import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Navigate } from 'react-router';
import { TextField, Button, CircularProgress, Container, Typography } from '@mui/material';
import { AppProvider } from '@toolpad/core/AppProvider';
import logo from "/src/assets/logo.png";
const Login: React.FC = () => {
  const [credentials, setCredentials] = useState<{ email: string; contrasena: string }>({
    email: '',
    contrasena: '',
  });

  const dispatch = useDispatch<AppDispatch>();
  const { loading, error, isAuthenticated } = useSelector((state: RootState) => state.auth);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCredentials({
      ...credentials,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    dispatch(loginUser(credentials));
  };

  // Si el usuario está autenticado, redirigir
  if (isAuthenticated) {
    return <Navigate to="/" />;
  }

  return (
    <AppProvider branding={{ logo: <img src={logo} alt="MUI logo" style={{ height: 24 }} />, title: 'MUI' }}>
      <Container maxWidth="xs">
        <Typography variant="h5" align="center" gutterBottom>
          Iniciar Sesión
        </Typography>

        <form onSubmit={handleSubmit}>
          <TextField
            label="Email"
            name="email"
            type="email"
            value={credentials.email}
            onChange={handleChange}
            fullWidth
            margin="normal"
            variant="outlined"
            required
          />
          <TextField
            label="Contraseña"
            name="contrasena"
            type="password"
            value={credentials.contrasena}
            onChange={handleChange}
            fullWidth
            margin="normal"
            variant="outlined"
            required
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            disabled={loading}
            style={{ marginTop: '16px' }}
          >
            {loading ? <CircularProgress size={24} color="inherit" /> : 'Iniciar Sesión'}
          </Button>
        </form>

        {error && <Typography color="error" align="center" style={{ marginTop: '8px' }}>{error}</Typography>}
      </Container>
    </AppProvider>
  );
};

export default Login;
