import axios from "axios";

const Base_Url = import.meta.env.ENDPOINT_URL;

const Http = axios.create({
  baseURL: Base_Url,
  headers: {
    "Content-type": "application/json",
  },
});

export default Http;
