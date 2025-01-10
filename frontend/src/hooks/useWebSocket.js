import { useState } from "react";

const useWebSocket = (roomId) => {
  const [ws, setWs] = useState(null);
  const [connected, setConnected] = useState(false); // 接続状態を管理

  const connect = () => {
    if (ws) {
      console.log("既に接続済みです。");
      return; // 既に接続されている場合は何もしない
    }
    // WebSocketのインスタンスを作成
    const socket = new WebSocket(`ws://localhost:8080/ws/${roomId}`);

    // WebSocket接続時
    socket.onopen = () => {
      console.log("接続しました。");
      setConnected(true);
    };

    // メッセージ受信時
    socket.onmessage = (event) => {
      const message = event.data;
      console.log("メッセージ受信:", message);
    };

    // エラー発生時
    socket.onerror = (error) => {
      console.log("接続エラー:", error);
      // エラー時に再接続を試みる
    };

    // WebSocketが閉じられた時
    socket.onclose = (event) => {
      console.log("接続を閉じました", event);
      setConnected(false); // 接続状態を更新
    };

    setWs(socket); // WebSocketインスタンスを状態に保存
  };

  const sendMessage = (message) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(message); // メッセージを送信
    } else {
      console.log("接続ができません。メッセージが送信できませんでした");
    }
  };

  const receiveMessage = (callback) => {
    if (ws) {
      ws.onmessage = (event) => {
        callback(event.data);
      };
    }
  };

  const disconnect = () => {
    if (ws) {
      //ws.close();
    }
  };

  return { connect, sendMessage, receiveMessage, disconnect, connected };
};

export default useWebSocket;
