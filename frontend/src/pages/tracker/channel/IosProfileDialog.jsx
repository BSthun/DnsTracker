import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
} from '@mui/material';
import React, { useContext } from 'react';
import { ChannelContext } from '../../../contexts/ChannelContext.jsx';

const IosProfileDialog = ({open, close}) => {
	const { channel } = useContext(ChannelContext);
	const url = `${channel.base_url}/api/record/ios-profile?id=${channel.channel_id}&salt=${channel.salt}`
	return (
		<Dialog
			open={open}
			onClose={close}
		>
			<DialogTitle id="alert-dialog-title">
				Session #{channel.channel_id} iOS Profile
			</DialogTitle>
			<DialogContent>
				<img
					src={`https://chart.googleapis.com/chart?cht=qr&chs=250x250&chl=${encodeURIComponent(url)}`}
				/>
			</DialogContent>
			<DialogActions>
				<Button onClick={close} autoFocus>
					Close
				</Button>
			</DialogActions>
		</Dialog>
	);
};

export default IosProfileDialog;
