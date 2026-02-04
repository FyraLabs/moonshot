<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { SelectFile } from '../../bindings/moonshot/AppService';
	import { Events, Window } from '@wailsio/runtime';
	import { onMount } from 'svelte';
	import prettyBytes from 'pretty-bytes';
	import { Upload, File, CircleAlert } from '@lucide/svelte';
	import { goto } from '$app/navigation';
	import { appState } from './state.svelte';
	import * as Alert from '$lib/components/ui/alert/index.js';
	import { resolve } from '$app/paths';

	onMount(() => {
		Events.On('files-dropped', async ({ data: { files } }) => {
			if (appState.file) return;
			if (files.length > 0) {
				const imagePath = files[0];
				if (!/^.+(\.raw|\.iso|\.img)$/.test(imagePath)) return;
				appState.file = await SelectFile(files[0]);
			}
		});

		return () => Events.Off('files-dropped');
	});

	async function selectImage() {
		if (appState.file) return;
		appState.file = await SelectFile(null);
	}
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Select your image</h1>
		<h2 class="text-md text-muted-foreground">
			This is the operating system to flash to your drive.
		</h2>
	</div>

	{#if appState.file && !appState.file.validGPT}
		<Alert.Root variant="destructive">
			<CircleAlert />
			<Alert.Title>Potentially Invalid Image</Alert.Title>
			<Alert.Description
				>The image doesn't seem to be a valid GPT disk. Your operating system will probably not
				boot.</Alert.Description
			>
		</Alert.Root>
	{/if}

	<button
		type="button"
		class="flex w-full flex-1 flex-col items-center justify-center rounded border-2 border-dotted border-muted [.file-drop-target-active]:border-primary"
		onclick={selectImage}
		data-file-drop-target
	>
		<div class="flex flex-col items-center gap-2">
			{#if appState.file}
				<File />
				<p>{appState.file.basename} ({prettyBytes(appState.file.size)})</p>
				<Button
					variant="outline"
					onclick={(e) => {
						e.stopPropagation();
						appState.file = null;
					}}>Clear</Button
				>
			{:else}
				<Upload />
				<p>Drag and drop or <span>click to select a file</span></p>
			{/if}
		</div>
	</button>

	<div class="flex gap-3">
		<Button class="min-w-28" variant="outline" onclick={() => goto(resolve('/settings'))}
			>Settings</Button
		>
		<Button
			class="ml-auto min-w-28"
			disabled={!appState.file}
			onclick={() => goto(resolve('/drives'))}>Next</Button
		>
	</div>
</div>
