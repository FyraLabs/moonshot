import type { ListDrives, SelectFile } from '$lib/wailsjs/go/main/App';
import { EventsOn } from '$lib/wailsjs/runtime/runtime';

export const appState: {
	file: Awaited<ReturnType<typeof SelectFile>> | null;
	drive: Awaited<ReturnType<typeof ListDrives>>[0] | null;
	bytesWritten: number;
	rate: number;
	finished: boolean;
} = $state({
	file: null,
	drive: null,
	bytesWritten: 0,
	rate: 0,
	finished: false
});

export function resetAppState() {
	appState.file = null;
	appState.drive = null;
	appState.bytesWritten = 0;
	appState.rate = 0;
	appState.finished = false;
}

let bytesWritten = 0;

EventsOn('progress', (data) => {
	appState.bytesWritten += JSON.parse(data).written;
	bytesWritten += JSON.parse(data).written;
});

setInterval(() => {
	appState.rate = bytesWritten;
	bytesWritten = 0;
}, 1000);
