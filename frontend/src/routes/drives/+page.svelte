<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button } from '$lib/components/ui/button';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import * as Table from '$lib/components/ui/table/index.js';
	import { ListDrives } from '../../../bindings/moonshot/AppService';
	import { Drive } from '../../../bindings/moonshot/models';
	import prettyBytes from 'pretty-bytes';
	import { onMount } from 'svelte';
	import { appState } from '../state.svelte';
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import { resolve } from '$app/paths';
	import { showAllDrives } from '../settings.svelte';

	let drives: Awaited<ReturnType<typeof ListDrives>> = $state([]);
	let selectableDrives = $derived(drives.filter((d) => showAllDrives.current || d.removable));

	onMount(async () => {
		drives = await ListDrives();
	});

	function validDrive(drive: Drive): boolean {
		return Math.ceil(appState.file!.size / 512) * 512 <= drive.capacity;
	}
</script>

<div class="flex h-screen max-h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Select Your Drive</h1>
		<h2 class="text-md text-muted-foreground">
			This is the drive where the operating system will be flashed. All data will be erased.
		</h2>
	</div>

	<div class="h-0 flex-1 overflow-y-auto">
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head></Table.Head>
					<Table.Head class="w-[100px]">ID</Table.Head>
					<Table.Head>Model</Table.Head>
					<Table.Head>Size</Table.Head>
					<Table.Head></Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each selectableDrives as drive (drive)}
					<Table.Row
						onclick={() => {
							if (!validDrive(drive)) return;
							appState.drive = appState.drive === drive ? null : drive;
						}}
					>
						<Table.Cell
							><Checkbox
								checked={appState.drive?.name === drive.name}
								onCheckedChange={(v) => (appState.drive = v ? drive : null)}
								onclick={(e) => e.stopPropagation()}
								disabled={!validDrive(drive)}
							/></Table.Cell
						>
						<Table.Cell class="font-medium">{drive.name}</Table.Cell>
						<Table.Cell>{drive.model}</Table.Cell>
						<Table.Cell class={!validDrive(drive) ? 'text-destructive' : ''}
							>{prettyBytes(drive.capacity)}</Table.Cell
						>
						<Table.Cell class="text-end">
							{#if !validDrive(drive)}
								<Badge variant="destructive">Insufficient Space</Badge>
							{/if}
							{#if !drive.removable}
								<Badge variant="destructive">Internal Drive</Badge>
							{/if}
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>

	<div class="ml-auto flex gap-3">
		<Button class="min-w-28" onclick={() => goto(resolve('/'))} variant="outline">Back</Button>
		<Button class="min-w-28" disabled={!appState.drive} onclick={() => goto(resolve('/confirm'))}
			>Next</Button
		>
	</div>
</div>
