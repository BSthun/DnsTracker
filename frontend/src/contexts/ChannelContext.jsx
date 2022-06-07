import {
	createTheme,
	ThemeProvider,
} from '@mui/material/styles';
import React, {
	createContext,
	useMemo,
} from 'react';
import { create } from '../utils/api/axios.js';

export const ChannelContext = createContext(null);

export const ChannelContextProvider = ({ channel, children }) => {
	const value = useMemo(() => {
		const axios = create();
		axios.defaults.headers.common['X-CH-ID'] = channel.channel_id;
		axios.defaults.headers.common['X-CH-TOKEN'] = channel.channel_token;
		
		return {
			channel: channel,
			axios: axios,
		};
	}, [channel]);
	
	return <ChannelContext.Provider value={value}>{children}</ChannelContext.Provider>;
};
