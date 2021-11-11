import * as React from "react";
import axios from "axios";
import { Button, TextField } from "@mui/material";

class DisplayFromApi extends React.Component {
  constructor() {
    super();
    this.state = {
      pong: "pending",
    };
  }
  componentWillMount() {
    axios
      .get("api/ping")
      .then((response) => {
        this.setState(() => {
          return { pong: response.data.message };
        });
      })
      .catch(function (error) {
        console.log(error);
      });
  }

  render() {
    return <h1>Ping {this.state.pong}</h1>;
  }
}

class FormToApi extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      inputContent: "",
    };
    // this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChange(e) {
    e.preventDefault();
    this.setState({ inputContent: e.target.value });
  }
  handleSubmit() {
    axios
      .get("/api/tobehashed", {
          params: {
              inputValue: this.state.inputContent
          }
      })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
  }
  render() {
    return (
      <>
        <TextField
          id="filled-basic"
          label="Here"
          variant="filled"
          onInput={(evt) => this.handleChange(evt)}
        />
        <Button
          variant="contained"
          onClick={(e) => this.handleSubmit(e)}
          sx={{ p: 2 }}
        >
          Wyślij
        </Button>
      </>
    );
  }
}

class HashedFuncTextField extends React.Component {
    state = {
        textValue:''
    }
    // componentDidMount(){
    //     axios.get('/hashed').then(function(res) {console.log(res)})
    // }
    render() {
        if (this.state.textValue === '') {
            return(
                <span>
                    tutaj res
                </span>
            )
        }
        return (
            <>
            </>
        );
    }
}
class Content extends React.Component {
  render() {
    return (
      <span>
        <DisplayFromApi />
        <FormToApi />
        <br/>
        <br/>
        <HashedFuncTextField/>
      </span>
    );
  }
}

export default Content;
