module.exports = function (eleventyConfig) {
    eleventyConfig.addPassthroughCopy('src/bundle.css');

    return {
        dir: {
            input: 'src',
            output: 'dist'
        }
    };
};
