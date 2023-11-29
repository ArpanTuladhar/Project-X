import React, { useState, useEffect } from 'react';

const Home = () => {
  const [tweets, setTweets] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Fetch tweets from the backend
    const fetchTweets = async () => {
      try {
        const response = await fetch('http://localhost:8080/tweets');
        if (!response.ok) {
          throw new Error('Failed to fetch tweets');
        }
        const data = await response.json();
        setTweets(data);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching tweets:', error);
      }
    };

    fetchTweets();
  }, []);

  return (
    <div>
      <h2>Home</h2>

      {loading ? (
        <p>Loading tweets...</p>
      ) : (
        <ul>
          {tweets.map((tweet) => (
            <li key={tweet.id}>{tweet.content}</li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Home;
