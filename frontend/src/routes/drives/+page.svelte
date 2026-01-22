<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button } from '$lib/components/ui/button';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import * as Table from '$lib/components/ui/table/index.js';
	import { ListDrives } from '$lib/wailsjs/go/main/App';
	import prettyBytes from 'pretty-bytes';
	import { onMount } from 'svelte';
	import { appState } from '../state.svelte';

	let drives: Awaited<ReturnType<typeof ListDrives>> = $state([]);
	let selectableDrives = $derived(drives);
	// let selectableDrives = $derived(drives.filter((d) => d.removable));

	onMount(async () => {
		drives = await ListDrives();
	});
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Select your drive</h1>
		<h2 class="text-md text-muted-foreground">
			This is the drive where the operating system will be flashed. All data will be erased.
		</h2>
	</div>

	<div class="flex-1">
		<Table.Root>
			<!-- <Table.Caption>A list of your recent invoices.</Table.Caption> -->
			<Table.Header>
				<Table.Row>
					<Table.Head></Table.Head>
					<Table.Head class="w-[100px]">ID</Table.Head>
					<Table.Head>Model</Table.Head>
					<Table.Head>Size</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each selectableDrives as drive (drive)}
					<Table.Row>
						<Table.Cell
							><Checkbox
								checked={appState.drive?.name === drive.name}
								onCheckedChange={(v) => (appState.drive = v ? drive : null)}
							/></Table.Cell
						>
						<Table.Cell class="font-medium">{drive.name}</Table.Cell>
						<Table.Cell>{drive.model}</Table.Cell>
						<Table.Cell>{prettyBytes(drive.size_bytes)}</Table.Cell>
						<!-- <Table.Cell class="text-end">$250.00</Table.Cell> -->
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>

	<div class="ml-auto flex gap-3">
		<Button class="min-w-28" onclick={() => goto('/')} variant="outline">Back</Button>
		<Button class="min-w-28" disabled={!appState.drive} onclick={() => goto('/confirm')}
			>Next</Button
		>
	</div>
</div>
