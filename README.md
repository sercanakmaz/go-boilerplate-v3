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