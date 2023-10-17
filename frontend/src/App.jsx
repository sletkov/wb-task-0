
import {useState, useEffect } from 'react';
import axios from 'axios'
import './App.css';
function App() {
  const [id, setId] = useState('')
  const [order, setOrder] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      var response = await axios.get(`http://localhost:8080/orders/${id}`, {
        validateStatus: function(status) {
          return true
        }
      });
      console.log(response.data)
      setOrder(response.data);
    } catch (error) {
      console.error(error);
    }
  }



  return (
    <div className="App">
      <h1>Get info about order</h1>
      <form onSubmit={handleSubmit}>
        <input value={id} type="text"  placeholder='Enter the order id' onChange={(e) => setId(e.target.value)}/>
        <button type='sumbit'>Get</button>
      </form>
      {order !== null ? 
      
        <div className='order'>
          <h2>Order uid: {order.order_uid}</h2>
          <p>Track number: {order.track_number}</p>
          <p>Entry: {order.entry}</p>
          <div className='delivery'>
            <h3>Delivery</h3>
            <p>name: {order.delivery.name}</p>
            <p>phone: {order.delivery.phone}</p>
            <p>zip: {order.delivery.zip}</p>
            <p>city: {order.delivery.city}</p>
            <p>address: {order.delivery.address}</p>
            <p>region: {order.delivery.region}</p>
            <p>email: {order.delivery.email}</p>
          </div>

          <div className='items'>
            <h3>Items: </h3>
            {order.items.map((item ,idx) => (
              <div className='item'>
                <h4>Item № {idx + 1}</h4>
                <p>chrt_id: {item.chrt_id}</p>
                <p>track_number: {item.track_number}</p>
                <p>price: {item.price}</p>
                <p>rid: {item.rid}</p>
                <p>name: {item.name}</p>
                <p>sale: {item.sale}</p>
                <p>size: {item.size}</p>
                <p>total_price: {item.total_price}</p>
                <p>nm_id: {item.nm_id}</p>
                <p>brand: {item.brand}</p>
                <p>status: {item.status}</p>
              </div>
            ))}
           
          </div>

          <div className='payment'>
            <h3>Payment</h3>
            <p>transaction: {order.payment.transaction}</p>
            <p>request_id: {order.payment.request_id}</p>
            <p>currency: {order.payment.currency}</p>
            <p>provider: {order.payment.provider}</p>
            <p>amount: {order.payment.amount}</p>
            <p>payment_dt: {order.payment.payment_dt}</p>
            <p>bank: {order.payment.bank}</p>
            <p>delivery_cost: {order.payment.delivery_cost}</p>
            <p>goods_total: {order.payment.goods_total}</p>
            <p>custom_fee: {order.payment.custom_fee}</p>
          </div>
          <p>locale: {order.locale}</p>
          <p>internal_signature: {order.internal_signature}</p>
          <p>customer_id: {order.customer_id}</p>
          <p>delivery_service: {order.delivery_service}</p>
          <p>shardkey: {order.shardkey}</p>
          <p>sm_id: {order.sm_id}</p>
          <p>date_created: {order.date_created}</p>
          <p>oof_shard: {order.oof_shard}</p>
        </div>
        :
      
      
        <div></div>
      }
    </div>
  );
}

export default App;
