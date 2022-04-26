import React from 'react';
import './List.css'
export default function List(props) {
  return (
    <div>
        <ul className="list-group mt-3">
            { props.lists.map((todo, index) => (
            <li
                key={index}
                className="list-group-item d-flex justify-content-between align-items-center"
            >
              <input
                type="checkbox"
                checked={todo.isfinished}
                onChange={() => props.handleChangeFinish(todo)}
              />
              <div style={{
                textDecoration: todo.isfinished ? 'line-through' : 'none',
                }}>{ todo.name }</div>
              <button className="delete-button" onClick={() => props.handleDelete(todo)}>
                DELETE
              </button>
            </li>
            ))}
        </ul>
    </div>

  )
}