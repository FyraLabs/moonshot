export namespace block {
	
	export class Partition {
	    name: string;
	    label: string;
	    mount_point: string;
	    size_bytes: number;
	    type: string;
	    read_only: boolean;
	    uuid: string;
	    filesystem_label: string;
	
	    static createFrom(source: any = {}) {
	        return new Partition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.label = source["label"];
	        this.mount_point = source["mount_point"];
	        this.size_bytes = source["size_bytes"];
	        this.type = source["type"];
	        this.read_only = source["read_only"];
	        this.uuid = source["uuid"];
	        this.filesystem_label = source["filesystem_label"];
	    }
	}
	export class Disk {
	    name: string;
	    size_bytes: number;
	    physical_block_size_bytes: number;
	    drive_type: number;
	    removable: boolean;
	    storage_controller: number;
	    bus_path: string;
	    vendor: string;
	    model: string;
	    serial_number: string;
	    wwn: string;
	    wwnNoExtension: string;
	    partitions: Partition[];
	
	    static createFrom(source: any = {}) {
	        return new Disk(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.size_bytes = source["size_bytes"];
	        this.physical_block_size_bytes = source["physical_block_size_bytes"];
	        this.drive_type = source["drive_type"];
	        this.removable = source["removable"];
	        this.storage_controller = source["storage_controller"];
	        this.bus_path = source["bus_path"];
	        this.vendor = source["vendor"];
	        this.model = source["model"];
	        this.serial_number = source["serial_number"];
	        this.wwn = source["wwn"];
	        this.wwnNoExtension = source["wwnNoExtension"];
	        this.partitions = this.convertValues(source["partitions"], Partition);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace main {
	
	export class SourceFile {
	    path: string;
	    basename: string;
	    size: number;
	    validGPT: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SourceFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.basename = source["basename"];
	        this.size = source["size"];
	        this.validGPT = source["validGPT"];
	    }
	}

}

