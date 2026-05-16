<template>
  <div class="animate-fade-in-up">
    <!-- Header -->
    <div class="mb-16">
      <h1 class="text-3xl font-bold mb-4 text-zinc-900 tracking-tight">Projects</h1>
      <p class="text-zinc-500 font-light leading-relaxed max-w-2xl text-base">
        A collection of things I've built, ranging from open-source tools to personal experiments. 
        Click on any project to explore the code or live demo.
      </p>
    </div>

    <!-- Projects Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <a 
        v-for="project in projects" 
        :key="project.id" 
        :href="project.link" 
        target="_blank"
        class="group flex flex-col justify-between p-6 border border-zinc-100 rounded-2xl hover:border-zinc-300 hover:shadow-sm hover:bg-zinc-50/50 transition duration-300"
      >
        <div>
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-xl font-medium text-zinc-800 group-hover:text-zinc-900 transition-colors">
              {{ project.name }}
            </h3>
            <span class="text-zinc-300 group-hover:text-zinc-500 transition-colors">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <line x1="7" y1="17" x2="17" y2="7"></line>
                <polyline points="7 7 17 7 17 17"></polyline>
              </svg>
            </span>
          </div>
          <p class="text-zinc-500 text-sm font-light leading-relaxed mb-6">
            {{ project.description }}
          </p>
        </div>
        
        <div class="flex flex-wrap gap-2">
          <span 
            v-for="tech in project.techStack" 
            :key="tech" 
            class="text-[11px] font-mono tracking-wider uppercase px-2 py-0.5 border border-zinc-200 text-zinc-500 rounded-md bg-white"
          >
            {{ tech }}
          </span>
        </div>
      </a>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// Since there's no backend API for projects yet, we use a static list
const projects = ref([
  {
    id: 1,
    name: 'PersonalSite',
    description: 'A minimalist personal blog and digital garden built with Golang, Vue 3, and Tailwind CSS. Features an elegant writing experience.',
    link: '#', // Add your Github link here
    techStack: ['Go', 'Vue', 'Tailwind']
  },
  {
    id: 2,
    name: 'Terminal-Chat',
    description: 'A command-line based chat application that allows developers to communicate securely over local networks.',
    link: '#', 
    techStack: ['Rust', 'WebSocket']
  },
  {
    id: 3,
    name: 'AI Code Reviewer',
    description: 'A GitHub Action that automatically reviews pull requests using LLMs, leaving comments on potential bugs and style issues.',
    link: '#',
    techStack: ['Python', 'OpenAI API']
  },
  {
    id: 4,
    name: 'UI Component Library',
    description: 'An open-source set of accessible, highly customizable Vue components designed for modern dashboards.',
    link: '#',
    techStack: ['Vue 3', 'TypeScript']
  }
])
</script>