import { useState, useEffect } from "react";
import { getRooms } from "../services/roomService";
import { getNextRaceByCourse } from "../services/raceService";

const useRooms = () => {
  const [rooms, setRooms] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [raceSchedules, setRaceSchedules] = useState({});

  useEffect(() => {
    const fetchRoomsAndRaceSchedules = async () => {
      try {
        const roomData = await getRooms();
        setRooms(roomData);

        const schedules = {};
        // 各部屋に対応する次のレース情報を取得
        for (const room of roomData) {
          const raceData = await getNextRaceByCourse(room.Name);  // 部屋名を基に次のレース情報を取得
          if (raceData) {
            schedules[room.Name] = raceData;
          }
        }
        setRaceSchedules(schedules);  // レース情報を保存
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchRoomsAndRaceSchedules();
  }, []);

  return { rooms, loading, error , raceSchedules};
};

export default useRooms;