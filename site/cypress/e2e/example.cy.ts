export {} // Prevents linting errors

describe('Example Test', () => {
  it('Visit the home page', () => {
    cy.visit('http://localhost:9000/#/')
  })
})