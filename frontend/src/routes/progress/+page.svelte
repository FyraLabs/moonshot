<script>
	import { Progress } from '$lib/components/ui/progress';
	import prettyBytes from 'pretty-bytes';
	import { appState } from '../state.svelte';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Button } from '$lib/components/ui/button';
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
				<Card.Title>Card Title</Card.Title>
				<Card.Description>Card Description</Card.Description>
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
				<Card.Title>Card Title</Card.Title>
				<Card.Description>Card Description</Card.Description>
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
				<Card.Description>Support Moonshot on our GitHub Sponsors!</Card.Description>
			</Card.Header>
			<!-- <Card.Content>
				<p>Card Content</p>
			</Card.Content>
			<Card.Footer>
				<p>Card Footer</p>
			</Card.Footer> -->
		</Card.Root>
	</div>

	<div class="flex gap-4">
		<div class="flex flex-1 flex-col gap-2">
			<p class="text-sm text-muted-foreground">
				{prettyBytes(appState.bytesWritten)} / {prettyBytes(appState.file?.size)} ({prettyBytes(
					appState.rate
				)}/s)
			</p>
			<Progress value={appState.bytesWritten} max={appState.drive?.size_bytes} />
		</div>

		<Button disabled>Finish</Button>
	</div>
</div>
