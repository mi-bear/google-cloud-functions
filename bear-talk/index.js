const config = require('./config.json');

exports.bearTalk = (req, res) => {
  return Promise.resolve().then(() => {
    // method
    if (req.method !== 'POST') {
      const error = new Error('Only POST! bear is angry!!!');
      error.code = 405;
      throw error;
    }
    // verify
    verify(req.body);

    return talk(req.body);
  }).then((response) => {
    // response
    res.json(response);
  }).catch((err) => {
    // errors
    console.error(err);
    res.status(err.code || 500).send(err);
    return Promise.reject(err);
  });
};

function verify(body) {
  if (!body || body.token !== config.SLACK_TOKEN) {
    const error = new Error('Invalid credentials! bear is angry!!!');
    error.code = 401;
    throw error;
  }
}

function talk(body) {
  // TODO: add message valiations
  let response;
  switch (body.text) {
    case 'のどがかわいた':
      response = `<@${body.user_id}|${body.user_name}> :beer:`;
      break;
    default:
      response = message;
      break;
  }

  return {
    response_type: 'in_channel',
    text: response
  }
}
