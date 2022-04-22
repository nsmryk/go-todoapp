import React from 'react';

export default function List(props) {
  return (
    <div>
        <h1>ToDo List</h1>
        <ul>
            { props.lists.map(todo => (
            <li>{ todo.name }: {todo.isfinished} </li>
            ))}
        </ul>
    </div>

  )
}