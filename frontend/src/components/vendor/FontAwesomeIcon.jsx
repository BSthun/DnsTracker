import React from 'react';

const FontAwesomeIcon = ({ icon, size = '1em', margin = '0' }) => {
	return (
		<i className={icon} style={{ fontSize: size, margin: margin }} />
	);
};

export default FontAwesomeIcon;
