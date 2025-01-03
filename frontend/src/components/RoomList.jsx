import { Link } from "react-router-dom";
import useRooms from "../hooks/useRooms";

const RoomList = () => {
  const { rooms, loading, error } = useRooms();

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  const getTimeOfDay = () => {
    const hour = new Date().getHours();
    if (hour >= 6 && hour < 12) {
      return "morning";
    } else if (hour >= 12 && hour < 18) {
      return "afternoon";
    } else {
      return "night";
    }
  };

  const timeOfDay = getTimeOfDay();
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
    <div className="grid grid-cols-4 grid-rows-4 gap-2">
      {rooms.map((room) => (
        <Link
          key={room.ID}
          to={`/room/${room.ID}`}
          className="border-2 border-black box-border rounded-lg flex flex-col items-center justify-between h-24 w-40 sm:w-40 md:w-46 lg:w-56 max-w-full p-2"
        >
          <div className={`text-lg font-bold text-black w-full ${bgColorClass} rounded-t-lg`}>
            {room.Name}
          </div>
          
          <div className="text-sm text-black w-full border-t border-b border-black bg-white">
            何日目
          </div>
          
          <div className="text-sm text-black w-full p-1 bg-white rounded-b-lg">
            開始時間
          </div>
        </Link>
      ))}
    </div>
  );
};

export default RoomList;
