import React, { useState, useEffect } from 'react';
import { getRequest } from '../app/axiosClient';

/**
 * Component for rendering a list of todos.
 */
function Todos() {
  const [todos, setTodos] = useState([]);

  /**
   * Fetches the todos from the server and updates the state.
   */
  useEffect(() => {
    getRequest('gettodos').then(response => {
      setTodos(response.data);
    });
  }, []);

  return (
    <>
      {todos.map((todo) => (
        <div key={todo.ID} id={todo.ID}>{todo.name}</div>
      ))}
    </>
  );
}

export default Todos;