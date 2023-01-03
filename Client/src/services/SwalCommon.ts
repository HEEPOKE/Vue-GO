import Swal from "sweetalert2";

const SwalSuccess = () => {
  Swal.fire({
    icon: "success",
    title: "Success",
    confirmButtonColor: "#3085d6",
  });
};

const SwalErr = () => {
  Swal.fire({
    icon: "error",
    title: "Error try again later",
    confirmButtonColor: "#d33",
  });
};

const SwalCommon = {
  SwalSuccess,
  SwalErr,
};

export default SwalCommon;
