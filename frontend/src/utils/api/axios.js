import ax from 'axios';

export const create = () => {
	const axios = ax.create({
		baseURL: "/",
		withCredentials: true,
	});
	
	return axios;
};

export const axios = create();