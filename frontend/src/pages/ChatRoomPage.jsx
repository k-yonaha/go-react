import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import useWebSocket from "../hooks/useWebSocket";

import "../assets/styles/ChatRoomPage.css"; 

const ChatRoomPage = () => {
  const { roomId } = useParams();
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState("");

  const { connect, sendMessage, receiveMessage, disconnect,connected } = useWebSocket(roomId);
  useEffect(() => {
    if (!connected) {
      console.log("接続中...");
      connect();
    }
  
    return () => {
      disconnect();
    };
  }, [connect,connected,disconnect]);

  useEffect(() => {
    receiveMessage((msg) => {
      setMessages((prevMessages) => [...prevMessages, msg]);
    });
  }, [receiveMessage]);

  const handleSendMessage = () => {
    console.log(22222)
    if (newMessage.trim()) {
      sendMessage(newMessage);
      setNewMessage("");
    }
  };

  return (
    <div>
      <h1>Room {roomId} Chat</h1>
      <div className="messages">
        {messages.map((msg, idx) => (
          <div key={idx} className="message">
            {msg}
          </div>
        ))}
      </div>
      <input
        type="text"
        value={newMessage}
        onChange={(e) => setNewMessage(e.target.value)}
        placeholder="Type a message"
      />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
};

export default ChatRoomPage;