import Main from "./page/Main.svelte";
import NotFound from "./page/NotFound.svelte";
import About from "./page/about/About.svelte";
import Blog from "./page/blog/Blog.svelte";
import Contact from "./page/contact/Contact.svelte";
import Story from "./page/story/Story.svelte";
import Login from "./page/auth/Login.svelte";
import BlogDetail from "./page/blog/BlogDetail.svelte";

const routes = {
  "/": Main,

  // Using named parameters, with last being optional
  "/about": About,
  "/blog": Blog,
  "/blog/:id": BlogDetail,
  "/contact": Contact,
  "/story": Story,
  "/login": Login,

  //   // Wildcard parameter
  //   "/book/*": Book,
  // Catch-all route last
  "*": NotFound,
};

export default routes;
