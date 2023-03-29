/// <reference types="cypress" />



describe('decor-gator Post and Buy Page', () => {
  
  beforeEach (() => {
    cy.viewport(1920, 1080)
    // Visit Website
    cy.visit('http://localhost:3000/')
    
    // Sign up
    // cy.contains('Don\'t have an account? SIGN UP').click()
    // // Write in sign up information
    // cy.get('input[placeholder="First Name"]').type('Dwayne', {force: true})
    // cy.get('input[placeholder="Last Name"]').type('Johnson', {force: true})
    // cy.get('input[placeholder="Email"]').first().type('djohnson@gmail.com', {force: true})
    // cy.get('input[placeholder="Password"]').first().type('123456', {force: true})
    // cy.contains('SIGN UP').click({force: true})
    
    // cy.wait(5000)

    // // Test if successful message popped off
    // cy.on('window:alert', (t) => {
    //   expect(t).to.contains('User successfully created!');
    // })

    // cy.contains('LOG IN').click({force: true})
    
    cy.get('input[placeholder="Email"').last().type('djohnson@gmail.com')
    cy.get('input[placeholder="Password"').last().type('123456')
    cy.contains('SIGN IN').click()
    cy.wait(5000)
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })  

  it('Alerts if any of post information is missing', () => {
    // Go to Post Page
    cy.contains('+ Post').click()
    cy.url().should('be.equal', 'http://localhost:3000/PostPage')

    // Only fill in one field
    cy.get('.post-title').type('Nice old table')
    // Submit Post
    cy.get('.post-submit-button').click()
    // Check if submission was successful
    cy.on('window:alert', (t) => {
      expect(t).to.contains('Please enter all fields.');
    })
  })

  it('Clicks + Post button and submit a post', () => {

    cy.contains('+ Post').click()
    // Fill in post information
    cy.get('.post-title').type('SELLING BRICK CHAIR')
    cy.get('.post-furniture-type').type('Chair')
    // Submit Post
    cy.get('.post-submit-button').click()
    cy.wait(5000)
    // Check if submission was successful
    cy.on('window:alert', (t) => {
      expect(t).to.contains('Post successfully created!');
    })
    // Redirect to Post page
    cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })

  it('Shows New Posts Created', () => {
    // Check if the post is visible
    cy.get('.container').contains('SELLING BRICK CHAIR')
    cy.get('.container').contains('Chair')
  })
  
  it('Creates new page if posts go over 8', () => {
    for (let i = 1; i <= 8; ++i) {
      // Go to post page
      cy.contains('+ Post').click()
      // Make a new post
      cy.get('.post-title').type('' + i)
      cy.get('.post-furniture-type').type('Chair')
      cy.get('.post-submit-button').click()
    }
    cy.wait(5000)

    // Is able to go to next page and see post name '7'
    cy.contains('Next Page').click()
    cy.get('.container').first().contains('8')
  })

  it('Searches by post titles', () => {
    // Go to post page
    cy.contains('+ Post').click()
    // Make a new post
    cy.get('.post-title').type('SELLING BRICK SOFA')
    cy.get('.post-furniture-type').type('Sofa')
    cy.get('.post-submit-button').click()

    // Search the first post created
    cy.get('.search-text-input').type('BRICK')
    cy.get('.search-button').click()
    cy.contains('SELLING BRICK CHAIR')
    cy.contains('SELLING BRICK SOFA')
  })
})
  