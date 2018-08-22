import React, { Component } from 'react'
import * as React from 'react'

class App extends Component {
  constructor(props) {
    super(props)

    this.state = {
      message: ''
    }
  }

  componentDidMount() {
    fetch('/api/todo/')
    .then(x => x.json())
    .then(json => {
      this.setState({
        message: json.message
      })
    })
  }

  render() {
    const { message } = this.state

    return (
      <div>MESSAGE: { message }</div>
    )
  }
}

export default App