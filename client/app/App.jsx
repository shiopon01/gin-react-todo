import * as React from 'react';

export default class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      message: ''
    };
  }

  componentDidMount() {
    fetch('/api/todo/')
      .then(x => x.json())
      .then(json => {
        this.setState({
          message: json.message
        });
      })
  }

  render() {
    const { message } = this.state;

    return (
      <div>HELLO { message }</div>
    );
  }
}
