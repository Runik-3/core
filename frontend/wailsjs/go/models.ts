export namespace core {
	
	export class Config {
	    kindlegenPath: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.kindlegenPath = source["kindlegenPath"];
	    }
	}
	export class Response[[]github.com/runik-3/core/core.File] {
	
	
	    static createFrom(source: any = {}) {
	        return new Response[[]github.com/runik-3/core/core.File](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Response[github.com/runik-3/builder/dict.Dict] {
	
	
	    static createFrom(source: any = {}) {
	        return new Response[github.com/runik-3/builder/dict.Dict](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Response[github.com/runik-3/builder/wikiBot.WikiDetails] {
	
	
	    static createFrom(source: any = {}) {
	        return new Response[github.com/runik-3/builder/wikiBot.WikiDetails](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Response[github.com/runik-3/core/device.Device] {
	
	
	    static createFrom(source: any = {}) {
	        return new Response[github.com/runik-3/core/device.Device](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Response[string] {
	
	
	    static createFrom(source: any = {}) {
	        return new Response[string](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

export namespace dict {
	
	export class Dict {
	
	
	    static createFrom(source: any = {}) {
	        return new Dict(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

