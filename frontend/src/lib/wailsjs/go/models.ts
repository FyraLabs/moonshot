export namespace main {
	
	export class Drive {
	    name: string;
	    model: string;
	    capacity: number;
	    path: string;
	    removable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Drive(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.model = source["model"];
	        this.capacity = source["capacity"];
	        this.path = source["path"];
	        this.removable = source["removable"];
	    }
	}
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

