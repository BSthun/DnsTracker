import {
	Box,
	Container,
	Paper,
	Stack,
	Typography,
} from '@mui/material';
import FontAwesomeIcon from '../vendor/FontAwesomeIcon.jsx';
import React from 'react';

const Navbar = () => {
	return (
		<Paper square elevation={4} sx={{ position: 'fixed', inset: 'auto 0' }}>
			<Container maxWidth="xl">
				<Stack height={64} flexDirection="row" alignItems="center" justifyContent="space-between">
					<Box display="flex" alignItems="center" gap={2}>
						<FontAwesomeIcon icon="fa-solid fa-server" size="1.2em" />
						<Typography variant="h5">DNS Tracker</Typography>
					</Box>
					<Box textAlign="end">
						<Typography variant="body2">
							Developed by&nbsp;
							<a target="_blank" href="https://github.com/BSthun">
								BSthun
							</a>
						</Typography>
						<Typography variant="body2">
							Source code is available on
							<a target="_blank" href="https://github.com/BSthun/DnsTracker">
								<FontAwesomeIcon icon="fa-brands fa-github" margin="0 4px" />GitHub
							</a>
						</Typography>
					</Box>
				</Stack>
			</Container>
		</Paper>
	);
};

export default Navbar;
