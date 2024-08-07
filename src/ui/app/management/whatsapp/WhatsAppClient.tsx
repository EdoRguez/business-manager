'use client';

import useWebSocket from 'react-use-websocket';
import { useState } from "react";
import { WebsocketMessage, WebsocketMessageTypes } from '@/app/types/websocket';
import SimpleCard from '@/app/components/cards/SimpleCard';
import Connected from '@/app/components/whatsapp/Connected';
import ConnectQr from '@/app/components/whatsapp/ConnectQr';

const WS_URL = 'ws://localhost:50055/ws';

const WhatsAppClient = () => {
  const [showQR, setShowQR] = useState<boolean>(true);
  const [qrString, setQrString] = useState<string>('');
  const [showAlreadyConnected, setShowAlreadyConnected] = useState<boolean>(false);
  const [messageData, setMessageData] = useState<string>('');
  const { sendMessage, lastMessage, readyState } = useWebSocket(WS_URL, {
    share: true,
    shouldReconnect: () => false,
    onOpen: () => {
      console.log('WebSocket connection established.');
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      console.log('-------------------')
      const data: WebsocketMessage = JSON.parse(event?.data);
      console.log(data)
      if(messageData !== data?.message)
        setMessageData(data.message);

      if(data.messageType === WebsocketMessageTypes.QR_CODE) {
        setQrString(data.message);
        setShowQR(true);
      }

      if(data.messageType === WebsocketMessageTypes.ALREADY_CONNECTED_CODE) {
        setMessageData(data.message);
        setShowQR(false);
        setShowAlreadyConnected(true);
      }

      // if(data.messageType === WebsocketMessageTypes.CONVERSATIONS_CODE) {
      //   const result: WhatsappConversation[] = JSON.parse(data.message);
      //   setConversations(result);
      //   setShowQR(false);
      // }
    }
  });

  return (
    <SimpleCard>
      {
        showQR && (
          <ConnectQr qrString={qrString} />
        )
      }

      {
        showAlreadyConnected && (
          <Connected phone={messageData}/>
        )
      }

      {/*
        !showQR &&
        conversations && 
        <ChatView conversations={conversations} />
        */
      }
    </SimpleCard>
  )
}

export default WhatsAppClient;
