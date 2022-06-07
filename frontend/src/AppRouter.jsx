import {
	Box,
	Stack,
} from '@mui/material';
import Navbar from './components/navbar/index.jsx';
import Tracker from './pages/tracker/Tracker.jsx';
import React from 'react';
import {
	BrowserRouter,
	Route,
	Navigate,
	Routes,
	Outlet,
} from 'react-router-dom';

const AppRouter = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route
					element={
						<Stack>
							<Navbar />
							<Box marginTop={16} minHeight="calc(100vh - 64px)">
								<Outlet />
							</Box>
						</Stack>
					}
				>
					<Route path="/" element={<Navigate to="/tracker" />} />
					<Route path="/tracker" element={<Tracker />} />
				</Route>
			</Routes>
		</BrowserRouter>
	);
};

export default AppRouter;
