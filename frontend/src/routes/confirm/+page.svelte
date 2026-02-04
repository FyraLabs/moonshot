<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button } from '$lib/components/ui/button';
	import { HardDrive, File, ArrowDown } from '@lucide/svelte';
	import { appState } from '../state.svelte';
	import prettyBytes from 'pretty-bytes';
	import { FlashDrive } from '../../../bindings/moonshot/appservice';
	import { resolve } from '$app/paths';
	import { toast } from 'svelte-sonner';

	const file = $derived(appState.file!);
	const drive = $derived(appState.drive!);
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Confirm Selection</h1>
		<h2 class="text-md text-muted-foreground">
			Are you sure you want to flash this drive? All data will be erased.
		</h2>
	</div>

	<div class="flex flex-1 flex-col items-center justify-center gap-6">
		<div class="flex flex-col items-center gap-2">
			<File />
			<p>{file.basename} ({prettyBytes(file.size)})</p>
		</div>

		<ArrowDown />

		<div class="flex flex-col items-center gap-2">
			<HardDrive />
			<p>
				{drive.model} ({drive.name}, {prettyBytes(drive.capacity)})
			</p>
		</div>
	</div>

	<div class="ml-auto flex gap-3">
		<Button class="min-w-28" onclick={() => goto(resolve('/drives'))} variant="outline">Back</Button
		>
		<Button
			class="min-w-28"
			disabled={false}
			onclick={() => {
				FlashDrive(file.path, drive.name, drive.removable)
					.catch((e) =>
						toast.error(`Error flashing drive: ${e.message ?? e}`, {
							richColors: true,
							duration: Infinity,
							closeButton: true
						})
					)
					.finally(() => {
						appState.finished = true;
					});
				goto(resolve('/progress'));
			}}
			variant="destructive">Confirm</Button
		>
	</div>
</div>
