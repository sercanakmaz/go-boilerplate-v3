İş Listesi
- Event Handlers
- Http Logging Middleware
- Echo Logging Middleware
- Command Validation
- RabbitMQ Consumer & Publisher
- RabbitMQ Logging Middleware
- App Config Management
- App Secret Management
- Outbox Pattern
- APM
- Sidecar Service Applications (Config, Localization, Feaute Toggle)
- Sidecar Service Implementation
- Dependecy Injection?
- Packaging & Dependecy validation at build time
- Comprehensive documentation & readme

pkg
anticorruption / http, api vb external data sources
events / her domain'in eventi
model / her domain'in modelleri
domains
    Açıklama:
        - ddd, hexagonal packaging style. n-tier değil.
        - hiç bir domain başka bir domain'e erişemez.
        - eğer oms
    - oms
        - domain / aggregates
        - core domain: siparişi oluşturduğun ve statülerini değiştiriğin ana servis
        - supporting domain: siparişin kargo durumunu beslediğin yardımcı servis 
        - generic domain: siparişin tüm durumları ile ilgili sms vb bilgilendirme yaptığın servis
    - product
    - delivery
    - merchant


business
    X modeli=entity
    A modeli
data
    X modeli=entity
    A modeli multiple entities
application
    ihtiyaç Z modeli yoksa X modeli=entity
    ihtiyaç varsa B modeli yoksa A modeli 


Event Handler senkron = hata durumunda revert yapılır
Event Consumer asenkron = hata durumunda revert yapılmaz, error queue'ya düşer