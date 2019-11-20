const router = require('./router');
import 'babel-polyfill'

// export main function for server side rendering
global.main = router.renderToString;

// start app if it in the browser
if(typeof window !== 'undefined') {
    // Start main application here
    router.run();
}
