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

//console.log("Service worker registered"); // TEST

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open('v1').then((cache) => {
      return cache.addAll([
        '/app/',
        // Add other assets as needed
      ]);
    })
  );
});

self.addEventListener('fetch', (event) => {
  console.log("Fetch event", event.request.url) // TEST
  console.log("Fetch event2 ", event.request) // TEST
  event.respondWith(
    // Try to fetch the resource from the network
    fetch(event.request).then((networkResponse) => {
      // If successful, open the cache, store the response, and return it
      return caches.open('v1').then((cache) => {
        cache.put(event.request, networkResponse.clone());
        return networkResponse;
      });
    }).catch(() => {
      // If the network is unavailable, try to get the resource from the cache
      return caches.match(event.request);
    })
  );
});