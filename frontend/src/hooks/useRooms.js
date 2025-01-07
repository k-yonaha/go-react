import { useState, useEffect } from "react";
import { getRooms } from "../services/roomService";
import { getRaceSchedulesByDate } from "../services/raceService";

const useRooms = () => {
  const [rooms, setRooms] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [raceSchedules, setRaceSchedules] = useState({});

  useEffect(() => {
    const fetchRooms = async () => {
      try {
        const roomsData = await getRooms();
        setRooms(roomsData);
      } catch (err) {
        setError(err.message);
      }
    };
    const fetchRaceSchedules = async () => {
      try {
        const raceSchedulesData = await getRaceSchedulesByDate();
        setRaceSchedules(raceSchedulesData);
      } catch (err) {
        console.log(err)
        setError(err.message);
      }
    };

    fetchRooms();
    fetchRaceSchedules();
  }, []);
  useEffect(() => {
    if (rooms.length > 0) {
      setLoading(false);
    }
  }, [rooms]);
  return { rooms, loading, error, raceSchedules };
};

export default useRooms;
