"use client";

import React, { useCallback, useEffect, useState } from "react";
import {
  Box,
  Button,
  Container,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Stack,
  Switch,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
  VStack,
  Avatar,
  useColorModeValue,
  useToast,
} from "@chakra-ui/react";
import SimpleCard from "@/app/components/cards/SimpleCard";
import ImagesUpload from "@/app/components/uploads/ImagesUpload";
import { Company, EditCompany } from "@/app/types/company";
import { EditCompanyRequest, GetCompanyRequest } from "@/app/api/companies/route";
import { CurrentUser } from "@/app/types/auth";
import getCurrentUser from "@/app/actions/getCurrentUser";
import useLoading from "@/app/hooks/useLoading";
import { EditUser, User } from "@/app/types/user";
import { GetUserRequest } from "@/app/api/users/route";

const AccountClient = () => {

  const isLoading = useLoading();
  const toast = useToast();

  const [companyFormData, setCompanyFormData] = useState<EditCompany>({
    id: 0,
    name: '',
    imageUrl: ''
  });
  const [userFormData, setUserFormData] = useState<EditUser>({
    id: 0,
    roleId: 0,
    email: '',
    password: ''
  });

  const handleCompanyFormChange = (event: any) => {
    const { name, value } = event.target;
    setCompanyFormData((prev) => ({...prev, [name]: value}));
  }
  
  const handleUserFormChange = (event: any) => {
    const { name, value } = event.target;
    setUserFormData((prev) => ({...prev, [name]: value}));
  }

  const getCompany = useCallback(async () => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if(currentUser) {
      isLoading.onStartLoading();
      let company: Company = await GetCompanyRequest({
        id: currentUser?.companyId,
      });
      if (company) {
        setCompanyFormData({
          id: company.id ?? 0,
          name: company.name ?? "",
          imageUrl: company.imageUrl ?? ""
        });
      }

      isLoading.onEndLoading();
    }
  }, [])
  
  const getUser = useCallback(async () => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if(currentUser) {
      isLoading.onStartLoading();
      let user: User = await GetUserRequest({
        id: currentUser?.id,
      });
      if (user) {
        setUserFormData({
          id: user.id ?? 0,
          roleId: user.roleId ?? 0,
          email: user.email ?? '',
          password: ''
        });
      }

      isLoading.onEndLoading();
    }
  }, [])

  useEffect(() => {
    getCompany();
    getUser();
  }, [getCompany])

  const isCompanyFormValid = (): boolean => {
    if (!companyFormData.name) return false;

    return true;
  };

  const onCompanySubmit = async () => {
    if (isCompanyFormValid()) {
      isLoading.onStartLoading();
      let editCompany: any = await EditCompanyRequest(companyFormData);
      if (editCompany?.error) {
        isLoading.onEndLoading();
        showErrorMessage(editCompany.error);
      } else {
        isLoading.onEndLoading();
        showSuccessCreationMessage("Empresa modificada exitosamente");
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  }

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: "Empresa",
      description: msg,
      variant: "customsuccess",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  };

  const showErrorMessage = (msg: string) => {
    toast({
      title: "Error",
      description: msg,
      variant: "customerror",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  };

  return (
    <Container maxW="4xl" py={8}>
      <SimpleCard>
        <div className="p-4">
          <Flex align="center" gap={4}>
            <Avatar
              size="xl"
              src="/placeholder.svg?height=80&width=80"
            />
            <Box>
              <Heading size="lg" className="break-all">
                { companyFormData.name }
              </Heading>
            </Box>
          </Flex>
          <Tabs isFitted variant="enclosed" className="mt-8">
            <TabList mb="1em">
              <Tab>Empresa</Tab>
              <Tab>Cuenta</Tab>
            </TabList>
            <TabPanels>
              <TabPanel>
                <VStack spacing={4} align="stretch">
                  <Heading size="md" mb={2}>
                    Información de la Empresa
                  </Heading>
                  <Text size="sm" color="gray.500">
                    Actualiza los datos de tu empresa aquí
                  </Text>
                  <FormControl>
                    <label className="text-sm">Nombre</label>
                    <Input
                      size="sm"
                      name="name"
                      maxLength={50}
                      value={companyFormData.name}
                      onChange={handleCompanyFormChange}
                    />
                  </FormControl>
                  <FormControl>
                    <label className="text-sm">Imagen</label>
                    <div className="border rounded py-5 px-3">
                        <ImagesUpload maxImagesNumber={1} showAddImage={true} />
                    </div>
                  </FormControl>
                  <Button variant="main" alignSelf="flex-start" className="mt-4" onClick={onCompanySubmit}>
                    Guardar Cambios
                  </Button>
                </VStack>
              </TabPanel>
              <TabPanel>
                <VStack spacing={4} align="stretch">
                  <Heading size="md" mb={2}>
                    Ajuste de Cuenta
                  </Heading>
                  <Text size="sm" color="gray.500">Actualiza el correo de tu usuario</Text>
                  <FormControl>
                    <label className="text-sm">Correo</label>
                    <Input size="sm" name="email" value={userFormData.email} onChange={handleUserFormChange} />
                  </FormControl>
                  <Button variant="main" alignSelf="flex-start" className="mt-4">
                    Actualizar Correo
                  </Button>
                  <hr className="my-4" />
                  <Text size="sm" color="gray.500">Cambiar tu contraseña</Text>
                  <FormControl>
                    <label className="text-sm">Contraseña Nueva</label>
                    <Input size="sm" type="password" />
                  </FormControl>
                  <FormControl>
                    <label className="text-sm">Repetir Contraseña Nueva</label>
                    <Input size="sm" type="password" />
                  </FormControl>
                  <Button variant="main" alignSelf="flex-start" className="mt-4">
                    Cambiar Contraseña
                  </Button>
                </VStack>
              </TabPanel>
            </TabPanels>
          </Tabs>
        </div>
      </SimpleCard>
    </Container>
  );
};

export default AccountClient;
