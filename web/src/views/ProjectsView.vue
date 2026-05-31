<template>
  <div class="animate-fade-in-up">
    <!-- Header -->
    <div class="relative mb-16">
      <div class="absolute -top-20 -right-16 w-64 h-64 rounded-full bg-amber-100/20 dark:bg-amber-900/8 blur-3xl -z-10 pointer-events-none"></div>
      <h1 class="font-display text-4xl font-bold mb-5 text-stone-900 dark:text-stone-100 tracking-tight">Projects</h1>
      <p class="text-stone-500 dark:text-stone-400 leading-relaxed max-w-xl">
        A collection of things I've built, from open&#8209;source tools to personal experiments.
      </p>
    </div>

    <div class="divider"><span>&sect;</span></div>

    <!-- Projects Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
      <a
        v-for="project in projects"
        :key="project.id"
        :href="project.link"
        target="_blank"
        rel="noopener noreferrer"
        class="group flex flex-col justify-between p-6 bg-stone-50/60 dark:bg-stone-900/40 border border-stone-200/60 dark:border-stone-800/60 rounded-xl hover:bg-white dark:hover:bg-stone-900 hover:border-amber-300/60 dark:hover:border-amber-700/40 hover:shadow-lg hover:shadow-stone-200/50 dark:hover:shadow-stone-950/50 transition-all duration-500"
      >
        <div>
          <div class="flex items-center justify-between mb-4">
            <h3 class="font-display text-xl font-semibold text-stone-800 dark:text-stone-200 group-hover:text-amber-700 dark:group-hover:text-amber-500 transition-colors duration-500">
              {{ project.name }}
            </h3>
            <span class="text-stone-300 dark:text-stone-600 group-hover:text-amber-500 dark:group-hover:text-amber-500 transition-all duration-500 group-hover:translate-x-0.5 group-hover:-translate-y-0.5">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                <line x1="7" y1="17" x2="17" y2="7"></line>
                <polyline points="7 7 17 7 17 17"></polyline>
              </svg>
            </span>
          </div>
          <p class="text-stone-500 dark:text-stone-400 text-sm leading-relaxed mb-6">
            {{ project.description }}
          </p>
        </div>

        <div class="flex flex-wrap gap-2">
          <span
            v-for="tech in project.techStack"
            :key="tech"
            class="text-[10px] font-mono tracking-wider uppercase px-2.5 py-1 border border-stone-200 dark:border-stone-700 text-stone-400 dark:text-stone-500 rounded-md bg-transparent"
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

const projects = ref([
  {
    id: 1,
    name: 'PersonalSite',
    description: 'A minimalist personal blog and digital garden built with Golang, Vue 3, and Tailwind CSS.',
    link: '#',
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
