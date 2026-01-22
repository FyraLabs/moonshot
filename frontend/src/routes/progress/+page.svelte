<script>
	import { Progress } from '$lib/components/ui/progress';
</script>

<div class="flex h-screen flex-col gap-6 p-6">
	<div>
		<h1 class="text-2xl font-bold">Flashing your drive</h1>
		<h2 class="text-md text-muted-foreground">
			This may take a few minutes. Please do not disconnect your drive.
		</h2>
	</div>

	<Progress />
</div>
