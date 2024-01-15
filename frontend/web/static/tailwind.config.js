/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './out/*.{html,js}',
        '../templates/*.templ'
    ],
    theme: {
        extend: {
            maxWidth: {
                '1/2': '50%',
            },
        },
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ],
}

