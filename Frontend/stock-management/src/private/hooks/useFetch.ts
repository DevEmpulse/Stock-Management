import { useEffect, useState } from "react";

type Data<T> = T | null
type ErrorType= Error | null

interface Params<T>{
    data: Data<T>;
    loading: boolean;
    error: ErrorType
}

export const useFetch= <T>(url: string): Params<T> =>{
    const [data, setData] = useState<Data<T>>(null)
    const[loading, setLoading]= useState(true)
    const [error, setError]= useState<ErrorType>(null)


    useEffect(() =>{
        const controller= new AbortController()
        setLoading(true)
        setError(null);

        const fetchData = async ()=>{
            try {
                const response = await fetch(url, { signal: controller.signal })
                if (!response.ok){
                    throw new Error(`Error en la peticion : ${response.status}`)
                }
                const jsonData: T = await response.json()
                setData(jsonData)
            } catch (err){
                if (err instanceof DOMException && err.name === "AbortError") {
                    console.log("Solicitud abortada por el usuario");
                    return; // No establece error en caso de aborto
                }
                setError(err as Error)
            } finally {
                setLoading(false)
            }
        }
        fetchData()
        return () =>{
            controller.abort()
        }
    },[url])

    return {data, loading, error}
}