import Http from "../http/http";
import SwalCommon from "./SwalCommon";

const ReadProduct = () => {
  return Http.get("api/product/read");
};

const GetProductById = (id: number) => {
  return Http.get(`api/product/get/${id}`);
};

const AddProduct = (data: any) => {
  return Http.post("/api/product/add", data)
    .then((res: any) => {
      window.location.reload();
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const UpdateProduct = (id: number, data: any) => {
  return Http.put(`/api/product/update/${id}`, data);
};

const DeleteProduct = (id: number) => {
  return Http.delete(`/api/product/delete/${id}`)
    .then((res: any) => {
      SwalCommon.SwalSuccess();
      window.location.reload();
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const ProductServices = {
  ReadProduct,
  GetProductById,
  AddProduct,
  UpdateProduct,
  DeleteProduct,
};

export default ProductServices;
