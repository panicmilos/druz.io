import { AuthService } from "../services/AuthService";
import { AuthResponse } from "../models/AuthResponse";
import {  } from "../../auth-context/utils";
import { FC, useContext, useEffect, useState } from "react";
import * as Yup from "yup";
import { createUseStyles } from "react-jss";
import { useNavigate } from "react-router-dom";
import {
  Card,
  Col,
  Container,
  Form,
  Theme,
  Button,
  FormTextInput,
  AuthContext,
  NotificationService,
  extractErrorMessage,
  setAxiosInterceptors,
} from '../imports';
import axios, { AxiosError } from "axios";
import { EMAIL_REGEX, LOWER_CASE_REGEX, NUMERIC_REGEX, SPECIAL_CHARACTERS_REGEX } from "../../../core/constants";

type Props = {};

const useStyles = createUseStyles((theme: Theme) => ({
  outerContainer: {
    height: "100%",
    width: "100%",
    overflow: "hidden",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  container: {
    maxHeight: "250px",
    maxWidth: "400px",
  },
  submitButton: {
    marginTop: "1em",
    maxWidth: "100px",
    marginLeft: "auto",
    marginRight: "auto"
  },
}));

export const LoginForm: FC<Props> = () => {
  const classes = useStyles();

  const [authService] = useState(new AuthService());
  const [notificationService] = useState(new NotificationService());
  const { isAuthenticated, setUser, setAuthenticated } = useContext(AuthContext);
  const nav = useNavigate();

  const configureInterceptors = () => {
    setAxiosInterceptors(axios, () => {
      setAuthenticated(false);
      nav('/');
      sessionStorage.removeItem('jwt-token');
    });
  }

  useEffect(() => {
    if(isAuthenticated) {
      nav("/");
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated]);

  const schema = Yup.object().shape({
    email: Yup.string()
      .required(() => ({ email: "Email must be provided." }))
      .matches(EMAIL_REGEX, () => ({ 
        email: "Must be a valid email." 
      })),
    password: Yup.string()
      .required(() => ({ password: "Password must be provided." }))
      .matches(LOWER_CASE_REGEX, () => ({
        password: "Password must contain lower letters.",
      }))
      .matches(NUMERIC_REGEX, () => ({
        password: "Password must contain numbers.",
      }))
      .matches(SPECIAL_CHARACTERS_REGEX, () => ({
        password: "Password must contain special characters.",
      })),
  });


  const handleAuthResponse = (response: AuthResponse) => {
    sessionStorage.setItem("jwt-token", response.Jwt);
    setAuthenticated(true);
    setUser(response.Profile);

    configureInterceptors();
  };

  return (
    <Container className={classes.outerContainer}>
      <Container className={classes.container}>
        <Card title="Login">
          <Form
            schema={schema}
            onSubmit={(values) => {
              authService
                .login(values.email, values.password)
                .then(handleAuthResponse)
                .catch((error: AxiosError) => notificationService.error(extractErrorMessage(error.response?.data)));
            }}
          >
            <Container>
              <Col>
                <FormTextInput label="Email" name="email" type="text" />
              </Col>
              <Col>
                <FormTextInput
                  label="Password"
                  name="password"
                  type="password"
                />
              </Col>
              <Button className={classes.submitButton} type="submit">
                Login
              </Button>
            </Container>
          </Form>
        </Card>
      </Container>
    </Container>
  );
};
