# druž.io
Druž.io je web aplikacija koja predstavlja društvenu mrežu.

# Funkcionalnosti:
Funkcionalnosti koje druž.io pruža:

### Ne ulogovani korisnik:
* Prijava
* Registracija (unos osnovnih informacija)
* Obnova lozinke

### Ulogovani korisnik:
* Rad sa svojim nalogom:
  * Pregled naloga
  * Izmena naloga (unos dodatnih podataka koji će se koristiti u pretrazi, izmena lozinke)
  * Deaktivacija naloga
* Rad sa objavama:
  * Pisanje objava
  * Brisanje objava
  * Komentarisanje objava
  * Reagovanje na objavu
* Rad sa drugim korisnicima:
  * Pretraga korisnika
  * Dodavanje prijatelja
  * Pregled svih prijatelja
  * Uklanjanje korisnika iz prijatelja
  * Blokiranje korisnika
  * Prijava korisnika
* Interakcija:
  * Praćenje online statusa korisnika
  * Komunikacija sa korisnicima  

### Administrator:
  * Pretraga korisnika
  * Pregled prijavljenih korisnika
  * Brisanje korisnika
  
# Arhitektura sistema:

Web aplikacija će biti bazirana na mikroservisnoj arhitekturi. Za kontejnerizaciju će biti korišten Docker.

- UserMicroservice - servis za rad sa korisnicima i autorizacijom. Tehnologije: Go, MariaDB.
- UserRelationsMicroservice - servis za sad sa odnosima između korisnika (prijateljstvo, blokiranje). Tehnologije: Go, MariaDB.
- PostMicroservice - servis za rad sa objavama, komentarima, reakcijama. Tehnologije: Rust, RavenDB.
- ChatMicroservice - servis za interakciju korisnika. Tehnologije: Go, RavenDB, WebSockets.
- EmailService - servis za slanje emailova. Sa ovim servisom će se komunicirati preko RabbitMQ. Tehnologije: Python, RabbitMQ.
- ImageService - servis za rad sa slikama. Tehnologije: Rust.
- Frontend - klijentska aplikacija. Tehnologije: React.

Dodatno, za replikaciju podataka između mikroservisa će biti korišćen RabbitMQ.
