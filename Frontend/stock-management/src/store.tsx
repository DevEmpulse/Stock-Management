import { configureStore } from '@reduxjs/toolkit';
import authReducer from './features/auth/authSlice';

export type RootState = {
    auth: ReturnType<typeof authReducer>
}

export type AppDispatch = typeof store.dispatch

const store = configureStore({
    reducer:{
        auth:authReducer,
    },
})

export default store 