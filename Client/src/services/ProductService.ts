import Http from "../http/http";
import ProductModel from "../models/product";

const AddProduct = (data: any) => {
  return Http.post<ProductModel>("/api/product/add", data)
    .then((res: any) => {
      console.log(res);
    })
    .catch((err: any) => {
      console.log(err);
    });
};

const ProductServices = { AddProduct };

export default ProductServices;
