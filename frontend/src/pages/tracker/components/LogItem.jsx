import {
	Box,
	Stack,
	Typography,
} from '@mui/material';
import React from 'react';

const LogItem = ({ log }) => {
	return (
		<Stack direction="row" alignItems="center" sx={{"&:hover": { backgroundColor: "#EAEAEA"}}}>
			<Typography
				fontFamily="monospace"
				variant="caption"
				textAlign="center"
				minWidth={256}
			>
				{log.time}
			</Typography>
			<Box
				mb="5px"
				mt="3px"
				width={12}
				height={12}
				borderRadius={6}
				bgcolor={log.status === 0 ? '#00E676' : '#FF3D00'}
			/>
			<Typography ml={3} fontFamily="monospace" variant="caption">{log.hostname}</Typography>
		</Stack>
	);
};

export default LogItem;
