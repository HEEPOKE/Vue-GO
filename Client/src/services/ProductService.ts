import Http from "../http/http";
import authHeader from '../http/token';
import SwalCommon from "./SwalCommon";

const ReadProduct = () => {
  return Http.get("api/product/read", { headers: authHeader() });
};

const GetProductById = (id: number) => {
  return Http.get(`api/product/get/${id}`, { headers: authHeader() });
};

const AddProduct = (data: any) => {
  return Http.post("/api/product/add", data, { headers: authHeader() })
    .then((res: any) => {
      window.location.reload();
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const UpdateProduct = (id: number, data: any) => {
  return Http.put(`/api/product/update/${id}`, data, { headers: authHeader() });
};

const DeleteProduct = (id: number) => {
  return Http.delete(`/api/product/delete/${id}`, { headers: authHeader() })
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
