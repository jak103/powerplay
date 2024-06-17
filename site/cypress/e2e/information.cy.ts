export {}

// Tests Static Information Pages
describe('Test Information Pages', () => {
  it('Checks information index exists', () => {
    cy.visit('http://localhost:9000/#/information')
  }),

  it('Checks how to join exists', () => {
    cy.visit('http://localhost:9000/#/information/how-to-join')
  }),

  it('Checks for players exists', () => {
    cy.visit('http://localhost:9000/#/information/players')
  }),

  it('Players page how to join link works', () => {
    cy.visit('http://localhost:9000/#/information/players')
    cy.get('a').contains('How to Join').click().url().should('include', '/information/how-to-join')
  }),

  it('Checks for managers exists', () => {
    cy.visit('http://localhost:9000/#/information/maangers')
  }),

  it('Checks staff page exists', () => {
    cy.visit('http://localhost:9000/#/information/staff')
  }),

  it('Checks rink page exists', () => {
    cy.visit('http://localhost:9000/#/information/rink')
  }),

  it('Checks substitutions page exists', () => {
    cy.visit('http://localhost:9000/#/information/substitution')
  })
})