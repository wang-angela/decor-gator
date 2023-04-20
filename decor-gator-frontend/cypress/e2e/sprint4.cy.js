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

    cy.get('.post-display').last().click()
    cy.get('label.post-furniture-type-2').contains('Sofa')
    cy.get('label.post-price-2').contains('$22')
    cy.get('.post-description-2').contains('Selling 13 years old brick sofa.')

    cy.contains('← Back').click()

    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })

  it('Searches posts by category and title', () => {
    cy.wait(3000)
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

  it('Changes current email and password', () => {
    // Displays current email
    cy.get('button.editUser-button').click()
    cy.contains('Change Email').click()
    cy.contains('Current Email: djohnson@gmail.com')

    cy.get('input.search-text-input').first().type('hi@gmail.com')
    cy.get('input.search-text-input').last().type('123456')
    cy.get('.change-userInfo-button').click()
    cy.wait(15000)

    cy.visit('http://localhost:3000/')
    cy.get('input[placeholder="Email"]').last().type('hi@gmail.com')
    cy.get('input[placeholder="Password"]').last().type('123456')
    cy.contains('SIGN IN').click()
    cy.wait(5000)
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')

    cy.get('button.editUser-button').click()
    cy.contains('Change Password').click()

    cy.get('.search-text-input').first().type('123456')
    cy.get('.search-text-input').last().type('1')
    cy.contains('Enter').click()
    cy.wait(13000)

    cy.visit('http://localhost:3000/')
    cy.get('input[placeholder="Email"]').last().type('hi@gmail.com')
    cy.get('input[placeholder="Password"]').last().type('1')
    cy.contains('SIGN IN').click()
    cy.wait(5000)
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })

})
  