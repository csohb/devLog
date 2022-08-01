import Main from "./page/Main.svelte";
import NotFound from "./page/NotFound.svelte";
import About from "./page/about/About.svelte";
import AboutEditor from "./page/about/AboutEditor.svelte";
import Blog from "./page/blog/Blog.svelte";
import Contact from "./page/contact/Contact.svelte";
import Story from "./page/story/Story.svelte";
import Login from "./page/auth/Login.svelte";
import BlogDetail from "./page/blog/BlogDetail.svelte";
import BlogEditor from "./page/blog/BlogEditor.svelte";
import StoryDetail from "./page/story/StoryDetail.svelte";
import StoryEditor from "./page/story/StoryEditor.svelte";

const routes = {
  "/": Main,

  // Using named parameters, with last being optional
  "/about": About,
  "/about/edit": AboutEditor,
  "/blog": Blog,
  "/blog/:id": BlogDetail,
  "/blog/edit/:id": BlogEditor,
  "/contact": Contact,
  "/story": Story,
  "/story/:id": StoryDetail,
  "/story/edit/:id": StoryEditor,
  "/login": Login,

  //   // Wildcard parameter
  //   "/book/*": Book,
  // Catch-all route last
  "*": NotFound,
};

export default routes;
