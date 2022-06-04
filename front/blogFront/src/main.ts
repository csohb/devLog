import App from './App.svelte';
import 'swiper/swiper-bundle.css'

const app = new App({
	target: document.body,
	props: {
		name: 'world'
	}
});

export default app;