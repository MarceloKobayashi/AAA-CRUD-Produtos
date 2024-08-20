import { createRouter, createWebHistory } from 'vue-router';
import InitialPage from '../views/InitialPage.vue';
import InsertProduct from '../views/InsertProduct.vue';
import ProductsDashboard from '../views/ProductsDashboard.vue';
import DeleteProduct from '../views/DeleteProduct.vue';
import UpdateProduct from '../views/UpdateProduct.vue';

const routes = [
    {
      path: '/',
      name: 'initial',
      component: InitialPage
    },
    {
      path: '/produto',
      name: 'insert product',
      component: InsertProduct
    },
    {
      path: '/produtos',
      name: 'show products',
      component: ProductsDashboard
    },
    {
      path: '/produto/deletar/:name',
      name: 'deleteProduct',
      component: DeleteProduct
    },
    {
      path: '/produto/alterar/:name',
      name: 'updateProduct',
      component: UpdateProduct
    }
  ];

  const router = createRouter({
    history: createWebHistory(),
    routes
  });

export default router;
