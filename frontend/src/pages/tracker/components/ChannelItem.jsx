import {
	Avatar,
	ButtonBase,
	Divider,
	Stack,
	Typography,
} from '@mui/material';
import React from 'react';

const ChannelItem = ({ index, channel, setChannel }) => {
	return (
		<ButtonBase
			onClick={() => setChannel(channel)}
			sx={{ height: 64, alignItems: 'center', justifyContent: 'flex-start', padding: 4, gap: 4 }}
		>
			<Avatar>{index + 1}</Avatar>
			<Typography color="text.primary">Session #{channel.channel_id}</Typography>
		</ButtonBase>
	);
};

export default ChannelItem;
