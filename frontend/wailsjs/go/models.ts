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
	export class File {
	    Name: string;
	    Display: string;
	    Extension: string;
	    Size: number;
	    // Go type: time
	    Modified: any;
	
	    static createFrom(source: any = {}) {
	        return new File(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Display = source["Display"];
	        this.Extension = source["Extension"];
	        this.Size = source["Size"];
	        this.Modified = this.convertValues(source["Modified"], null);
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
	export class Response___github_com_runik_3_core_core_File_ {
	    Data: File[];
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Response___github_com_runik_3_core_core_File_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = this.convertValues(source["Data"], File);
	        this.Error = source["Error"];
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
	export class Response_github_com_runik_3_builder_dict_Dict_ {
	    Data: dict.Dict;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Response_github_com_runik_3_builder_dict_Dict_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = this.convertValues(source["Data"], dict.Dict);
	        this.Error = source["Error"];
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
	export class Response_github_com_runik_3_builder_wikiBot_WikiDetails_ {
	    Data: wikiBot.WikiDetails;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Response_github_com_runik_3_builder_wikiBot_WikiDetails_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = this.convertValues(source["Data"], wikiBot.WikiDetails);
	        this.Error = source["Error"];
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
	export class Response_github_com_runik_3_core_core_Config_ {
	    Data: Config;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Response_github_com_runik_3_core_core_Config_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = this.convertValues(source["Data"], Config);
	        this.Error = source["Error"];
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
	export class Response_github_com_runik_3_core_device_Device_ {
	    Data: any;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Response_github_com_runik_3_core_device_Device_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = source["Data"];
	        this.Error = source["Error"];
	    }
	}
	export class Response_string_ {
	    Data: string;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new Response_string_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = source["Data"];
	        this.Error = source["Error"];
	    }
	}

}

export namespace dict {
	
	export class Entry {
	    Word: string;
	    Definition: string;
	    Synonyms: string[];
	
	    static createFrom(source: any = {}) {
	        return new Entry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Word = source["Word"];
	        this.Definition = source["Definition"];
	        this.Synonyms = source["Synonyms"];
	    }
	}
	export class Dict {
	    Name: string;
	    ApiUrl: string;
	    Lang: string;
	    Lexicon: Entry[];
	
	    static createFrom(source: any = {}) {
	        return new Dict(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.ApiUrl = source["ApiUrl"];
	        this.Lang = source["Lang"];
	        this.Lexicon = this.convertValues(source["Lexicon"], Entry);
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

export namespace frontend {
	
	export class FileFilter {
	    DisplayName: string;
	    Pattern: string;
	
	    static createFrom(source: any = {}) {
	        return new FileFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DisplayName = source["DisplayName"];
	        this.Pattern = source["Pattern"];
	    }
	}
	export class OpenDialogOptions {
	    DefaultDirectory: string;
	    DefaultFilename: string;
	    Title: string;
	    Filters: FileFilter[];
	    ShowHiddenFiles: boolean;
	    CanCreateDirectories: boolean;
	    ResolvesAliases: boolean;
	    TreatPackagesAsDirectories: boolean;
	
	    static createFrom(source: any = {}) {
	        return new OpenDialogOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DefaultDirectory = source["DefaultDirectory"];
	        this.DefaultFilename = source["DefaultFilename"];
	        this.Title = source["Title"];
	        this.Filters = this.convertValues(source["Filters"], FileFilter);
	        this.ShowHiddenFiles = source["ShowHiddenFiles"];
	        this.CanCreateDirectories = source["CanCreateDirectories"];
	        this.ResolvesAliases = source["ResolvesAliases"];
	        this.TreatPackagesAsDirectories = source["TreatPackagesAsDirectories"];
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

export namespace utils {
	
	export class GetRequestOptions {
	    ForceTLS12: boolean;
	    Timeout: number;
	
	    static createFrom(source: any = {}) {
	        return new GetRequestOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ForceTLS12 = source["ForceTLS12"];
	        this.Timeout = source["Timeout"];
	    }
	}

}

export namespace wikiBot {
	
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
	export class WikiDetails {
	    ApiUrl: string;
	    SiteName: string;
	    MainPage: string;
	    Lang: string;
	    Logo: string;
	    Pages: number;
	    Articles: number;
	    Languages: Lang[];
	    RequestOpts: utils.GetRequestOptions;
	
	    static createFrom(source: any = {}) {
	        return new WikiDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ApiUrl = source["ApiUrl"];
	        this.SiteName = source["SiteName"];
	        this.MainPage = source["MainPage"];
	        this.Lang = source["Lang"];
	        this.Logo = source["Logo"];
	        this.Pages = source["Pages"];
	        this.Articles = source["Articles"];
	        this.Languages = this.convertValues(source["Languages"], Lang);
	        this.RequestOpts = this.convertValues(source["RequestOpts"], utils.GetRequestOptions);
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

