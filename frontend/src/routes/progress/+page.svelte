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
		<h1 class="text-2xl font-bold">Flashing Your Drive...</h1>
		<h2 class="text-md text-muted-foreground">
			This may take a few minutes. Please do not disconnect your drive.
		</h2>
	</div>

	<div class="flex flex-1 flex-wrap gap-4">
		<Card.Root class="flex-1">
			<Card.Header>
				<Card.Title>Love Moonshot?</Card.Title>
				<Card.Description
					>Support us today, and help us keep Moonshot great for you.</Card.Description
				>
			</Card.Header>
			<Card.Footer class="mt-auto flex gap-2">
				<Button
					variant="outline"
					size="icon"
					onclick={() => Browser.OpenURL('https://github.com/sponsors/fyralabs')}
					class="text-[#EA4AAA]"
				>
					<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
						><title>GitHub Sponsors</title><path
							fill="currentColor"
							d="M17.625 1.499c-2.32 0-4.354 1.203-5.625 3.03-1.271-1.827-3.305-3.03-5.625-3.03C3.129 1.499 0 4.253 0 8.249c0 4.275 3.068 7.847 5.828 10.227a33.14 33.14 0 0 0 5.616 3.876l.028.017.008.003-.001.003c.163.085.342.126.521.125.179.001.358-.041.521-.125l-.001-.003.008-.003.028-.017a33.14 33.14 0 0 0 5.616-3.876C20.932 16.096 24 12.524 24 8.249c0-3.996-3.129-6.75-6.375-6.75zm-.919 15.275a30.766 30.766 0 0 1-4.703 3.316l-.004-.002-.004.002a30.955 30.955 0 0 1-4.703-3.316c-2.677-2.307-5.047-5.298-5.047-8.523 0-2.754 2.121-4.5 4.125-4.5 2.06 0 3.914 1.479 4.544 3.684.143.495.596.797 1.086.796.49.001.943-.302 1.085-.796.63-2.205 2.484-3.684 4.544-3.684 2.004 0 4.125 1.746 4.125 4.5 0 3.225-2.37 6.216-5.048 8.523z"
						/></svg
					>
				</Button>
				<Button
					variant="outline"
					size="icon"
					onclick={() => Browser.OpenURL('https://ko-fi.com/fyralabs')}
					class="text-[#FF6433]"
				>
					<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
						><title>Ko-fi</title><path
							fill="currentColor"
							d="M11.351 2.715c-2.7 0-4.986.025-6.83.26C2.078 3.285 0 5.154 0 8.61c0 3.506.182 6.13 1.585 8.493 1.584 2.701 4.233 4.182 7.662 4.182h.83c4.209 0 6.494-2.234 7.637-4a9.5 9.5 0 0 0 1.091-2.338C21.792 14.688 24 12.22 24 9.208v-.415c0-3.247-2.13-5.507-5.792-5.87-1.558-.156-2.65-.208-6.857-.208m0 1.947c4.208 0 5.09.052 6.571.182 2.624.311 4.13 1.584 4.13 4v.39c0 2.156-1.792 3.844-3.87 3.844h-.935l-.156.649c-.208 1.013-.597 1.818-1.039 2.546-.909 1.428-2.545 3.064-5.922 3.064h-.805c-2.571 0-4.831-.883-6.078-3.195-1.09-2-1.298-4.155-1.298-7.506 0-2.181.857-3.402 3.012-3.714 1.533-.233 3.559-.26 6.39-.26m6.547 2.287c-.416 0-.65.234-.65.546v2.935c0 .311.234.545.65.545 1.324 0 2.051-.754 2.051-2s-.727-2.026-2.052-2.026m-10.39.182c-1.818 0-3.013 1.48-3.013 3.142 0 1.533.858 2.857 1.949 3.897.727.701 1.87 1.429 2.649 1.896a1.47 1.47 0 0 0 1.507 0c.78-.467 1.922-1.195 2.623-1.896 1.117-1.039 1.974-2.364 1.974-3.897 0-1.662-1.247-3.142-3.039-3.142-1.065 0-1.792.545-2.338 1.298-.493-.753-1.246-1.298-2.312-1.298"
						/></svg
					>
				</Button>
				<Button
					variant="outline"
					size="icon"
					onclick={() => Browser.OpenURL('https://liberapay.com/fyra')}
					class="text-[#F6C915]"
				>
					<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"
						><title>Liberapay</title><path
							fill="currentColor"
							d="M2.32 0A2.321 2.321 0 0 0 0 2.32v19.36A2.321 2.321 0 0 0 2.32 24h19.36A2.32 2.32 0 0 0 24 21.68V2.32A2.32 2.32 0 0 0 21.68 0zm9.208 3.98l-2.27 9.405a2.953 2.953 0 0 0-.073.539.853.853 0 0 0 .09.432.7.7 0 0 0 .334.302c.157.077.378.126.661.147l-.49 2.008c-.772 0-1.38-.1-1.82-.3-.441-.203-.757-.477-.947-.826a2.391 2.391 0 0 1-.278-1.2c.005-.452.068-.933.188-1.445l2.074-8.67zm3.9 3.888c.61 0 1.135.092 1.576.277.44.185.802.438 1.085.76.283.32.493.696.629 1.126.136.43.204.89.204 1.379v.001c0 .794-.13 1.52-.392 2.179a5.16 5.16 0 0 1-1.086 1.706 4.84 4.84 0 0 1-1.665 1.118c-.648.267-1.353.4-2.114.4-.37 0-.74-.033-1.11-.098l-.735 2.956H9.403l2.71-11.298c.435-.13.934-.248 1.494-.351a10.045 10.045 0 0 1 1.821-.155zm-.31 2.041a4.67 4.67 0 0 0-.98.098l-1.143 4.752c.185.044.413.065.685.065.425 0 .812-.079 1.16-.237a2.556 2.556 0 0 0 .89-.661c.244-.283.435-.623.571-1.02a4.03 4.03 0 0 0 .204-1.315c0-.468-.104-.865-.31-1.192-.207-.326-.566-.49-1.077-.49z"
						/></svg
					>
				</Button>
			</Card.Footer>
		</Card.Root>
		<Card.Root class="flex-1">
			<Card.Header>
				<Card.Title>Get Involved</Card.Title>
				<Card.Description>Moonshot is a project of Fyra Labs.</Card.Description>
			</Card.Header>
			<Card.Footer class="mt-auto flex gap-2">
				<Button onclick={() => Browser.OpenURL('https://fyralabs.com')}>About us</Button>
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
				<Button
					variant="outline"
					size="icon"
					class="text-[#181717]"
					onclick={() => Browser.OpenURL('https://github.com/fyralabs/moonshot')}
				>
					<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
						><title>GitHub</title><path
							d="M12 .297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385.6.113.82-.258.82-.577 0-.285-.01-1.04-.015-2.04-3.338.724-4.042-1.61-4.042-1.61C4.422 18.07 3.633 17.7 3.633 17.7c-1.087-.744.084-.729.084-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22 0 1.606-.015 2.896-.015 3.286 0 .315.21.69.825.57C20.565 22.092 24 17.592 24 12.297c0-6.627-5.373-12-12-12"
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
