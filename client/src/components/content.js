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
    return <h3>{this.state.pong}</h3>;
  }
}

class FormToApi extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      inputContent: "",
      responseContent:"",
    };
  }
  handleChange = (e) => {
    e.preventDefault();
    this.setState({ inputContent: e.target.value });
  }

  handleSubmit = () => {
    axios
      .get("/api/tobehashed", {
        params: {
          inputValue: this.state.inputContent,
        },
      })
      .then((response) => this.props.handleTextMessage1(response.data.msg))
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
          color="info"
          sx={{bgcolor:"#EFFFDF"}}
          onInput={(evt) => this.handleChange(evt)}
        />
        <Button
          variant="contained"
          onClick={(e) => this.handleSubmit(e)}
          sx={{ p: 2}}
        >
          Wyślij
        </Button>
      </>
    );
  }
}
class Content extends React.Component {
  state = {
    textMessage1: "",
  };

  handleTextMessage1 = (text) => {
    this.setState({ textMessage1: text });
  };
  render() {
    return (
      <span>
        <DisplayFromApi />
        <FormToApi handleTextMessage1={this.handleTextMessage1} />
        <br />
        <br />
        <HashedFuncTextField textMessage1={this.state.textMessage1} />
      </span>
    );
  }
}

class HashedFuncTextField extends React.Component {
  render() {
    return (
      <>{this.props.textMessage1 ? this.props.textMessage1 : <div></div>}</>
    );
  }
}
export default Content;
