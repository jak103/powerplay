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
  e.waitUntil(self.registration.showNotification('Hello World'));
  console.log("Done showing notification");
})



self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open('v1').then((cache) => {
      return cache.addAll([
        '/app/frontend/app.vue',
        // Add other assets as needed
      ]);
    })
  );
});

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      return response || fetch(event.request);
    })
  );
});