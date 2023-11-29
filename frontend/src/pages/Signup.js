import React, { useState } from 'react';

const Signup = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleSignup = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        throw new Error('Registration failed');
      }

      // Handle successful registration (e.g., show success message, redirect to login)
      setMessage('Account created successfully!');

      // You can add additional logic here, such as redirecting to the login page
    } catch (error) {
      console.error('Error:', error);
      setMessage('Registration failed. Please try again.');
    }
  };

  return (
    <div>
      <h2>Create Account</h2>
      <form onSubmit={handleSignup}>
        <label htmlFor="username">Username:</label>
        <input
          type="text"
          id="username"
          name="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />

        <label htmlFor="password">Password:</label>
        <input
          type="password"
          id="password"
          name="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        <button type="submit">Create Account</button>
      </form>

      {message && <p>{message}</p>}

      {/* Link to Login page */}
      <p>
        Already have an account? <a href="/">Log in</a>
      </p>
    </div>
  );
};

export default Signup;
