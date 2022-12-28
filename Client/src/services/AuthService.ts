import Http from "../http/http";
import RegisterModel from "../models/auth/register";
import LoginModel from "../models/auth/login";

const RegisterService = (data: any) => {
  return Http.post<RegisterModel>("/api/auth/register", data)
    .then((res: any) => {
      console.log(data);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const LoginService = (data: any) => {
  return Http.post<LoginModel>("api/auth/login", data)
    .then((res: any) => {
      console.log(res);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const AuthService = { RegisterService, LoginService };

export default AuthService;
