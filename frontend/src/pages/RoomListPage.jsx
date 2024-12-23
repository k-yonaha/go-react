import { Link } from 'react-router-dom';

const RoomListPage = () => {
  const rooms = [
    { id: 1, name: 'Room 1' },
    { id: 2, name: 'Room 2' },
    { id: 3, name: 'Room 3' },
  ];

  return (
    <div>
      <h1>Room List</h1>
      <ul className="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-4 gap-4">
        {rooms.map((room) => (

          <li key={room.id} className="p-4 rounded-lg bg-indigo-300 dark:bg-indigo-800 dark:text-indigo-400">
            <Link to={`/room/${room.id}`}>{room.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default RoomListPage;