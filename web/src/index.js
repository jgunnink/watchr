import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';
import 'antd/dist/antd.css';
import { LocaleProvider } from 'antd';
import enUS from 'antd/lib/locale-provider/en_US';

ReactDOM.render(
  <LocaleProvider locale={enUS}><App /></LocaleProvider>,
  document.getElementById('root')
);
