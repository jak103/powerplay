export {} // Prevents linting errors

describe('Example spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:9001/#/') 
  })

  // Define the teams array as per your application
  const teams = [
    {
      id: 1,
      name: 'District 5',
      logo: 'path/to/district5_logo.png',
      league: 'B',
      manager: 'Captain Hook',
    },
  ];

  beforeEach(() => {
    // Visit the page where the items are rendered
    cy.visit('http://localhost:9001/#/'); // Change this to the actual URL
  });

  teams.forEach(team => {
    it(`should click on the team item with id team-${team.id}`, () => {
      // Find the team item with the specified id and click it
      cy.get(`#team-${team.id}`).click();

      // Optionally, you can add assertions to verify the navigation
      cy.url().should('include', `/teaminfo/District%25205`); // Adjust this as needed
    });

    it(`should click on the league item with id league-${team.league}`, () => {
      // Find the league item with the specified id and click it
      cy.get(`#league-${team.league}`).click();

      // Optionally, you can add assertions to verify the navigation
      cy.url().should('include', `/leagueinfo/${team.league}`); // Adjust this as needed
    });
  });
})