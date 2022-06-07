import {
	Button,
	CircularProgress,
	IconButton,
	Paper,
	Stack,
	Typography,
} from '@mui/material';
import React, {
	useContext,
	useEffect,
	useState,
} from 'react';
import IosProfileDialog from './IosProfileDialog.jsx';
import { axios } from '../../../utils/api/axios.js';
import { caller } from '../../..//utils/api/caller.js';
import LogItem from '../components/LogItem.jsx';
import { ChannelContext } from '../../../contexts/ChannelContext.jsx';
import {
	Refresh,
} from '@mui/icons-material';

const ChannelFragment = ({ terminate }) => {
	const { channel, axios } = useContext(ChannelContext);
	const [logs, setLogs] = useState([]);
	const [status, setStatus] = useState('refreshing');
	const [dialog, setDialog] = useState(false);
	
	var socket = null;
	
	const refresh = () => {
		setStatus('refreshing');
		
		caller(axios.get('/api/record/query-history'))
			.then((res) => {
				setLogs(res.data.reverse());
				
				socket?.close();
				socket = new WebSocket(`${channel.websocket_url}/ws/log?channel_id=${channel.channel_id}&channel_token=${channel.channel_token}`);
				socket.addEventListener('message', (event) => {
					const message = JSON.parse(event.data);
					if (message.event === 'log/update') {
						setLogs((logs) => (
							[message.payload, ...logs]),
						);
					}
				});
			})
			.catch((err) => {
				alert(err.message);
			})
			.finally(() => {
				setStatus('done');
			});
	};
	
	useEffect(() => {
		refresh();
		return () => {
			socket?.close();
		};
	}, [channel]);
	
	return (
		<Stack paddingY={6} paddingX={8} spacing={2}>
			<Stack
				direction="row"
				alignItems="center"
				justifyContent="space-between"
			>
				<Stack>
					<Typography variant="h6">Session #{channel.channel_id}</Typography>
					<Typography variant="caption" fontFamily="monospace">Token <b>{channel.channel_token}</b>
					</Typography>
					<Typography variant="caption" fontFamily="monospace">URL <b>{channel.doh_url}</b></Typography>
					<Stack direction="row" my={4} spacing={2}>
						<Button variant="outlined" onClick={() => setDialog(true)}>iOS Profile</Button>
						<Button variant="outlined" color="error" onClick={() => terminate(channel.channel_id)}>Terminate
							session</Button>
					</Stack>
				</Stack>
				{
					status === 'refreshing' ?
						<CircularProgress size="20px" sx={{ mr: 2 }} /> :
						<IconButton onClick={refresh}><Refresh /></IconButton>
				}
			</Stack>
			<Stack
				sx={{ border: '1px solid #dadce0' }}
				borderRadius={4}
				height="calc(100vh - 260px)"
				overflow="scroll"
			>
				{
					logs.map((el, index) => (
						<LogItem key={index} log={el} />))
				}
				{
					logs.length === 0 &&
					<Stack justifyContent="center" alignItems="center" height="250px">
						<Typography color="text.secondary">There are no DNS query log yet.</Typography>
					</Stack>
				}
			</Stack>
			<IosProfileDialog open={dialog} close={() => setDialog(false)} />
		</Stack>
	);
};

export default ChannelFragment;
