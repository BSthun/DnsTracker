import {
	createTheme,
	ThemeProvider,
} from '@mui/material/styles';
import React, { useMemo } from 'react';

export const ThemeContextProvider = ({ children }) => {
	const theme = useMemo(
		() =>
			createTheme({
				spacing: 4,
				typography: {
					fontFamily: [
						'Arial',
						'Roboto',
						'sans-serif',
					].join(','),
				},
				shape: {
					borderRadius: 4,
				},
			}),
		[],
	);
	
	return <ThemeProvider theme={theme}>{children}</ThemeProvider>;
};
