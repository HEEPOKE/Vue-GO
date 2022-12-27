import Http from "../http/http";
import RegisterModel from "../models/auth/register";

const RegisterService = (data: any) => {
  return Http.post("/api/auth/register", data)
    .then((res: any) => {
      console.log(data);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const AuthService = { RegisterService };

export default AuthService;
