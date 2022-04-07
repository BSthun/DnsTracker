import {
	CircularProgress,
	Container,
	Stack,
	Typography,
} from '@mui/material';
import React, {
	useEffect,
	useState,
} from 'react';
import { io } from 'socket.io';

const Tracker = () => {
	const [no, setNo] = useState(null);
	
	useEffect(() => {
		const socket = io();
		
		socket.on('connect', () => {
			console.log(socket.connected); // true
			const engine = socket.io.engine;
			console.log(engine.transport.name);
		});
		
		socket.on('disconnect', () => {
			console.log(socket.connected); // false
		});
		
	}, []);
	
	return (
		<Container maxWidth="xl">
			{
				no === null &&
				<Stack height="calc(100vh - 64px)" alignItems="center" justifyContent="center">
					<CircularProgress />
					<Typography variant="body1" mt={4}>Retrieving DNS Tracking Number...</Typography>
				</Stack>
			}
		</Container>
	);
};

export default Tracker;
