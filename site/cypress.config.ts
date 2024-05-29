import { defineConfig } from 'cypress';

export default defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      console.log(on); // Temporarily prevents lint errors
      console.log(config); // Temporarily prevents lint errors
      // TODO: implement node event listeners here
    },
  },
});
