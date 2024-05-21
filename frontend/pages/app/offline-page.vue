<template>
	<div class="offline-container">
	  <h1>You are offline</h1>
	  <p>Please check your internet connection.</p>
	  <button @click="tryReconnect">Try Reconnecting</button>
	</div>
  </template>
  
  <script>
  export default {
	data() {
	  return {
		// Store a reference to the bound method so it can be added and removed consistently
		boundRedirectUser: null,
	  };
	},
	mounted() {
	  this.boundRedirectUser = this.redirectUser.bind(this);
	  // Check if the user is already back online when the component mounts
	  if (navigator.onLine) {
		this.redirectUser();
	  } else {
		// Listen for the online event to detect when the user comes back online
		window.addEventListener('online', this.boundRedirectUser);
	  }
	},
	beforeDestroy() {
	  // Clean up the event listener when the component is destroyed
	  if (this.boundRedirectUser) {
		window.removeEventListener('online', this.boundRedirectUser);
	  }
	},
	methods: {
	  tryReconnect() {
		// Attempt to reload the page
		this.redirectUser();
	  },
	  redirectUser() {
		window.history.back();
	  },
	},
  };
  </script>
  
  <style scoped>
  .offline-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    text-align: center;
  }
  
  button {
    margin-top: 20px;
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
  }
  </style>