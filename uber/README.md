# Comparision of different multi error libraries

## Currently best: [`uber-go/muliterr`](https://github.com/uber-go/multierr)

- Doesn't introduce a own error type
- Append (known syntax)
- Convenience function for collection of deferred errors
- Compatible with `errors.Is` and `errors.As`
- Stable

## Notes on [hashicorp/go-multierror](https://github.com/hashicorp/go-multierror)

- Cutstom formatting
- Append (known syntax)
- Compatible with `errors.Is` and `errors.As`
- Stable
