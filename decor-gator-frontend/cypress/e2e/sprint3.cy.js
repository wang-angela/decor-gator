/// <reference types="cypress" />



describe('decor-gator Post and Buy Page', () => {
  
  beforeEach (() => {
    cy.viewport(1920, 1080)
    // Visit Website
    cy.visit('http://localhost:3000/')
    
    cy.get('input[placeholder="Email"]').last().type('djohnson@gmail.com')
    cy.get('input[placeholder="Password"]').last().type('123456')
    cy.contains('SIGN IN').click()
    cy.wait(7000)
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
    cy.get('.post-furniture-type').select('Chair')
    cy.get('input.post-price').type('22')
    cy.get('textarea.post-description').type('Selling 13 years old brick chair.')

    //upload image
    cy.get('button.file-upload-display').click()
    cy.wait(10000)

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
  })
  
  it('Creates new page if posts go over 8', () => {
    for (let i = 1; i <= 6; ++i) {
      // Go to post page
      cy.contains('+ Post').click()
      // Make a new post
      cy.get('.post-title').type('BRICK CHAIR '+i)
      cy.get('.post-furniture-type').select('Chair')
      cy.get('input.post-price').type('22')
      cy.get('textarea.post-description').type('Selling 13 years old brick chair.')

      //upload image
      cy.get('button.file-upload-display').click()
      cy.wait(5000)

      cy.get('.post-submit-button').click()
    }
    cy.wait(5000)

    // Is able to go to next page and see post name '7'
    cy.get('button.next-page-button').click()
    cy.get('.container').contains('6')
  })

  it('Searches by post titles', () => {
    // Go to post page
    cy.contains('+ Post').click()
    // Make a new post
    cy.get('.post-title').type('BRICK SOFA')
    cy.get('.post-furniture-type').select('Sofa')
    cy.get('input.post-price').type('22')
    cy.get('textarea.post-description').type('Selling 13 years old brick sofa.')
    cy.get('button.file-upload-display').click()
    cy.wait(5000)

    cy.get('.post-submit-button').click()

    // Search the first post created
    cy.get('.search-text-input').type('BRICK')
    cy.get('.search-button').click()
    cy.wait(3000)
    cy.get('.search-button').click()
    cy.wait(3000)
    cy.contains('SELLING BRICK CHAIR')
  })
})
  