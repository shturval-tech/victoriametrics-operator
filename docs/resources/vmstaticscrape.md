# VMStaticScrape

The `VMStaticScrape` CRD provides mechanism for scraping metrics from static targets, configured by CRD targets.

`VMStaticScrape` object generates part of [VMAgent](https://docs.victoriametrics.com/operator/resources/vmagent.html) 
configuration with [static "service discovery"](https://docs.victoriametrics.com/sd_configs.html#static_configs).
It has various options for scraping configuration of target (with basic auth,tls access, by specific port name etc.).

By specifying configuration at CRD, operator generates config 
for [VMAgent](https://docs.victoriametrics.com/operator/resources/vmagent.html) and syncs it. 
It's useful for external targets management, when service-discovery is not available. 
`VMAgent` `staticScrapeSelector` must match `VMStaticScrape` labels.

More information about selectors you can find in [this doc](https://docs.victoriametrics.com/operator/resources/vmagent.html#scraping).

## Specification

You can see the full actual specification of the `VMStaticScrape` resource in
the [API docs -> VMStaticScrape](https://docs.victoriametrics.com/operator/api.html#vmstaticscrape).

<!-- TODO: examples -->