import {
	Box,
	CircularProgress,
	Container,
	Divider,
	Stack,
	Typography,
} from '@mui/material';
import React, {
	useEffect,
	useState,
} from 'react';
import ChannelFragment from './channel/ChannelFragment.jsx';
import { ChannelContextProvider } from '../../contexts/ChannelContext.jsx';
import ChannelAdd from './components/ChannelAdd.jsx';
import ChannelItem from './components/ChannelItem.jsx';

const Tracker = () => {
	const [channels, setChannels] = useState([]);
	const [currentChannel, setCurrentChannel] = useState(null);
	
	useEffect(() => {
		const lc = localStorage.getItem('channels_v1');
		let channels = JSON.parse(lc) || [];
		setChannels(channels);
	}, []);
	
	const addChannel = (channel) => {
		const c = [...channels, channel];
		setChannels(c);
		localStorage.setItem('channels_v1', JSON.stringify(c));
	};
	
	const terminateChannel = (id) => {
		const c = channels.filter((el) => (el.channel_id !== id));
		setChannels(c);
		localStorage.setItem('channels_v1', JSON.stringify(c));
	};
	
	return (
		<Container
			maxWidth="xl"
			sx={{
				marginY: 2,
				paddingX: { 'sm': 0 },
				overflow: 'hidden',
				boxShadow: '0 5px 15px -3px rgb(0 0 0 / 0.2), 0 4px 6px -4px rgb(0 0 0 / 0.1)',
				borderRadius: 2,
			}}
		>
			<Stack height="calc(100vh - 80px)" direction="row" justifyContent="center">
				<Stack width={300} sx={{ borderRight: '1px solid #dadce0' }} divider={<Divider />}>
					{
						channels.map((channel, index) => (
							<ChannelItem key={index} index={index} channel={channel} setChannel={setCurrentChannel} />
						))
					}
					<ChannelAdd addChannel={addChannel} />
				</Stack>
				<Box flex={1}>
					{
						currentChannel ?
							<ChannelContextProvider
								channel={currentChannel}
							>
								<ChannelFragment terminate={terminateChannel}/>
							</ChannelContextProvider> :
							<Stack justifyContent="center" alignItems="center" height="100%">
								<Typography color="text.secondary">No channel selected</Typography>
							</Stack>
					}
				</Box>
			</Stack>
		</Container>
	);
};

export default Tracker;
