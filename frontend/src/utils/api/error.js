export class ApiCallError extends Error {
	code;
	snackbarOptions;
	
	constructor(message, code, snackbarOptions) {
		super(message);
		this.message = message;
		this.code = code;
		this.snackbarOptions = snackbarOptions;
	}
}