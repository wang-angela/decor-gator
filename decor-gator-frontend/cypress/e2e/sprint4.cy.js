/// <reference types="cypress" />



describe('decor-gator Post and Buy Page', () => {
  
  beforeEach (() => {
    cy.viewport(1920, 1080)
    // Visit Website
    cy.visit('http://localhost:3000/')
    
    cy.get('input[placeholder="Email"]').last().type('djohnson@gmail.com')
    cy.get('input[placeholder="Password"]').last().type('123456')
    cy.contains('SIGN IN').click()
    cy.wait(5000)
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })  

  it('Displays details of a post when clicked on', () => {
    // Go to Post Page
    cy.contains('BRICK CHAIR 1').click()

    cy.get('label.post-title-2').contains('SEELING BRICK CHAIR 1')
    cy.get('label.post-furniture-type-2').contains('Chair')
    cy.get('label.post-price-2').contains('$22')
    cy.get('post-description-2').contains('Selling 13 years old brick chair.')

    cy.get('button.post-submit-button-2').click()

    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })

  it('Searches posts by category and title', () => {
    // Search post by category
    cy.get('select.post-furniture-type-3').select('Sofa')
    cy.get('button.search-button').click()
    cy.contains('SELLING BRICK SOFA')
    
    // Reload all posts
    cy.get('select.post-furniture-type-3').select('All Categories')
    cy.get('button.search-button').click()

    // Search post by category + title
    cy.get('select.post-furniture-type-3').select('Desk')
    cy.get('input.search-text-input').type('Beautiful')
    cy.get('button.search-button').click()
    cy.contains("Beautiful Desk")
  })

  it('Changes current email', () => {
    // Displays current email
    cy.get('button.editUser-button').click()
    cy.get('button.makePost-button').first.click()
    cy.contains('Current Email: djohnson@gmail.com')


    cy.get('input.search-text-input').first.type('hi@gmail.com')
    cy.get('input.search-text-input').last.type('123456')
    cy.get('button.makePost-button').click()
    cy.on('window:alert', (t) => {
      expect(t).to.contains('Email successfully updated!');
    })

    cy.visit('http://localhost:3000/')
    cy.get('input[placeholder="Email"]').last().type('hi@gmail.com')
    cy.get('input[placeholder="Password"]').last().type('123456')
    cy.contains('SIGN IN').click()
    cy.wait(5000)
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })

  it('Changes current password', () => {
    // Displays current email
    cy.get('button.editUser-button').click()
    cy.get('button.makePost-button').last.click()

    cy.contains('Current Password').first.type('123456')
    cy.contains('New Password').type('1')
    cy.contains('Enter').click()
    cy.on('window:alert', (t) => {
      expect(t).to.contains('Password successfully updated!');
    })

    cy.visit('http://localhost:3000/')
    cy.get('input[placeholder="Email"]').last().type('hi@gmail.com')
    cy.get('input[placeholder="Password"]').last().type('1')
    cy.contains('SIGN IN').click()
    cy.wait(5000)
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })
})
  