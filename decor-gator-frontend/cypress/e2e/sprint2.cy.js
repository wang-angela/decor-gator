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
  
  it('Alerts if login information is invalid', () => {
    // Put in invalid passsword
    cy.get('input[placeholder="Email"').last().type('@')
    cy.get('input[placeholder="Password"').last().type('123456')
    cy.contains('SIGN IN').click()
    
    cy.wait(5000)

    // Test if invalid password message pops off
    cy.on('window:alert', (t) => {
      expect(t).to.contains('Please enter valid credentials.');
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
  })
    
})
  