import { ApiCallError } from './error.js';
import { AxiosResponse } from 'axios';

export const caller = function (c, options) {
	return new Promise((resolve, reject) => {
		c.then(({ data, statusText }) => {
			if (data.success === true) {
				resolve(data);
			}
			if (data.success === false && data.message) {
				reject(
					new ApiCallError(data.message, data.code, { variant: 'warning' }),
				);
			}
			
			// In case of no data value
			reject(
				new ApiCallError(statusText, 'NO_DATA_ERROR', { variant: 'error' }));
		}).catch((e) => {
			reject(new ApiCallError(e.message, 'AXIOS_ERROR', { variant: 'error' }));
		});
	});
};
