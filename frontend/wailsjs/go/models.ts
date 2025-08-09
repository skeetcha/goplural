export namespace themes {
	
	export class ThemeColor {
	    R: number;
	    G: number;
	    B: number;
	    A: number;
	
	    static createFrom(source: any = {}) {
	        return new ThemeColor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.R = source["R"];
	        this.G = source["G"];
	        this.B = source["B"];
	        this.A = source["A"];
	    }
	}
	export class ThemeColors {
	    primary: ThemeColor;
	    secondary: ThemeColor;
	    success: ThemeColor;
	    info: ThemeColor;
	    warning: ThemeColor;
	    danger: ThemeColor;
	    light: ThemeColor;
	    dark: ThemeColor;
	    bg: ThemeColor;
	    fg: ThemeColor;
	    selectbg: ThemeColor;
	    selectfg: ThemeColor;
	    border: ThemeColor;
	    inputfg: ThemeColor;
	    inputbg: ThemeColor;
	
	    static createFrom(source: any = {}) {
	        return new ThemeColors(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.primary = this.convertValues(source["primary"], ThemeColor);
	        this.secondary = this.convertValues(source["secondary"], ThemeColor);
	        this.success = this.convertValues(source["success"], ThemeColor);
	        this.info = this.convertValues(source["info"], ThemeColor);
	        this.warning = this.convertValues(source["warning"], ThemeColor);
	        this.danger = this.convertValues(source["danger"], ThemeColor);
	        this.light = this.convertValues(source["light"], ThemeColor);
	        this.dark = this.convertValues(source["dark"], ThemeColor);
	        this.bg = this.convertValues(source["bg"], ThemeColor);
	        this.fg = this.convertValues(source["fg"], ThemeColor);
	        this.selectbg = this.convertValues(source["selectbg"], ThemeColor);
	        this.selectfg = this.convertValues(source["selectfg"], ThemeColor);
	        this.border = this.convertValues(source["border"], ThemeColor);
	        this.inputfg = this.convertValues(source["inputfg"], ThemeColor);
	        this.inputbg = this.convertValues(source["inputbg"], ThemeColor);
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
	export class Theme {
	    type: number;
	    colors: ThemeColors;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Theme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.colors = this.convertValues(source["colors"], ThemeColors);
	        this.name = source["name"];
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

