import React, { useState } from 'react';

function App() {
  const [items, setNumber] = useState('');
  const [result, setResult] = useState('');

  const [packSizes, setText] = useState('250, 500, 1000, 2000, 5000');
  const [responseMessage, setResponseMessage] = useState('');

  const handlePackSizes = (e) => {
    setText(e.target.value);
  };

  const handleItemsChange = (e) => {
    setNumber(e.target.value);
  };

  const handleGetPackConfig = async () => {
    try {
      const response = await fetch(`http://localhost:8881/packs?items=${items}`);
      const textResult = await response.text();
      setResult(textResult); 
    } catch (error) {
      console.error('Error fetching data: ', error);
      setResult('Error fetching data');
    }
  };

  const handleSubmitPackSizes = async () => {
    setResponseMessage("");
    try {
      const requestOptions = {
        method: 'PUT',
        headers: { 'Content-Type': 'text/plain' },
        body: packSizes,
      };

      const response = await fetch(`http://localhost:8881/packs`, requestOptions);
      const result = await response.text();
      setResponseMessage(result);
    } catch (error) {
      console.error('Error:', error);
      setResponseMessage("error");
    }
  };

  return (
    <div>
      <input
        type="text"
        value={packSizes}
        onChange={handlePackSizes}
        placeholder="Update pack sizes (comma separated values)"
      />
      <button onClick={handleSubmitPackSizes}>Enter pack sizes (comma separated values)</button>
      <div>Response: {responseMessage}</div>

      <input
        type="text"
        value={items}
        onChange={handleItemsChange}
        placeholder="Enter a number"
      />
      <button onClick={handleGetPackConfig}>Get pack config</button>
      <div>Result: {result}</div>
    </div>
  );
}

export default App;