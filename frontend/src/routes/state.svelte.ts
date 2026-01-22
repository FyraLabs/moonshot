import type { ListDrives, SelectFile } from '$lib/wailsjs/go/main/App';

export const appState: {
	file: Awaited<ReturnType<typeof SelectFile>> | null;
	drive: Awaited<ReturnType<typeof ListDrives>>[0] | null;
	bytesWritten: number;
} = $state({
	file: null,
	drive: null,
	bytesWritten: 0
});
