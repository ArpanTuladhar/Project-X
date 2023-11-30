// app.js

import React, { useState, useEffect } from 'react';
import './App.css';

const TweetActions = ({ onLike, onComment, onEdit, onDelete }) => (
  <div className="tweet-actions mt-2">
    <button onClick={onLike} className="btn btn-primary mr-2">
      Like
    </button>
    <button onClick={onComment} className="btn btn-secondary mr-2">
      Comment
    </button>
    <button onClick={onEdit} className="btn btn-warning mr-2">
      Edit
    </button>
    <button onClick={onDelete} className="btn btn-danger">
      Delete
    </button>
  </div>
);


function App() {
  const [tweetContent, setTweetContent] = useState('');
  const [tweets, setTweets] = useState([]);
  const [message, setMessage] = useState('');
  const [editTweetId, setEditTweetId] = useState(null);

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

  const handleEdit = async (id, content) => {
    setEditTweetId(id);
    setTweetContent(content);
  };

  const handleEditSubmit = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/edit_tweet/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `tweetContent=${encodeURIComponent(tweetContent)}`,
        mode: 'cors',
      });

      if (!response.ok) {
        throw new Error(`An error occurred while editing the tweet: ${response.status} ${response.statusText}`);
      }

      // Fetch tweets again to update the state
      await fetchTweets();

      setEditTweetId(null);
      setTweetContent('');
      setMessage(`Tweet with ID ${id} edited successfully`);
    } catch (error) {
      console.error('Error:', error);
      setMessage(`Error editing tweet with ID ${id}`);
    }
  };

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/delete_tweet/${id}`, {
        method: 'DELETE',
        mode: 'cors',
      });

      if (!response.ok) {
        throw new Error(`An error occurred while deleting the tweet: ${response.status} ${response.statusText}`);
      }

      // Fetch tweets again to update the state
      await fetchTweets();

      setMessage(`Tweet with ID ${id} deleted successfully`);
    } catch (error) {
      console.error('Error:', error);
      setMessage(`Error deleting tweet with ID ${id}`);
    }
  };

  const handleLike = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/like_tweet/${id}`, {
        method: 'POST',
        mode: 'cors',
      });

      if (!response.ok) {
        throw new Error(`An error occurred while liking the tweet: ${response.status} ${response.statusText}`);
      }

      // Fetch tweets again to update the state
      await fetchTweets();

      setMessage(`Liked tweet with ID ${id}`);
    } catch (error) {
      console.error('Error:', error);
      setMessage(`Error liking tweet with ID ${id}`);
    }
  };

  const handleComment = async (id) => {
    // Implement commenting logic
  };

  return (
    <div className="App">
      <h1>Create a Tweet</h1>
      <form onSubmit={editTweetId !== null ? () => handleEditSubmit(editTweetId) : handleSubmit}>
        <div className="form-group">
          <label htmlFor="tweetContent">Tweet Content:</label>
          <textarea
            className="form-control"
            id="tweetContent"
            name="tweetContent"
            rows="4"
            cols="50"
            required
            value={tweetContent}
            onChange={(e) => setTweetContent(e.target.value)}
          ></textarea>
        </div>
        <button type="submit" className="btn btn-primary">
          {editTweetId !== null ? 'Edit Tweet' : 'Create Tweet'}
        </button>
      </form>

      {message && (
        <div className={message.includes('error') ? 'alert alert-danger mt-3' : 'alert alert-success mt-3'}>
          {message}
        </div>
      )}

      <div className="tweet-list">
        <h2>Tweets</h2>
        {tweets !== null && tweets.length > 0 ? (
          tweets.map((tweet) => (
            <div key={tweet.id} className="tweet border p-3 mb-3 rounded">
              <p>{tweet.content}</p>
              <p>Likes: {tweet.likesCount}</p>
              <p>Comments: {tweet.commentsCount}</p>
              <TweetActions
                onLike={() => handleLike(tweet.id)}
                onComment={() => handleComment(tweet.id)}
                onEdit={() => handleEdit(tweet.id, tweet.content)}
                onDelete={() => handleDelete(tweet.id)}
              />
            </div>
          ))
        ) : (
          <p>No tweets available</p>
        )}
      </div>
    </div>
  );
}

export default App;
