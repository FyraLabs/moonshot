<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button } from '$lib/components/ui/button';
	import { ArrowRight, HardDrive, File, ArrowDown } from '@lucide/svelte';
	import { appState } from '../state.svelte';
	import prettyBytes from 'pretty-bytes';
	import { FlashDrive } from '$lib/wailsjs/go/main/App';
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Confirm selection</h1>
		<h2 class="text-md text-muted-foreground">
			Are you sure you want to flash this drive? All data will be erased.
		</h2>
	</div>

	<div class="flex flex-1 flex-col items-center justify-center gap-6">
		<div class="flex flex-col items-center gap-2">
			<File />
			<p>{appState.file.basename} ({prettyBytes(appState.file.size)})</p>
		</div>

		<ArrowDown />

		<div class="flex flex-col items-center gap-2">
			<HardDrive />
			<p>
				{appState.drive.model} ({appState.drive?.name}, {prettyBytes(appState.drive?.size_bytes)})
			</p>
		</div>
	</div>

	<div class="ml-auto flex gap-3">
		<Button class="min-w-28" onclick={() => goto('/drives')} variant="outline">Back</Button>
		<Button
			class="min-w-28"
			disabled={false}
			onclick={() => {
				FlashDrive(appState.file.path, '/dev/rdisk9');
				goto('/progress');
			}}
			variant="destructive">Confirm</Button
		>
	</div>
</div>
