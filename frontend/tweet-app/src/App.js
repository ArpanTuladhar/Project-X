import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [tweetContent, setTweetContent] = useState('');
  const [tweets, setTweets] = useState([]);
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetchTweets();
  }, []);

  const fetchTweets = async () => {
    try {
      const response = await fetch('http://localhost:8080/tweets');
  
      if (!response.ok) {
        throw new Error(`Failed to fetch tweets: ${response.status} ${response.statusText}`);
      }
  
      const data = await response.json();
      setTweets(data);
    } catch (error) {
      console.error('Error fetching tweets:', error);
    }
  };
  

  const handleSubmit = async (event) => {
    event.preventDefault();
  
    try {
      const response = await fetch('http://localhost:8080/create_tweet', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `tweetContent=${encodeURIComponent(tweetContent)}`,
        mode: 'cors',
      });
  
      if (!response.ok) {
        throw new Error('An error occurred while creating the tweet.');
      }
  
      // Fetch tweets again to update the state
      await fetchTweets();
  
      setTweetContent('');
      setMessage('Tweet created successfully');
    } catch (error) {
      console.error('Error:', error);
      setMessage('An error occurred while creating the tweet.');
    }
  };
  

  const handleEdit = async (id, updatedContent) => {
    // Similar to your previous code
  };

  const handleDelete = async (id) => {
    // Similar to your previous code
  };

  return (
    <div className="App">
      <h1>Create a Tweet</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="tweetContent">Tweet Content:</label>
        <br />
        <textarea
          id="tweetContent"
          name="tweetContent"
          rows="4"
          cols="50"
          required
          value={tweetContent}
          onChange={(e) => setTweetContent(e.target.value)}
        ></textarea>
        <br />
        <br />
        <input type="submit" value="Create Tweet" />
      </form>

      {message && (
        <div className={message.includes('error') ? 'error-message' : 'success-message'}>
          {message}
        </div>
      )}

      <div className="tweet-list">
        <h2>Tweets</h2>
        {tweets.map((tweet) => (
          <div key={tweet.id} className="tweet">
            <p>{tweet.content}</p>
            <button onClick={() => handleEdit(tweet.id, 'Updated Content')}>Edit</button>
            <button onClick={() => handleDelete(tweet.id)}>Delete</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
