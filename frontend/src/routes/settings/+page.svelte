<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { Button } from '$lib/components/ui/button';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import * as Field from '$lib/components/ui/field';
	import { enableAnalytics, showAllDrives } from '../settings.svelte';
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Settings</h1>
		<h2 class="text-md text-muted-foreground">
			Advanced settings to customize your Moonshot experience
		</h2>
	</div>
	<div class="flex-1">
		<Field.Group>
			<Field.Field orientation="horizontal">
				<Checkbox id="show-all-drives" bind:checked={showAllDrives.current} />
				<Field.Content>
					<Field.Label for="show-all-drives">Show all drives</Field.Label>
					<Field.Description
						>This will show all drives including internal drives, which may be generally dangerous
						to flash.</Field.Description
					>
				</Field.Content>
			</Field.Field>
			<Field.Field orientation="horizontal">
				<Checkbox id="enable-analytics" bind:checked={enableAnalytics.current} />
				<Field.Content>
					<Field.Label for="enable-analytics">Enable privacy preserving analytics</Field.Label>
					<Field.Description
						>We use a self-hosted, completely anonymous analytics system to understand how we can
						improve Moonshot.</Field.Description
					>
				</Field.Content>
			</Field.Field>
		</Field.Group>
	</div>
	<div class="ml-auto flex gap-3">
		<Button class="min-w-28" onclick={() => goto(resolve('/'))}>Back</Button>
	</div>
</div>
