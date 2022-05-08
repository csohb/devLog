<script lang="ts">
    import {link} from 'svelte-spa-router'
</script>

<header class="header">
<div class="inner">
    <h1>
        <a href="/">Devlog</a>
    </h1>
    <nav>
        <a href="/about" use:link>About</a>
        <a href="/blog" use:link>Blog</a>
        <a href="/story" use:link>Story</a>
        <a href="/contact" use:link>Contact</a>
    </nav>
</div>
</header>
