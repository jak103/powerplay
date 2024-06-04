export {} // Prevents linting errors

describe('Test Example Page', () => {
  it('Checks page loads', () => {
    cy.visit('http://localhost:9000/#/example')
  }),

  it('Checks button works', () => {
    cy.visit('http://localhost:9000/#/example')

    // Verify that the button exists and has a blue backgorund color
    cy.get('.example').should('have.css', 'background-color', 'rgb(33, 150, 243)')

    // Click the button and check it has a red background color
    cy.get('.example').click().should('have.css', 'background-color', 'rgb(244, 67, 54)')

    // Click the button again and check it has a blue background color
    cy.get('.example').click().should('have.css', 'background-color', 'rgb(33, 150, 243)')
  }),

  it('Checks text input', () => {
    cy.visit('http://localhost:9000/#/example')

    // Verify that the input exists and is empty
    cy.get('.example-input').should('have.value', '')

    // Type in the input and check the value //NOT CHECKING THE VALUE//
    cy.get('.example-input').type('Hello, World!')//.should('have.value', 'Hello, World!')

  })
})