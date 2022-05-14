<script lang="ts">
    import {link} from 'svelte-spa-router'

    let toggleEl:HTMLElement
    let navbarEl:HTMLElement

    let isToggle:boolean = false

    function toggleBtn():void{
        isToggle = !isToggle
    
        if(isToggle){
            toggleEl.classList.add('on')
            navbarEl.classList.add('on')
        }else{
            toggleEl.classList.remove('on')
            navbarEl.classList.remove('on')
        }
    }

</script>

<header class="header">
<div class="inner">
    <h1>
        <a href="/">Devlog</a>
    </h1>
    <nav>
        <ul class="navbar-menu" bind:this={navbarEl}>
            <li>
                <a href="/about" use:link>About</a>
            </li>
            <li>
                <a href="/blog" use:link>Blog</a>
            </li>
            <li>
                <a href="/story" use:link>Story</a>
            </li>
            <li>
                <a href="/contact" use:link>Contact</a>
            </li>
        </ul>

        <button class="navbar-toggleBtn" on:click={toggleBtn} bind:this={toggleEl}>
            <span></span><span></span><span></span>
        </button>
    </nav>
</div>
</header>
