import Http from "../http/http";

const RegisterService = (data: any) => {
  return Http.post("/api/auth/register", data)
    .then((res: any) => {
      console.log(data);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const LoginService = (data: any) => {
  return Http.post("api/auth/login", data)
    .then((res: any) => {
      const token = res.data["token"];
      
      sessionStorage.setItem("token", token)
      window.location.href = "/";
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const AuthService = { RegisterService, LoginService };

export default AuthService;
