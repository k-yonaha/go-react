const API_URL = "http://localhost:8080/api/rooms";

// 部屋一覧を取得する関数
export const getRooms = async () => {
  try {
    const response = await fetch(API_URL);
    if (!response.ok) {
      throw new Error("部屋情報の取得に失敗しました");
    }
    return await response.json();
  } catch (error) {
    throw new Error(error.message);
  }
};