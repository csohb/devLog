<script lang="ts">
    import { createEventDispatcher, onMount } from 'svelte';
    import {link} from 'svelte-spa-router'

    const dispatcher = createEventDispatcher()

    export let colorClass:boolean = false

    let toggleEl:HTMLElement
    let navbarEl:HTMLElement
    let headerEl:HTMLElement
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

    $: console.log('class:',colorClass)
    
    onMount(()=>{
        dispatcher('headerH', headerEl.clientHeight)
    })

    function headerBackgroundHandler(color: boolean){
       if(headerEl){
        if(color){
            headerEl.classList.add('bg') 
        }else{
            headerEl.classList.remove('bg')
        }
       }
    }

    $: headerBackgroundHandler(colorClass)


</script>

<header class="header" bind:this={headerEl}>
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
