import {
	Avatar,
	ButtonBase,
	CircularProgress,
	Typography,
} from '@mui/material';
import { deepOrange } from '@mui/material/colors';
import React, { useState } from 'react';
import { axios } from '../../../utils/api/axios.js';
import { caller } from '../../../utils/api/caller.js';

const ChannelAdd = ({ addChannel }) => {
	const [loading, setLoading] = useState(false);
	
	const newSession = () => {
		setLoading(true);
		caller(axios.post('/api/account/session/new'))
			.then((res) => {
				addChannel(res.data);
			})
			.catch((err) => {
				alert(err.message);
			})
			.finally(() => {
				setLoading(false);
			});
	};
	
	return (
		<ButtonBase
			onClick={() => newSession()}
			sx={{ height: 64, alignItems: 'center', justifyContent: 'flex-start', padding: 4, gap: 4 }}
		>
			{
				loading ?
					<CircularProgress /> :
					<Avatar sx={{ bgcolor: '#E3F2FD', color: '#2196F3' }}>+</Avatar>
			}
			<Typography color="text.primary">New session</Typography>
		</ButtonBase>
	);
};

export default ChannelAdd;
