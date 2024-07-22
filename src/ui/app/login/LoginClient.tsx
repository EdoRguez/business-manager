'use client';

import Image from 'next/image'
import SimpleCard from '../components/cards/SimpleCard';
import { Button, Container, Input } from '@chakra-ui/react';
import { useState } from 'react';
import { Login } from '../types/auth';
import { login } from '../api/auth/route';

const LoginClient = () => {
  const [formData, setFormData] = useState<Login>({ email: '', password: '' });

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  }

  const onLogin = async () => {
    console.log('epa')
    console.log(formData)
    let result: any = await login(formData);
    console.log('vemoas')
    console.log(result)
  }

  return (
    <>
      <div className="grid grid-cols-8 h-screen">
        <div className="hidden md:flex col-span-3 bg-gradient-to-b from-thirdcolor to-fourthcolor flex-col items-center justify-center">
          <h1 className='mb-10 text-white font-medium text-xl text-center'>¡Ingresa a tu cuenta y empieza a administrar tu negocio!</h1>
          <Image className='rounded-lg' src='/images/login/main_image.png' alt="Logo" width={300} height={300} />
        </div>
        <div className="col-span-8 md:col-span-5 bg-graybackground flex items-center">
          <Container>
            <div className='relative'>
              <div className='absolute rounded-full border-[7px] border-white top-[-60px] left-[40%]'>
                <Image className='rounded-full border-rose-500 ' src='/images/logo.png' width={100} height={100} alt='logo' />
              </div>
              <SimpleCard>
                <div className='px-1 py-8'>
                  <div className='mt-2'>
                    <label className='text-sm'>Correo</label>
                    <Input size="sm" type='email' name='email' value={formData.email} onChange={handleChange} />
                  </div>
                  <div className='mt-2'>
                    <label className='text-sm'>Contraseña</label>
                    <Input size="sm" type='password' name='password' value={formData.password} onChange={handleChange} />
                  </div>
                  <div className="mt-3">
                    <Button variant="main" className='w-full' onClick={onLogin}>
                      Iniciar Sesión
                    </Button>
                  </div>
                </div>
              </SimpleCard>
            </div>
          </Container>
        </div>
      </div>
    </>
  )
}

export default LoginClient;
