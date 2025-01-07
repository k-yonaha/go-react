import { Link } from "react-router-dom";
import useRooms from "../hooks/useRooms";

const RoomList = () => {
  const { rooms, loading, error, raceSchedules } = useRooms();

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
 
  const getTimeOfDay = (raceTime) => {
    const time = new Date(raceTime);
    const options = { timeZone: "Asia/Tokyo", hour: "2-digit", minute: "2-digit" };
    const localTime = time.toLocaleString("ja-JP", options);
    const hour = parseInt(localTime.split(":")[0]);

    if (hour >= 6 && hour < 9) {
      return "morning";
    } else if (hour >= 9 && hour < 15) {
      return "afternoon";
    } else if (hour >= 15) {
      return "night";
    } else {
      return "";
    }
  };

  return (
    <div className="grid grid-cols-2 sm:grid-cols-4 lg:grid-cols-4 grid-rows-4 gap-2">
      {rooms.map((room) => {
        // map内で変数の定義
        const raceSchedule = raceSchedules[room.Name] ? raceSchedules[room.Name][0] : null;

        const raceTime = raceSchedule ? raceSchedule.RaceTime : 0

        const timeOfDay = getTimeOfDay(raceTime);
        let bgColorClass;

        switch (timeOfDay) {
          case "morning":
            bgColorClass = "bg-yellow-200";
            break;
          case "afternoon":
            bgColorClass = "bg-blue-200";
            break;
          case "night":
            bgColorClass = "bg-purple-500";
            break;
          default:
            bgColorClass = "bg-gray-300";
            break;
        }

        return (
          <Link
            key={room.ID}
            to={`/room/${room.ID}`}
            className={`border-2 border-black box-border rounded-lg flex flex-col items-center justify-between h-24 w-40 sm:w-40 md:w-46 lg:w-56 max-w-full p-2 ${
              !raceSchedule ? "disabled opacity-50 cursor-not-allowed pointer-events-none" : ""
            } ${bgColorClass}`}
          >
            <div
              className={`text-lg font-bold text-black w-full ${bgColorClass} rounded-t-lg`}
            >
              {room.Name}
            </div>

            <div className="text-sm text-black w-full border-t border-b border-black bg-white">
            <span className="mr-2">{raceSchedule ? raceSchedule.RaceType : "--"}</span>
              {raceSchedule ? raceSchedule.RaceDay : "--"}
            </div>

            <div className="text-sm text-black w-full p-1 bg-white rounded-b-lg">
              <span className="mr-2">
                {raceSchedule ? raceSchedule.RaceNumber + "R" : "--"}
              </span>
              {raceSchedule
                ? new Date(raceSchedule.RaceTime).toLocaleTimeString()
                : "-- --"}
            </div>
          </Link>
        );
      })}
    </div>
  );
};


export default RoomList;
