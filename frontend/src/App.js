import { useState } from 'react';
import './App.css';
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';

function App() {
  const [orderId, setOrderId] = useState('')
  const [orderInfo, setOrderInfo] = useState('')
  const [errorInfo, setErrorInfo] = useState('')

  function handleOrderIdField(e) {
    setOrderId(e.target.value)
  }

  function handleSending() {
    setOrderInfo('')
    setErrorInfo('')
    if (orderId !== '') {
      fetch('http://localhost:8000/getOrder/' + orderId)
      .then((response) => response.json())
      .then((data) => {
        if (data.message === 'OK') {
          var stringifiedOrder = JSON.stringify(data.order, null, 2)
          setOrderInfo(stringifiedOrder)
        } else if (data.message === 'ORDERS_NOT_FOUND') {
          setErrorInfo('Заказы по данному order_uid не найдены')
        } else {
          console.log(data.message)
        }})
      .catch((error) => {
        console.log(error)
      })
    } else {
      setErrorInfo('Запрос пуст')
    }
  }

  return (
    <Box className='app'>
      <Typography variant="h4">Введите order_uid:</Typography>
      <Box display='flex' justifyContent='center'>
        <TextField variant="outlined" value={orderId} onChange={handleOrderIdField}/>
        <Button variant="outlined" onClick={handleSending}>Получить</Button>
      </Box>
      {errorInfo && (
        <Typography variant="h4">{errorInfo}</Typography>
      )}
      {orderInfo && (
        <Box display='inline-block' textAlign='left'>
          <pre>{orderInfo}</pre>
        </Box>
      )}
    </Box>
  );
}

export default App;
