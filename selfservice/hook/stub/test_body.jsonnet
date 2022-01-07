function(ctx) {
  flow_id: ctx.flow.id,
  identity_id: if ctx["identity"] != null then ctx.identity.id,
  headers: ctx.request_headers,
  url: ctx.request_url,
  method: ctx.request_method,
  body: if std.objectHas(ctx, "request_body")  then ctx.request_body
}
