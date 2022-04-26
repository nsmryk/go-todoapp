import React from 'react';
import './InputField.css'
export default function InputField(props) {
  return (
    <div className="todo-input-field">
      <input placeholder="入力してください" value={props.text} onChange={e => props.handleChange(e)}></input>
      <button className="add-button" onClick={props.handleSubmit}>ADD</button>
    </div>
  )
}