"use client";

import SimpleCard from "@/app/components/cards/SimpleCard";
import { validNumbers } from "@/app/utils/InputUtils";
import { Button, Input } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import { useState } from "react";

// const WS_URL = 'ws://localhost:50055/ws';

// const WhatsAppClient = () => {
//   const [showQR, setShowQR] = useState<boolean>(true);
//   const [qrString, setQrString] = useState<string>('');
//   const [showAlreadyConnected, setShowAlreadyConnected] = useState<boolean>(false);
//   const [messageData, setMessageData] = useState<string>('');
//   const { sendMessage, lastMessage, readyState } = useWebSocket(WS_URL, {
//     share: true,
//     shouldReconnect: () => false,
//     onOpen: () => {
//       console.log('WebSocket connection established.');
//     },
//     onMessage: (event: WebSocketEventMap['message']) => {
//       console.log('-------------------')
//       const data: WebsocketMessage = JSON.parse(event?.data);
//       console.log(data)
//       if(messageData !== data?.message)
//         setMessageData(data.message);

//       if(data.messageType === WebsocketMessageTypes.QR_CODE) {
//         setQrString(data.message);
//         setShowQR(true);
//       }

//       if(data.messageType === WebsocketMessageTypes.ALREADY_CONNECTED_CODE) {
//         setMessageData(data.message);
//         setShowQR(false);
//         setShowAlreadyConnected(true);
//       }

//       // if(data.messageType === WebsocketMessageTypes.CONVERSATIONS_CODE) {
//       //   const result: WhatsappConversation[] = JSON.parse(data.message);
//       //   setConversations(result);
//       //   setShowQR(false);
//       // }
//     }
//   });

//   return (
//     <SimpleCard>
//       {
//         showQR && (
//           <ConnectQr qrString={qrString} />
//         )
//       }

//       {
//         showAlreadyConnected && (
//           <Connected phone={messageData}/>
//         )
//       }

//       {/*
//         !showQR &&
//         conversations &&
//         <ChatView conversations={conversations} />
//         */
//       }
//     </SimpleCard>
//   )
// }

const WhatsAppClient = () => {
  const [phone, setPhone] = useState<string>("");
  const [isEditMode, setIsEditMode] = useState<boolean>(false);

  const handleNumberChange = (event: any) => {
    const { value } = event.target;
    if (value && !validNumbers(value)) return;
    setPhone(value);
  };

  const handleEditMode = () => {
    setIsEditMode((prev) => !prev);
  };
  return (
    <SimpleCard>
      <h1 className="font-bold">Whatsapp Business</h1>
      <div className="grid grid-cols-1 md:grid-cols-5 gap-2 ">
        <div className="md:col-span-4 mt-2">
          <label className="text-sm">Teléfono</label>
          <Input
            size="sm"
            name="phone"
            placeholder="04141234567"
            value={phone}
            onChange={handleNumberChange}
            maxLength={11}
          />
        </div>
        <div className="flex items-end">
          {isEditMode ?
             (
              <Button
                size="sm"
                variant="main"
                className="mx-1"
                onClick={handleEditMode}
              >
                <Icon icon="lucide:edit" />
              </Button>
            ) : (
            <Button
              size="sm"
              variant="main"
              className="mx-1"
              onClick={handleEditMode}
            >
              <Icon icon="lucide:edit" />
            </Button>
          )}
        </div>
      </div>
      <p className="my-2 text-xs text-gray-600">
        Este número será usado para notificarte cuando tengas una nueva order y
        para que tus clientes se comuniquen contigo
      </p>
    </SimpleCard>
  );
};

export default WhatsAppClient;
