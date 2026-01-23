import type { ListDrives, SelectFile } from '$lib/wailsjs/go/main/App';
import { EventsOn } from '$lib/wailsjs/runtime/runtime';

export const appState: {
	file: Awaited<ReturnType<typeof SelectFile>> | null;
	drive: Awaited<ReturnType<typeof ListDrives>>[0] | null;
	bytesWritten: number;
	rate: number;
	finished: boolean;
	stage: string;
} = $state({
	file: null,
	drive: null,
	bytesWritten: 0,
	rate: 0,
	finished: false,
	stage: 'flash'
});

export function resetAppState() {
	appState.file = null;
	appState.drive = null;
	appState.bytesWritten = 0;
	appState.rate = 0;
	appState.finished = false;
	appState.stage = 'flash';
}

let bytesWritten = 0;

EventsOn('progress', (data) => {
	const parsedData = JSON.parse(data);
	if (parsedData.stage != appState.stage) {
		appState.stage = parsedData.stage;
		appState.bytesWritten = 0;
		bytesWritten = 0;
	}
	appState.bytesWritten += parsedData.written;
	bytesWritten += parsedData.written;
});

setInterval(() => {
	appState.rate = bytesWritten;
	bytesWritten = 0;
}, 1000);
