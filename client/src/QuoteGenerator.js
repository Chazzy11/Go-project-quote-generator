import React from 'react';
import { useState, useEffect } from 'react';

export default function QuoteGenerator() {
  const [quote, setQuote] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/quotenumber')
      .then(Response => Response.json())
      .then(data => setQuote(data.quote));
  // catch errors
  }, []); // Empty array ensures that effect is only run on mount

  return (
    <div>
    {quote !== null ? (
      <div>
        {quote === 1 ? <p>It does not matter how slowly you go as long as you do not stop</p> : null}
        {quote === 2 ? <p>Wherever you go, go with all your heart</p> : null}
        {quote === 3 ? <p>Never give a sword to a man who can't dance</p> : null}
        {quote === 4 ? <p>To know what you know and what you do not know, that is true knowledge</p> : null}
        {quote === 5 ? <p>I hear and I forget. I see and I remember. I do and I understand</p> : null}
        {quote === 6 ? <p>When it is obvious that the goals cannot be reached, don't adjust the goals, adjust the action steps</p> : null}
        {quote === 7 ? <p>Everything has beauty, but not everyone can see</p> : null}
        {quote === 8 ? <p>Choose a job you love, and you will never have to work a day in your life</p> : null}
        {quote === 9 ? <p>Our greatest glory is not in never falling, but in rising every time we fall</p> : null}
      </div>
    ) : (
      <p>Loading...</p>
    )}
  </div>
);
};
