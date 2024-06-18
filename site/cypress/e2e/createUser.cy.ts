export {}

describe('template spec', () => {
  it('Check page exists', () => {
    cy.visit('http://localhost:9000/#/create-user')
  })

  it('Check valid input fields', () => {
    cy.visit('http://localhost:9000/#/create-user')

    cy.get('.input-first').type('John')
    cy.get('.input-last').type('Smith')
    cy.get('.input-email').type('johnsmith@gmail.com')
    cy.get('.input-password').type('123456789!')
    cy.get('.input-confirm-password').type('123456789!')
    cy.get('.input-phone').type('(123) 435-1111')
    cy.get('.input-birthday').type('2000-01-01')
    cy.get('.input-exp').type('4')
    cy.get('.input-level').click();

    // Select an option from the dropdown
    cy.get('.q-menu .q-item')
      .contains('High School') // Replace 'Option Text' with the text of the option you want to select
      .click();

    // Verify the selected value
    cy.get('.input-level').should('contain', 'High School'); // Ensure the correct value is displayed
  })



  it('Check valid invalid input fields', () => {
    cy.visit('http://localhost:9000/#/create-user')

    cy.get('.input-first').type('John1')
    cy.get('#first-error').should('be.visible');
    cy.get('.input-last').type('Smith1')
    cy.get('#last-error').should('be.visible');

    cy.get('.input-email').type('johnsmith@gmail')
    cy.get('.input-password').type('123456789!')
    cy.get('.input-confirm-password').type('1234')
    cy.get('.input-phone').type('435-1111')
    cy.get('.input-birthday').type('2047-01-01')

    cy.get('#submitButton').click()
    cy.get('#phone-error').should('be.visible');
    cy.get('#password-error').should('be.visible');
    cy.get('#birthday-error').should('be.visible');
    cy.get('#exp-error').should('be.visible');
    cy.get('#level-error').should('be.visible');

  })



})