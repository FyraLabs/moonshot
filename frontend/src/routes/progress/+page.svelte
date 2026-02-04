<script lang="ts">
	import { Progress } from '$lib/components/ui/progress';
	import prettyBytes from 'pretty-bytes';
	import { appState, resetAppState } from '../state.svelte';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Button } from '$lib/components/ui/button';
	import { Browser } from '@wailsio/runtime';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { Throttled } from 'runed';

	const throttledbytesWritten = new Throttled(() => appState.bytesWritten, 100);
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Flashing your drive</h1>
		<h2 class="text-md text-muted-foreground">
			This may take a few minutes. Please do not disconnect your drive.
		</h2>
	</div>

	<div class="flex flex-1 flex-wrap gap-4">
		<Card.Root class="flex-1">
			<Card.Header>
				<Card.Title>Sponsored thing</Card.Title>
				<Card.Description
					>Our sponsor is very sigma. Get $67 off of your Labubu when you use code ohio.</Card.Description
				>
			</Card.Header>
		</Card.Root>
		<Card.Root class="flex-1">
			<Card.Header>
				<Card.Title>Love Moonshot?</Card.Title>
				<Card.Description>Moonshot is a project of Fyra Labs.</Card.Description>
			</Card.Header>
			<Card.Footer class="mt-auto flex gap-2">
				<Button onclick={() => Browser.OpenURL('https://github.com/sponsors/FyraLabs')}
					>Sponsor</Button
				>
				<Button variant="secondary" onclick={() => Browser.OpenURL('https://fyralabs.com')}
					>About us</Button
				>
				<Button
					variant="outline"
					size="icon"
					class="text-[#5865F2]"
					onclick={() => Browser.OpenURL('https://fyralabs.com/discord')}
				>
					<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
						><path
							fill="currentColor"
							d="M20.317 4.3698a19.7913 19.7913 0 00-4.8851-1.5152.0741.0741 0 00-.0785.0371c-.211.3753-.4447.8648-.6083 1.2495-1.8447-.2762-3.68-.2762-5.4868 0-.1636-.3933-.4058-.8742-.6177-1.2495a.077.077 0 00-.0785-.037 19.7363 19.7363 0 00-4.8852 1.515.0699.0699 0 00-.0321.0277C.5334 9.0458-.319 13.5799.0992 18.0578a.0824.0824 0 00.0312.0561c2.0528 1.5076 4.0413 2.4228 5.9929 3.0294a.0777.0777 0 00.0842-.0276c.4616-.6304.8731-1.2952 1.226-1.9942a.076.076 0 00-.0416-.1057c-.6528-.2476-1.2743-.5495-1.8722-.8923a.077.077 0 01-.0076-.1277c.1258-.0943.2517-.1923.3718-.2914a.0743.0743 0 01.0776-.0105c3.9278 1.7933 8.18 1.7933 12.0614 0a.0739.0739 0 01.0785.0095c.1202.099.246.1981.3728.2924a.077.077 0 01-.0066.1276 12.2986 12.2986 0 01-1.873.8914.0766.0766 0 00-.0407.1067c.3604.698.7719 1.3628 1.225 1.9932a.076.076 0 00.0842.0286c1.961-.6067 3.9495-1.5219 6.0023-3.0294a.077.077 0 00.0313-.0552c.5004-5.177-.8382-9.6739-3.5485-13.6604a.061.061 0 00-.0312-.0286zM8.02 15.3312c-1.1825 0-2.1569-1.0857-2.1569-2.419 0-1.3332.9555-2.4189 2.157-2.4189 1.2108 0 2.1757 1.0952 2.1568 2.419 0 1.3332-.9555 2.4189-2.1569 2.4189zm7.9748 0c-1.1825 0-2.1569-1.0857-2.1569-2.419 0-1.3332.9554-2.4189 2.1569-2.4189 1.2108 0 2.1757 1.0952 2.1568 2.419 0 1.3332-.946 2.4189-2.1568 2.4189Z"
						/></svg
					>
				</Button>
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
			<Progress value={throttledbytesWritten.current} max={appState.file?.size} />
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
