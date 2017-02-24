import axios from 'axios';
import { Form, Icon, Input, Button, Checkbox } from 'antd';
import React, { Component } from 'react';
import './formStyle.css';

class RegisterForm extends Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log('Received values of form: ', values);
        axios.post('http://jk-hyperv-ubuntu:8000/register', values)
             .then(function (response) {
               console.log(response);
             })
      }
    });
  }
  render() {
    const { getFieldDecorator } = this.props.form;
    return (
      <Form onSubmit={this.handleSubmit} className="login-form">
        <Form.Item>
          <label>Email</label>
          {getFieldDecorator('userName', {
            rules: [{ required: true, message: 'Please enter your email address!' }],
          })(
            <Input addonBefore={<Icon type="mail" />} placeholder="Email" />
          )}
        </Form.Item>
        <Form.Item>
          <label>Password</label>
          {getFieldDecorator('password', {
            rules: [{ required: true, message: 'Please enter your password!' }],
          })(
            <Input addonBefore={<Icon type="lock" />} type="password" placeholder="Password" />
          )}
        </Form.Item>
        <Form.Item>
          {getFieldDecorator('remember', {
            valuePropName: 'checked',
            initialValue: true,
          })(
            <Checkbox>Remember me</Checkbox>
          )}
          <a className="login-form-forgot">Forgot password</a>
          <Button type="primary" htmlType="submit" className="login-form-button">
            Log in
          </Button>
          Or <a>register now!</a>
        </Form.Item>
      </Form>
    );
  }
}
const WrappedRegisterForm = Form.create()(RegisterForm);
export default WrappedRegisterForm
