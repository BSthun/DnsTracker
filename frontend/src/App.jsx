import React from 'react';
import AppRouter from './AppRouter.jsx';
import { ThemeContextProvider } from './contexts/ThemeContext.jsx';
import './styles/index.scss';

const App = () => {
	return (
		<ThemeContextProvider>
			<AppRouter />
		</ThemeContextProvider>
	);
};

export default App;
