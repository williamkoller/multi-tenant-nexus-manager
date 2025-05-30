# Domain Layer - Guia de Uso

Este documento explica como usar os componentes da camada de domínio compartilhada (`internal/shared/domain`) em sua arquitetura DDD com Go.

## 📁 Estrutura de Arquivos

```
internal/shared/domain/
├── entity.go              # BaseEntity com timestamps e soft delete
├── aggregate.go           # BaseAggregateRoot com domain events
├── events.go              # Sistema de domain events
├── repository.go          # Interfaces de repositório
└── value_objects/
    ├── email.go           # Validação de email
    ├── cpf.go             # CPF brasileiro
    ├── cnpj.go            # CNPJ brasileiro
    ├── money.go           # Valores monetários
    ├── address.go         # Endereço brasileiro
    ├── phone.go           # Telefone brasileiro
    ├── date_range.go      # Períodos de data
    ├── percentage.go      # Porcentagens
    ├── code.go            # Códigos alfanuméricos
    ├── slug.go            # URLs amigáveis
    └── color.go           # Cores hexadecimais
```

## 🏗️ Entidades Base

### BaseEntity

Use `BaseEntity` para entidades simples que não são agregados:

```go
package user

import "your-project/internal/shared/domain"

type UserProfile struct {
    domain.BaseEntity
    UserID string `json:"user_id" gorm:"not null"`
    Bio    string `json:"bio"`
    Avatar string `json:"avatar"`
}
```

**Funcionalidades incluídas:**

- ID UUID automático
- `created_at`, `updated_at` automáticos
- Soft delete com `deleted_at`
- Métodos: `IsDeleted()`

### BaseAggregateRoot

Use `BaseAggregateRoot` para agregados que geram domain events:

```go
package user

import (
    "your-project/internal/shared/domain"
    "your-project/internal/shared/domain/value_objects"
)

type User struct {
    domain.BaseAggregateRoot
    Email    value_objects.Email `json:"email"`
    CPF      value_objects.CPF   `json:"cpf"`
    Phone    value_objects.Phone `json:"phone"`
    IsActive bool                `json:"is_active"`
}

func (u *User) Activate() {
    u.IsActive = true

    // Dispara domain event
    event := domain.NewBaseDomainEvent(
        "user.activated",
        u.ID,
        map[string]interface{}{
            "email": u.Email.String(),
            "activated_at": time.Now(),
        },
    )
    u.RaiseDomainEvent(event)
}
```

## 💰 Value Objects

### Email

```go
import "your-project/internal/shared/domain/value_objects"

// Criação
email, err := value_objects.NewEmail("user@example.com")
if err != nil {
    return err
}

// Uso
fmt.Println(email.String())    // "user@example.com"
fmt.Println(email.IsEmpty())   // false
```

### CPF e CNPJ

```go
// CPF
cpf, err := value_objects.NewCPF("123.456.789-00")
if err != nil {
    return err
}

fmt.Println(cpf.String())      // "12345678900"
fmt.Println(cpf.Formatted())   // "123.456.789-00"

// CNPJ
cnpj, err := value_objects.NewCNPJ("12.345.678/0001-90")
if err != nil {
    return err
}

fmt.Println(cnpj.String())     // "12345678000190"
fmt.Println(cnpj.Formatted())  // "12.345.678/0001-90"
```

### Money (Valores Monetários)

```go
// Criação
price := value_objects.NewMoney(199.99, "BRL")
tax := value_objects.NewMoney(20.00, "BRL")

// Operações
total, err := price.Add(tax)
if err != nil {
    return err
}

// Comparações
if total.GreaterThan(price) {
    fmt.Println("Total é maior que preço")
}

// Formatação
fmt.Println(total.FormattedBRL()) // "R$ 219.99"
fmt.Println(total.String())       // "219.99 BRL"

// Multiplicação/Divisão
discount := total.Multiply(0.1)   // 10% de desconto
fmt.Println(discount.FormattedBRL()) // "R$ 21.99"
```

### Address (Endereço)

```go
address, err := value_objects.NewAddress(
    "Rua das Flores",     // street
    "123",                // number
    "Centro",             // district
    "São Paulo",          // city
    "SP",                 // state
    "01234-567",          // zipCode
    "Brasil",             // country
)

if err != nil {
    return err
}

fmt.Println(address.FullAddress())      // Endereço completo formatado
fmt.Println(address.FormattedZipCode()) // "01234-567"
fmt.Println(address.IsComplete())       // true
```

### Phone (Telefone)

```go
phone, err := value_objects.NewPhone("11987654321")
if err != nil {
    return err
}

fmt.Println(phone.String())     // "11987654321"
fmt.Println(phone.Formatted())  // "(11) 98765-4321"
fmt.Println(phone.IsMobile())   // true
```

### DateRange (Período)

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)

period, err := value_objects.NewDateRange(start, end)
if err != nil {
    return err
}

fmt.Println(period.DurationInDays())   // 365
fmt.Println(period.DurationInMonths()) // 11
fmt.Println(period.Contains(time.Now())) // true/false

// Verificar sobreposição
other, _ := value_objects.NewDateRange(
    time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
    time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC),
)
fmt.Println(period.Overlaps(other)) // true
```

### Percentage (Porcentagem)

```go
discount, err := value_objects.NewPercentage(15.5) // 15.5%
if err != nil {
    return err
}

price := value_objects.NewMoney(100.00, "BRL")
discountAmount := discount.ApplyTo(price)

fmt.Println(discount.String())           // "15.50%"
fmt.Println(discount.Decimal())          // 0.155
fmt.Println(discountAmount.FormattedBRL()) // "R$ 15.50"
```

## 📋 Repositories

### Definindo Repository para Agregado

```go
package user

