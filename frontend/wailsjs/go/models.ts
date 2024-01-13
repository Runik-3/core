export namespace wikibot {
	
	export class Lang {
	    lang: string;
	    langname: string;
	    autonym: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new Lang(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lang = source["lang"];
	        this.langname = source["langname"];
	        this.autonym = source["autonym"];
	        this.url = source["url"];
	    }
	}

}

