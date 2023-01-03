import Swal from "sweetalert2";
import SwalCommon from "../SwalCommon";
import ProductServices from "../ProductService";

const DeleteSwal = (id: number, name: string) => {
  Swal.fire({
    title: "Are you sure?",
    text: `confirm delete ${name}`,
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: "Yes, delete it!",
  }).then((result) => {
    if (result.isConfirmed) {
      ProductServices.DeleteProduct(id);
    } else {
      SwalCommon.SwalErr();
    }
  });
};

const SwalProduct = {
  DeleteSwal,
};

export default SwalProduct;
