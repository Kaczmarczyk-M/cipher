import * as React from "react";
import axios from "axios";
import { Button, TextField } from "@mui/material";
import { styled } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";

const Item = styled(Paper)(({ theme }) => ({
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: "center",
  bgcolor: "#151300",
  fontSize: 20,
  whiteSpace:'normal',
  color: theme.palette.text.primary,
}));

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
    return (
      <Box p="30px">
        {this.state.pong}
      </Box>
    );
  }
}

class FormToApi extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      inputContent: "",
      responseContent: "", //is it in use?
    };
  }
  handleChange = (e) => {
    e.preventDefault();
    this.setState({ inputContent: e.target.value });
  };

  handleSubmit = () => {
    axios
      .get("/api/tobehashed", {
        params: {
          inputValue: this.state.inputContent,
        },
      })
      .then((response) => this.props.handleTextMessage1(response.data))
      .catch(function (error) {
        console.log(error);
      });
  };
  render() {
    return (
      <>
        <TextField
          id="filled-basic"
          label="Here"
          variant="filled"
          color="info"
          sx={{ bgcolor: "#EFFFDF" }}
          onInput={(evt) => this.handleChange(evt)}
        />

        <Button
          variant="contained"
          onClick={(e) => this.handleSubmit(e)}
          sx={{ p: 2, ml: 2 }}
        >
          Send
        </Button>
      </>
    );
  }
}
class Content extends React.Component {
  state = {
    textMessage1: "",
    jump: "",
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
        <HashedFuncTextField
          textMessage1={this.state.textMessage1}
          jump={this.state.jump}
        />
      </span>
    );
  }
}

class HashedFuncTextField extends React.Component {
  render() {
    return (
      <>
        {this.props.textMessage1 ? (
          <Box width="25em"sx={{ flexGrow: 1 }}>
            <Grid container spacing={2}>
              <Grid item xs={3}>
                <Item>Caesar Cipher</Item>
              </Grid>
              <Grid item xs={7}>
                <Item>{this.props.textMessage1.msgCaesar}</Item>
              </Grid>
              <Grid item xs={2}>
                <Item>Jump: {this.props.textMessage1.jump}</Item>
              </Grid>
              {/* next */}
              <Grid item xs={3}>
                <Item>Hash: MD5</Item>
              </Grid>
              <Grid item xs={9}>
                <Item>{this.props.textMessage1.md5}</Item>
              </Grid>
            </Grid>
          </Box>
        ) : (
          <></>
        )}
      </>
    );
  }
}
export default Content;
