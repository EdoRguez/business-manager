'use server'

import { Login } from "@/app/types/auth"
import { cookies } from "next/headers";

export async function login(
  request: Login,
) {
  try {
    const res = await fetch('http://localhost:3001/api/auth/login', {
      method: 'POST',
      body: JSON.stringify(request)
    });
    let response = await res.json();

    if(!response.error) {
      cookies().set('token', `Bearer ${response.token}`);
    }

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}
