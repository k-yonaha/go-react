import { useEffect, useState } from 'react';

import './App.css';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    // GoのAPIを呼び出してメッセージを取得
    fetch('http://localhost:8080/api/hello')
      .then(response => response.json())
      .then(data => setMessage(data.message))
      .catch(err => console.error('Error:', err));
  }, []);

  return (
    <div className="App border border-gray-400 rounded-2xl">
      <h1>Message from Go Backend:</h1>
      <p>{message}</p>
    </div>
  );
}

export default App;