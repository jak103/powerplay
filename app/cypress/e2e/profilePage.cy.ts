export {}

describe('Profile page tests', () => {
    it('should load the page', () => {
        cy.visit('http://localhost:9001/#/profile');
    }),

    it('should take me to the edit-profile page', () => {
        cy.visit('http://localhost:9001/#/profile');
        cy.get('#edit-profile-button').click();
        cy.url().should('eq', 'http://localhost:9001/#/profile/edit-profile');
    })

    it('should take me to the replace-image page', () => {
        cy.visit('http://localhost:9001/#/profile');
        cy.get('#replace-image-button').click();
        cy.url().should('eq', 'http://localhost:9001/#/profile/replace-image');
    })

    it('should take me to the C league page', () => {
        cy.visit('http://localhost:9001/#/profile');
        cy.get('a[href="#/leagueinfo/C"]').click();
        cy.url().should('eq', 'http://localhost:9001/#/leagueinfo/C');
    })

    it('should take me to the D league page', () => {
        cy.visit('http://localhost:9001/#/profile');
        cy.get('a[href="#/leagueinfo/D"]').click();
        cy.url().should('eq', 'http://localhost:9001/#/leagueinfo/D');
    })

    it('should take me to the District 5 team page', () => {
        cy.visit('http://localhost:9001/#/profile');
        cy.get('a[href="#/teaminfo/District 5"]').click();
        cy.url().should('eq', 'http://localhost:9001/#/teaminfo/District%205');
    })

    it('should take me to the Trash Pandas team page', () => {
        cy.visit('http://localhost:9001/#/profile');
        cy.get('a[href="#/teaminfo/Trash Pandas"]').click();
        cy.url().should('eq', 'http://localhost:9001/#/teaminfo/Trash%20Pandas');
    })
})

describe('Replace image tests', () => {
    it('should grow the image', () => {
         cy.visit('http://localhost:9001/#/profile/replace-image');

    cy.get('.q-slider__track')
    .trigger('mousedown', 'topRight')
    .trigger('mouseup');

    cy.get('.q-img')
    .should('have.attr', 'style')
    .and('include', 'width: 220px; height: 220px;')
});

    it('should shrink the image', () => {
        cy.visit('http://localhost:9001/#/profile/replace-image')
        cy.get('.q-slider__track-container')
        .trigger('mousedown', 'topLeft')
        .trigger('mouseup');
        cy.get('.q-img')
        .should('have.attr', 'style')
        .and('include', 'width: 50px; height: 50px;')
    })

    it('should take you back to the main profile page', () => {
        cy.visit('http://localhost:9001/#/profile/replace-image');
        cy.contains('span', 'Cancel').click();
        cy.url().should('eq', 'http://localhost:9001/#/profile');
    })
});