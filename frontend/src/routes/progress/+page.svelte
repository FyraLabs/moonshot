<script>
	import { Progress } from '$lib/components/ui/progress';
	import prettyBytes from 'pretty-bytes';
	import { appState, resetAppState } from '../state.svelte';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Button } from '$lib/components/ui/button';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Flashing your drive</h1>
		<h2 class="text-md text-muted-foreground">
			This may take a few minutes. Please do not disconnect your drive.
		</h2>
	</div>

	<div class="grid flex-1 grid-cols-2 grid-rows-2 gap-4">
		<Card.Root class="row-span-2">
			<Card.Header>
				<Card.Title>Need help?</Card.Title>
				<Card.Description>Join our Discord server!</Card.Description>
			</Card.Header>
			<!-- <Card.Content>
				<p>Card Content</p>
			</Card.Content>
			<Card.Footer>
				<p>Card Footer</p>
			</Card.Footer> -->
			<Card.Footer class="mt-auto flex gap-2">
				<!-- <p>Card Footer</p> -->
				<Button onclick={() => BrowserOpenURL('https://fyralabs.com/discord')}>Join Discord</Button>
			</Card.Footer>
		</Card.Root>
		<Card.Root>
			<Card.Header>
				<Card.Title>Sponsored thing</Card.Title>
				<Card.Description
					>Our sponsor is very sigma. Get $67 off of your Labubu when you use code ohio.</Card.Description
				>
			</Card.Header>
			<!-- <Card.Content>
				<p>Card Content</p>
			</Card.Content>
			<Card.Footer>
				<p>Card Footer</p>
			</Card.Footer> -->
		</Card.Root>
		<Card.Root>
			<Card.Header>
				<Card.Title>Love Moonshot?</Card.Title>
				<Card.Description>Moonshot is a project of Fyra Labs.</Card.Description>
			</Card.Header>
			<!-- <Card.Content>
				<p>Card Content</p>
			</Card.Content> -->
			<Card.Footer class="mt-auto flex gap-2">
				<!-- <p>Card Footer</p> -->
				<Button onclick={() => BrowserOpenURL('https://github.com/sponsors/FyraLabs')}
					>Sponsor</Button
				>
				<Button variant="secondary" onclick={() => BrowserOpenURL('https://fyralabs.com')}
					>About us</Button
				>
			</Card.Footer>
		</Card.Root>
	</div>

	<div class="flex gap-4">
		<div class="flex flex-1 flex-col gap-2">
			<p class="text-sm text-muted-foreground">
				{#if appState.stage === 'flash'}
					Flashing:
				{:else if appState.stage === 'verify'}
					Verifying:
				{/if}
				{prettyBytes(appState.bytesWritten)} / {prettyBytes(appState.file?.size ?? 0)} ({prettyBytes(
					appState.rate
				)}/s)
			</p>
			{#key appState.stage}
				<!-- TODO: This gets really weird at high speeds -->
				<Progress value={(appState.bytesWritten / appState.file?.size) * 100} />
			{/key}
		</div>

		<Button
			onclick={() => {
				resetAppState();
				goto(resolve('/'));
			}}
			disabled={!appState.finished}>Finish</Button
		>
	</div>
</div>
