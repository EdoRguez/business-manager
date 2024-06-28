'use client';

import { Tab, TabList, TabPanel, TabPanels, Tabs } from "@chakra-ui/react";
import SimpleCard from "../components/cards/SimpleCard";
import { Icon } from "@iconify/react";
import ChatList from "../components/whatsapp/ChatList";
import MessageList from "../components/whatsapp/MessageList";
import MessageBar from "../components/whatsapp/MessageBar";
import UserBar from "../components/whatsapp/UserBar";
import useWebSocket from 'react-use-websocket';
import { useState } from "react";
import ConnectQr from "../components/whatsapp/ConnectQr";
import { WebsocketMessage, WebsocketMessageTypes } from "../types/websocket";
import { WhatsappConversation } from "../types/whatsapp";

const WS_URL = 'ws://localhost:50055/ws';

const WhatsAppClient = () => {
  const [showQR, setShowQR] = useState<boolean>(true);
  const [qrString, setQrString] = useState<string>('');
  const [messageData, setMessageData] = useState<string>('');
  const [conversations, setConversations] = useState<WhatsappConversation[]>();
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

      if(data.messageType === WebsocketMessageTypes.CONVERSATIONS_CODE) {
        const result: WhatsappConversation[] = JSON.parse(data.message);
        setConversations(result);
        setShowQR(false);
      }
    }
  });

  const sendWhatsappMessage = (message: string): void => {
    console.log('Enviar primero el mensaje');
    console.log(lastMessage);

    const result: any = {
      type: 'send_message',
      payload: {
        message: message,
        from: 'Eduardo'
      }
    };
    sendMessage(JSON.stringify(result));
  }

  // const handleClickSendMessage = useCallback(() => sendMessage('Hello'), []);

  return (
    <SimpleCard>
      {showQR && (
        <ConnectQr qrString={qrString} />
      )}

      {!showQR &&
       conversations &&(
        <div className="h-[85vh] flex my-3 rounded border-2 border-slate-200">
          <div className="w-3/6 border-r-2 border-r-slate-200 overflow-scroll">
            <Tabs colorScheme='green'>
              <TabList>
                <Tab><Icon icon="ic:baseline-message" /><span className="text-sm"> Mensajes</span></Tab>
                <Tab><Icon icon="fluent:person-support-16-filled" /><span className="text-sm flex items-center"> Atendiendo</span></Tab>
                <Tab><Icon icon="lets-icons:done-duotone" /><span className="text-sm flex items-center"> Finalizado</span></Tab>
              </TabList>

              <TabPanels>
                <TabPanel>
                  <ChatList conversations={conversations} />
                </TabPanel>
                <TabPanel>
                  <ChatList conversations={conversations} />
                </TabPanel>
                <TabPanel>
                  <ChatList conversations={conversations} />
                </TabPanel>
              </TabPanels>
            </Tabs>
          </div>
          <div className="w-full" style={{ backgroundImage: 'url(/images/whatsapp/ws_bg.jpg)', backgroundSize: 'contain', backgroundRepeat: 'repeat' }}>
            <div className="h-full flex flex-col">
              <UserBar />
              <div className="mt-auto overflow-y-auto">
                <MessageList messages={conversations[2].messages} />
              </div>
              <MessageBar onSendMessage={sendWhatsappMessage} />
            </div>

          </div>
        </div>
      )}
    </SimpleCard>
  )
}

export default WhatsAppClient;
