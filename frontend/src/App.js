import React, { Component } from 'react';
import { Button, Form } from 'semantic-ui-react'
import logo from './logo.svg';
import './App.css';

class LinkForm extends Component {
  render() {
    return (
      <Form>
        <Form.Input label='Link URL' type='text' />
        <Button type='submit'>Submit</Button>
      </Form>
    )
  }
}


class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        <LinkForm />
      </div>
    );
  }
}

export default App;
