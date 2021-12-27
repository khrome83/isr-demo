const getResult = (body, options) => ({
  promise: Promise.resolve(),
  waitUntil: Promise.resolve(),
  response: new Response(body, options),
});

_ENTRIES = typeof _ENTRIES === "undefined" ? {} : _ENTRIES;

_ENTRIES["middleware_pages/_middleware"] = {
  default: async function ({ request }) {
    // Body
    if (request.url.endsWith("/middleware-body")) {
      return getResult("hi from the edge", {});
    }

    // Rewrite
    if (request.url.endsWith("/middleware-rewrite")) {
      return getResult(null, {
        headers: {
          "x-middleware-rewrite": "/blog",
        },
      });
    }

    // Redirect
    if (request.url.endsWith("/middleware-redirect")) {
      return getResult(null, {
        status: 308,
        headers: {
          Location: "/blog",
          "x-middleware-redirect": "/blog",
        },
      });
    }

    if (request.url.endsWith("/time")) {
      return getResult(null, {
        headers: {
          "x-middleware-rewrite": "/current",
        },
      });
    }

    // Don't do anything
    return getResult(null, {
      headers: {
        "x-middleware-next": "1",
      },
    });
  },
};
