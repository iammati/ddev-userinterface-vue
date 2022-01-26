module.exports = {
    plugins: [
        require('tailwind-bootstrap-grid')({
            containerMaxWidths: { sm: '540px', md: '720px', lg: '960px', xl: '1140px' },
        }),
    ],
};
