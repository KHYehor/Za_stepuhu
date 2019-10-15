import fetch from 'node-fetch';
const URL = '';

export default {
  sendData: (battery, currentVideo) => (
    fetch(`${URL}/sendData`, {
      method: 'POST',
      body: JSON.stringify({
        battery, 
        currentVideo,
        deviceTime: new Date()
      }),
      headers: { 'Content-Type': 'application/json' }
    }).then(res => res.json())
  ),
  getData: count => (
    fetch(`${URL}/getData`, {
      method: 'POST',
      body: JSON.stringify({count}),
      headers: { 'Content-Type': 'application/json' }
    }).then(res => res.json())
  )
};