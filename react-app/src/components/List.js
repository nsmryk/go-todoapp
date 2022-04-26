import React from 'react';

export default function List(props) {
  return (
    <div>
        <h1>ToDo List</h1>
        <ul>
            { props.lists.map((todo, index) => (
            <li
                key={index}
                style={{
                textDecoration: todo.isfinished ? 'line-through' : 'none',
                }}
            >
            { todo.name }: { todo.isfinished } 
            <input
                type="checkbox"
                checked={todo.isfinished}
                onChange={() => props.handleChangeFinish(todo)}
            />
            <button onClick={() => props.handleDelete(todo)}>
              DELETE
            </button>
            </li>
            ))}
        </ul>
    </div>

  )
}