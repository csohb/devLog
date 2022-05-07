import Main from "./page/Main.svelte";
import NotFound from "./page/NotFound.svelte";
import About from "./page/about/About.svelte";
import Blog from "./page/blog/Blog.svelte";
import Contact from "./page/contact/Contact.svelte";
import Story from "./page/story/Story.svelte";

const routes = {
  "/": Main,

  // Using named parameters, with last being optional
  "/about": About,
  "/blog": Blog,
  "/contact": Contact,
  "/story": Story,

  //   // Wildcard parameter
  //   "/book/*": Book,
  // Catch-all route last
  "*": NotFound,
};

export default routes;
