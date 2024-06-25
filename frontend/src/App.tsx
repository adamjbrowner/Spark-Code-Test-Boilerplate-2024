import React, { useEffect, useState } from 'react';
import './App.css';
import Todo, { TodoType } from './Todo';

function App() {
  const [todos, setTodos] = useState<TodoType[]>([]);

  // Initially fetch todo
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const todos = await fetch('http://localhost:8080/');
        if (todos.status !== 200) {
          console.log('Error fetching data');
          return;
        }

        setTodos(await todos.json());
      } catch (e) {
        console.log('Could not connect to server. Ensure it is running. ' + e);
      }
    }

    fetchTodos()
  }, []);

  const [newToDo, setNewToDo] = useState<TodoType>({
    title: '',
    description: ''
  });

  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo) =>
          <Todo
            key={todo.title + todo.description}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form>
        <input placeholder="Title" name="title" value={newToDo.title} autoFocus={true} onChange={event =>
          setNewToDo({
            ...newToDo,
            title: event.target.value
          })
        } />
        <input placeholder="Description" name="description" value={newToDo.description} onChange={event => {
          setNewToDo({
            ...newToDo,
            description: event.target.value
          });
        }} />
        <button>Add Todo</button>
      </form>
    </div>
  );
}

export default App;
