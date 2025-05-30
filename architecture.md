```bash
/cmd
  /api          # Gateway/API única
  /worker       # Workers para processamento assíncrono

/internal
  /shared       # Código compartilhado entre contextos
    /events     # Event bus/messaging
    /middleware
    /database
    /auth

  /company      # Bounded Context: Company
    /domain
      /entity
      /repository
      /service
    /usecase
    /adapter
      /http
      /database
      /events

  /user         # Bounded Context: User
    /domain
    /usecase
    /adapter

  /payments     # Bounded Context: Payments
    /domain
    /usecase
    /adapter
```
