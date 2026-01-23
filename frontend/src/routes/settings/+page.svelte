<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { Button } from '$lib/components/ui/button';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import * as Field from '$lib/components/ui/field';
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Settings</h1>
		<h2 class="text-md text-muted-foreground">Settings for Moonshot.</h2>
	</div>
	<div class="flex-1">
		<Field.Group>
			<Field.Field orientation="horizontal">
				<Checkbox id="finder-pref-9k2-sync-folders-nep" checked />
				<Field.Content>
					<Field.Label for="finder-pref-9k2-sync-folders-nep">Show all drives</Field.Label>
					<Field.Description
						>This will show all drives including internal drives, which may be generally dangerous
						to flash.</Field.Description
					>
				</Field.Content>
			</Field.Field>
			<Field.Field orientation="horizontal">
				<Checkbox id="finder-pref-9k2-sync-folders-nep" checked />
				<Field.Content>
					<Field.Label for="finder-pref-9k2-sync-folders-nep"
						>Disable privacy preseving analytics</Field.Label
					>
					<Field.Description>We use these .</Field.Description>
				</Field.Content>
			</Field.Field>
		</Field.Group>
	</div>
	<div class="ml-auto flex gap-3">
		<Button class="min-w-28" onclick={() => goto(resolve('/'))}>Back</Button>
	</div>
</div>
