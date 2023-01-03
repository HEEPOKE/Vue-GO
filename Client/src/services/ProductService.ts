import Http from "../http/http";

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

const ProductServices = {
  ReadProduct,
  GetProductById,
  AddProduct,
  UpdateProduct,
};

export default ProductServices;
