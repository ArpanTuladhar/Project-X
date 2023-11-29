import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Login from './components/Login';
import Signup from './pages/Signup';
import Home from './pages/Home';

const App = () => {
  return (
    <Router>
      <div className="App">
        <Switch>
          <Route path="/signup" component={Signup} />
          <Route path="/home" component={Home} />
          <Route path="/" component={Login} />
        </Switch>
      </div>
    </Router>
  );
};

export default App;
