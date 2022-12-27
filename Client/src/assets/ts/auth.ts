import Swal from "sweetalert2";

const checkPassword = () => {
  Swal.fire({
    icon: "error",
    title: "รหัสผ่านไม่ตรงกัน",
    text: "กรุณากรอกใหม่",
    showConfirmButton: false,
    showDenyButton: true,
    denyButtonText: "ตกลง",
  });
};

const swalAuth = { checkPassword };

export default swalAuth;
