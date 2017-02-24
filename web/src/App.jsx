import React, { Component } from 'react';
import './App.css';
import { Layout } from 'antd';
import RegisterForm from './components/registerForm'
const { Content, Footer } = Layout;

class App extends Component {
  render() {
    return (
      <div>
        <Content style={{ padding: '0 50px' }}>
          <div style={{ background: '#fff', padding: 24, minHeight: 280 }}>
            <h1>HR Monitoring Registration</h1>
            <h4>Create your account and get notified!</h4>
            <div style={{ maxWidth: '300px', marginTop: '20px' }}>
              <RegisterForm />
            </div>
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          Watchr. Created for fun in Perth, Western Australia
        </Footer>
      </div>
    );
  }
}

export default App;
