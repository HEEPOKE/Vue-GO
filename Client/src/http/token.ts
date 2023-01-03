export default function authHeader() {
  const token = sessionStorage.getItem("token") ?? false;

  if (!token) {
    return { Authorization: "" };
  } else {
    return { Authorization: `Bearer ${token}` };
  }
}
