import React from 'react';
import InputField from './InputField';
import List from './List';
import http from './http';
import './Todos.css';

class Todo extends React.Component {
  constructor() {
    super();
    this.state = {
      name: "",
      lists: [],
    }
  }

  componentDidMount() {
    this.getTodoList();
  }

  getTodoList = () => {
    return http
      .get('/api/tasks')
      .then((response) => {
        console.log(response.data)
        this.setState({ lists: response.data })
      })
      .catch(error => {
        console.log(error)
      })
  }

  handleChange = e => {
    this.setState({ name: e.target.value })
  }

  handleSubmit = () => {
    if (this.state.name === "") {
      return window.alert("入力してください")
    }
    return http
      .post('/api/tasks', {
        name: this.state.name
      })
      .then(() => {
        this.setState({ name: "" });
        this.getTodoList();
      }
      )
      .catch(error => {
        console.log(error)
      })
  }

  handleChangeFinish = (list) => {
    return http
      .put(`/api/tasks/${list.id}`)
      .then(() =>
        this.getTodoList()
      )
      .catch(error => {
        console.log(error)
      })
  }
  
  handleDelete = (list) => {
    return http
      .delete(`/api/tasks/${list.id}`)
      .then(() =>
        this.getTodoList()
      )
      .catch(error => {
        console.log(error)
      })
  }

  render() {
    return (
      <div className="todo">
        <div className="todo-title">
          <h1 >TODO</h1>
        </div>
        <InputField
          name={this.state.name}
          handleChange={this.handleChange}
          handleSubmit={this.handleSubmit}
        />
        <List
          lists={this.state.lists}
          handleDelete={this.handleDelete}
          handleChangeFinish={this.handleChangeFinish}
        />
      </div>
    )
  }
}

export default Todo;