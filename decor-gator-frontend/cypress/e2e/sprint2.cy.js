/// <reference types="cypress" />



describe('decor-gator signup and login test', () => {

  beforeEach(() => {
    cy.visit('http://localhost:3000/')
  })

  it('Flips the login screen to signup page', () => {
    // Click to see if sign up button flips the card over
    cy.contains('Don\'t have an account? SIGN UP').click()
    cy.get('.menu-back-signup').should('be.visible')
    cy.get('.menu-front-login').should('not.be.visible')
  })

  it('Returns error message with no sign up information entered', () => {
      // Go to back page to sign up
      cy.contains('Don\'t have an account? SIGN UP').click()

      // Logs error if sign up fields are not fully entered
      cy.contains('SIGN UP').click({force: true})
      cy.on('window:alert', (t) => {
        expect(t).to.contains('Please enter all fields.');
      })      
  })

  it('Signs up with information entered', () => {
    cy.contains('Don\'t have an account? SIGN UP').click()
    // Write in sign up information
    cy.get('input[placeholder="First Name"]').type('Dwayne', {force: true})
    cy.get('input[placeholder="Last Name"]').type('Johnson', {force: true})
    cy.get('input[placeholder="Email"]').first().type('djohnson@gmail.com', {force: true})
    cy.get('input[placeholder="Password"]').first().type('123456', {force: true})
    cy.contains('SIGN UP').click({force: true})
    
    cy.wait(5000)

    // Test if successful message popped off
    cy.on('window:alert', (t) => {
      expect(t).to.contains('User successfully created!');
    })
    
  })

  it('Alerts if sign up attempted twice', () => {
    cy.contains('Don\'t have an account? SIGN UP').click()
    // Write in sign up information
    cy.get('input[placeholder="First Name"]').type('Dwayne', {force: true})
    cy.get('input[placeholder="Last Name"]').type('Johnson', {force: true})
    cy.get('input[placeholder="Email"]').first().type('djohnson@gmail.com', {force: true})
    cy.get('input[placeholder="Password"]').first().type('123456', {force: true})
    // Test if it alerts signup attempt with already existing email
    cy.contains('SIGN UP').click({force: true})
    
    cy.wait(5000)

    cy.on('window:alert', (t) => {
      expect(t).to.contains('Email already registered.');
    })
  })
  
  it('Alerts if login information is invalid', () => {
    // Put in invalid passsword
    cy.get('input[placeholder="Email"').last().type('djohnson@gmail.com')
    cy.get('input[placeholder="Password"').last().type('1')
    cy.contains('SIGN IN').click()
    
    cy.wait(5000)

    // Test if invalid password message pops off
    cy.on('window:alert', (t) => {
      expect(t).to.contains('Invalid password');
    })
  })

  it('Logs in and redirects to Buy Page', () => {
    // Attempt login
    cy.get('input[placeholder="Email"').last().type('djohnson@gmail.com')
    cy.get('input[placeholder="Password"').last().type('123456')
    cy.contains('SIGN IN').click()

    cy.wait(10000)
    cy.window().then((win) => {
      cy.spy(win.console, "log");
    });
    // Test if login successful message popped off
    // cy.on('window:alert', (t) => {
    //   expect(t).to.contains('Login successful!', {timeout: 5000});
    // })
    
    // // Redirects to Buy Page
    // cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
  })
    
})

  describe('decor-gator Posting', () => {

    beforeEach(() => {
      // Login and get redirected to Buy Page
      cy.visit('http://localhost:3000/')
      cy.get('input[placeholder="Email"').last().type('djohnson@gmail.com')
      cy.get('input[placeholder="Password"').last().type('123456')
      cy.contains('SIGN IN').click()
      cy.wait(7000)
      cy.url().should('be.equal', 'http://localhost:3000/BuyPage')
    })

    it('Clicks + Post button and submit a post', () => {
      // Click on + Post button, redirect to Post Page
      cy.contains('+ Post').click()
      cy.url().should('be.equal', 'http://localhost:3000/PostPage')

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
    })

    it('Clicks Show Post button and lists created post', () => {
      // Click on List Post button
      cy.contains('List Posts').click()

      // Check if user email, furniture type, post title are logged into console
      cy.window().then((win) => {
        cy.spy(win.console, "log");
      });
    })
  })
  