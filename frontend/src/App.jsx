import { BrowserRouter, Route, Routes } from "react-router-dom";
import RoomListPage from "./pages/RoomListPage";
import ChatRoomPage from "./pages/ChatRoomPage";

import "./App.css";

const App = () => {
  return (
    <BrowserRouter
      future={{
        v7_startTransition: true,
        v7_relativeSplatPath: true,
      }}
    >
      <Routes>
        {/* 部屋一覧ページ */}
        <Route path="/" element={<RoomListPage />} />
        <Route path="/room/:roomId" element={<ChatRoomPage />} />
        {/* チャットルームページ */}
        {/* <Route path="/room/:roomId" element={<ChatRoomPage />} />  */}
      </Routes>
    </BrowserRouter>
  );
};

export default App;

