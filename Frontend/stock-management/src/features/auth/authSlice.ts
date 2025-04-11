import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'
import axios, { AxiosError } from 'axios';

interface AuthState {
    user: { id: number, email: string, nombre: string } | null;
    token: string | null
    loading: boolean
    error: string | null
    isAuthenticated: boolean;
}

const initialState: AuthState = {
    user: null,
    token: localStorage.getItem('token') || null,
    loading: false,
    error: null,
    isAuthenticated: !!localStorage.getItem('token'), // Verdadero si hay token
  };
interface AuthError {
    message: string;
}

export const loginUser = createAsyncThunk<
    { user: AuthState['user']; token: string },
    { email: string; contrasena: string },
    { rejectValue: string }
>(
    'auth/loginUser',
    async (credentials, { rejectWithValue }) => {
        try {
            const response = await axios.post('http://localhost:3000/api/users/login', credentials);
            return response.data
        } catch (error) {
            // Tipamos el error como AxiosError con un tipo genérico para el cuerpo del error
            const axiosError = error as AxiosError<AuthError>;
            return rejectWithValue(
                axiosError.response?.data?.message || 'Error al iniciar sesión'
            );
        }
    }
);

export const checkAuth = createAsyncThunk<
  { user: AuthState['user']; token: string },
  void,
  { rejectValue: string }
>('auth/checkAuth', async (_, { rejectWithValue }) => {
  const token = localStorage.getItem('token');
  if (!token) {
    return rejectWithValue('No hay token');
  }

  try {
    const response = await axios.get('http://localhost:3000/api/users/getUser', {
      headers: {
        Authorization:`Bearer ${token}`,
      },
    });
    return { user: response.data, token };
  } catch (error) {
    localStorage.removeItem('token'); // Elimina el token si no es válido
    const axiosError = error as AxiosError<AuthError>;
    return rejectWithValue(
      axiosError.response?.data?.message || 'Sesión no válida'
    );
  }
});


const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        logout: (state) => {
            state.user = null
            state.token = null
            state.isAuthenticated = false
            localStorage.removeItem('token')
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(loginUser.pending, (state) => {
                state.loading = true
                state.error = null
            })
            .addCase(loginUser.fulfilled, (state, action: PayloadAction<{ user: AuthState['user']; token: string }>) => {
                state.loading = false
                state.user = action.payload.user
                state.token = action.payload.token
                state.isAuthenticated = true
                localStorage.setItem('token', action.payload.token)
            })
            .addCase(loginUser.rejected, (state, action: PayloadAction<string | undefined>) => {
                state.loading = false;
                state.error = action.payload || 'Error desconocido'
            })
            builder
            .addCase(checkAuth.pending, (state) => {
              state.loading = true;
              state.error = null;
            })
            .addCase(checkAuth.fulfilled, (state, action: PayloadAction<{ user: AuthState['user']; token: string }>) => {
              state.loading = false;
              state.user = action.payload.user;
              state.token = action.payload.token;
              state.isAuthenticated = true;
            })
            .addCase(checkAuth.rejected, (state, action: PayloadAction<string | undefined>) => {
              state.loading = false;
              state.user = null;
              state.token = null;
              state.isAuthenticated = false;
              state.error = action.payload || 'Error al verificar autenticación';
            });
    },
})

export const { logout } = authSlice.actions
export default authSlice.reducer