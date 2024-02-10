self.addEventListener('push', (e) => {
  console.log("Received push", e)
  console.log("Showing notitification");
  var options = {
    body: 'This notification was generate from a push.',
    vibrate: [100, 50, 100],
    data: {
      dateOfArrival: Date.now(),
      primaryKey: '2'
    },
    actions: [
      {
        action: 'explore',
        title: 'Explore this new world',
      },
      {
        action: 'close',
        title: 'Close',
      }
    ]
  };
  e.waitUntil(self.registration.showNotification('Hello World', options));
})
