export {} // Prevents linting errors

describe('Example spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:9001/#/leagueinfo/B') 
  })

  const items = [
    { icon: 'event', label: 'Schedule' },
    { icon: 'chat', label: 'Chat' },
    { icon: 'people', label: 'Roster' },
    { icon: 'list_alt', label: 'Sub' },
    { icon: 'assessment', label: 'Standings' },
    { icon: 'insights', label: 'Statistics' },
  ];

  beforeEach(() => {
    // Visit the page where the items are rendered
    cy.visit('http://localhost:9001/#/leagueinfo/B'); // Change this to the actual URL
  });

  items.forEach(item => {
    it(`should click on the team item with id team-${item.label}`, () => {
      // Find the team item with the specified id and click it
      cy.get(`#label-${item.label}`).click();

      // Optionally, you can add assertions to verify the navigation
      cy.url().should('include', `/${item.label.toLowerCase()}`); // Adjust this as needed
    });
  });
})