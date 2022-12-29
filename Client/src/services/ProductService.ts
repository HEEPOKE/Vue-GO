import Http from "../http/http";
import ProductModel from "../models/product";
import UpdateProductModel from "../models/updateProductModel";

const ReadProduct = () => {
  return Http.get<ProductModel>("api/product/read");
};

const AddProduct = (data: any) => {
  return Http.post<ProductModel>("/api/product/add", data)
    .then((res: any) => {
      console.log(res);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const UpdateProduct = ({ id, data }: UpdateProductModel) => {
  return Http.put<ProductModel>(`/api/product/update${id}`, data)
    .then((res: any) => {
      console.log(res);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const ProductServices = { ReadProduct, AddProduct, UpdateProduct };

export default ProductServices;
