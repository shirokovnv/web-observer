# Web Observer

**Lightweight http resource pinger, written in [Go][link-go]**

- Periodically checks http response statuses for your webservices

- Sends slack notification when desired code differs from the http response

- Can be deployed and served as [systemd unit][link-systemd]

## Dependencies

- [Go][link-go] >= 1.19
- [Make][link-make]

## Project setup

1. git clone https://github.com/shirokovnv/web-observer && cd web-observer
2. Run `make init`

## Build and run locally

1. Ensure environment variables are set (see `.env.example`)
2. Fill site and slack channel params in `./config` directory
3. Run `make build` or `make run`

## License

MIT. Please see the [license file](license.md) for more information.

[link-go]: https://go.dev/
[link-make]: https://www.gnu.org/software/make/manual/make.html
[link-systemd]: https://systemd.io/