import (
    "context"
    "your-project/internal/shared/domain"
)

type UserRepository interface {
    domain.Repository[*User]  // Herda Save, FindByID, Delete, Exists

    // Métodos específicos
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByCompany(ctx context.Context, companyID string) ([]*User, error)
}
```

### Implementando Repository

```go
package adapter

import (
    "context"
    "your-project/internal/user"
    "your-project/internal/shared/domain"
    "gorm.io/gorm"
)

type UserGormRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
    return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(ctx context.Context, u *user.User) error {
    return r.db.WithContext(ctx).Save(u).Error
}

func (r *UserGormRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
    var user user.User
    err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserGormRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
    var user user.User
    err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

## 📤 Domain Events

### Criando e Disparando Events

```go
package user

func (u *User) ChangeEmail(newEmail value_objects.Email) error {
    oldEmail := u.Email
    u.Email = newEmail

    // Dispara domain event
    event := domain.NewBaseDomainEvent(
        "user.email_changed",
        u.ID,
        map[string]interface{}{
            "old_email": oldEmail.String(),
            "new_email": newEmail.String(),
            "user_id":   u.ID,
        },
    )

    u.RaiseDomainEvent(event)
    return nil
}
```

### Processando Events no Use Case

```go
package usecase

func (uc *UserUseCase) ChangeUserEmail(ctx context.Context, userID string, newEmail string) error {
    // Busca usuário
    user, err := uc.userRepo.FindByID(ctx, userID)
    if err != nil {
        return err
    }

    // Cria value object
    email, err := value_objects.NewEmail(newEmail)
    if err != nil {
        return err
    }

    // Altera email (dispara domain event)
    err = user.ChangeEmail(email)
    if err != nil {
        return err
    }

    // Salva
    err = uc.userRepo.Save(ctx, user)
    if err != nil {
        return err
    }

    // Publica domain events
    for _, event := range user.GetDomainEvents() {
        uc.eventBus.Publish(ctx, event)
    }
    user.ClearDomainEvents()

    return nil
}
```

## 🎮 Exemplo Completo: Contexto User

```go
// internal/user/domain/user.go
package domain

import (
    "time"
    shared "your-project/internal/shared/domain"
    "your-project/internal/shared/domain/value_objects"
)

type User struct {
    shared.BaseAggregateRoot
    Email         value_objects.Email `json:"email" gorm:"unique;not null"`
    CPF           value_objects.CPF   `json:"cpf" gorm:"unique;not null"`
    Phone         value_objects.Phone `json:"phone"`
    FullName      string              `json:"full_name" gorm:"not null"`
    IsActive      bool                `json:"is_active" gorm:"default:true"`
    EmailVerified bool                `json:"email_verified" gorm:"default:false"`
    CompanyID     string              `json:"company_id" gorm:"not null"`
}

func NewUser(email, cpf, phone, fullName, companyID string) (*User, error) {
    emailVO, err := value_objects.NewEmail(email)
    if err != nil {
        return nil, err
    }

    cpfVO, err := value_objects.NewCPF(cpf)
    if err != nil {
        return nil, err
    }

    phoneVO, err := value_objects.NewPhone(phone)
    if err != nil {
        return nil, err
    }

    user := &User{
        Email:         emailVO,
        CPF:           cpfVO,
        Phone:         phoneVO,
        FullName:      fullName,
        IsActive:      true,
        EmailVerified: false,
        CompanyID:     companyID,
    }

    // Dispara evento de criação
    event := shared.NewBaseDomainEvent(
        "user.created",
        user.ID,
        map[string]interface{}{
            "email":      email,
            "company_id": companyID,
            "created_at": time.Now(),
        },
    )
    user.RaiseDomainEvent(event)

    return user, nil
}

func (u *User) Activate() {
    if !u.IsActive {
        u.IsActive = true

        event := shared.NewBaseDomainEvent(
            "user.activated",
            u.ID,
            map[string]interface{}{
                "activated_at": time.Now(),
            },
        )
        u.RaiseDomainEvent(event)
    }
}

func (u *User) VerifyEmail() {
    if !u.EmailVerified {
        u.EmailVerified = true

        event := shared.NewBaseDomainEvent(
            "user.email_verified",
            u.ID,
            map[string]interface{}{
                "verified_at": time.Now(),
            },
        )
        u.RaiseDomainEvent(event)
    }
}
```

## 🔄 Integração com Event Bus

```go
// No seu main.go ou dependency injection
eventBus := events.NewInMemoryEventBus()

// Registrar handlers para domain events
eventBus.Subscribe("user.created", func(ctx context.Context, event events.Event) error {
    // Enviar email de boas-vindas
    return emailService.SendWelcomeEmail(ctx, event.GetPayload())
})

eventBus.Subscribe("user.email_verified", func(ctx context.Context, event events.Event) error {
    // Atualizar permissões do usuário
    return permissionService.GrantBasicPermissions(ctx, event.GetAggregateID())
})
```

## ✅ Boas Práticas

1. **Value Objects**: Sempre valide dados no construtor
2. **Agregados**: Mantenha-os pequenos e coesos
3. **Domain Events**: Use para comunicação entre contextos
4. **Specifications**: Para regras de negócio complexas
5. **Repositories**: Uma interface por agregado
6. **Imutabilidade**: Value objects devem ser imutáveis

## 🚀 Próximos Passos

1. Implemente seus bounded contexts usando essas bases
2. Crie specifications específicas para suas regras de negócio
3. Configure o event bus para comunicação entre contextos
4. Adicione novos value objects conforme necessário
5. Considere migrar para microserviços quando os contextos crescerem

---

Este guia fornece uma base sólida para implementar DDD em Go com componentes reutilizáveis e type-safe!
