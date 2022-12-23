# Jobity

Jobity je veb aplikacija za pretragu poslova bazirana na mikroservisnoj arhitekturi.

Student - Milica Jovović, SW-15/2018

## Opis sistema

Glavni cilj Jobity-a je da pruži korisnicima priliku da lako pronađu posao koji odgovara njihovim sposobnostima i željama. Aplikacija omogućava pretragu, filtriranje i sortiranje postavljenih oglasa za posao, kao i prijavu na odabrane poslove. Takođe, korisnici mogu da komentarišu i ocenjuju poslodavce, ukoliko su radili za njih, kako bi pomogli drugim korisnicima pri izboru posla.

## Funkcionalnosti sistema

U sistemu postoje četiri kategorije korisnika: neautentifikovani korisnik, radnik, poslodavac i administrator.

### Neautentifikovani korisnik

Neautentifikovani korisnik ima ograničene mogućnosti unutar aplikacije. On može da pretražuje, filtrira i sortira oglase, ali ne može da se prijavljuje na njih. Takođe, može da pregleda ocene i komentare za svakog poslodavca, ali ne može da ostavlja svoje ocene i komentare. Da bi imao proširene mogućnosti, neautentifikovani korisnik mora da se registruje.

### Radnik

Radnik je autentifikovani korisnik koji se registrovao sa ciljem da pronađe posao. On ima sve mogućnosti kao i neautentifikovani korisnik, uz proširenje da može da se prijavi na odabrane oglase, kao i da ostavlja ocene i komentare za svoje poslodavce (trenutne ili bivše). Ukoliko poslodavac prihvati ili odbije prijavu za posao, radnik dobija e-mail obaveštenje o tome. Pri registraciji, radnik kreira svoj CV, odnosno unosi podatke o svom obrazovanju, poslovnom iskustvu i sposobnostima koje poseduje, ili prilaže postojeći CV u PDF formi. Na osnovu ovoga i izabranog poslovnog opredeljenja se prikazuju oglasi koji odgovaraju korisniku. Izabrano poslovno opredeljenje podrazumeva da korisnik odabere određene ključne reči vezane za kategorije poslova koji ga zanimaju (npr. IT, kuhinja, noćna smena, part-time, sezonski posao...). Radnik može naknadno menjati svoj CV i odabrano poslovno opredeljenje.

### Poslodavac

Poslodavac je autentifikovani korisnik koji se registrovao sa ciljem da pronađe radnika. On može da postavlja oglase za posao, pri čemu treba da unese opis posla, sposobnosti koje se traže, kao i ključne reči vezane za kategoriju tog posla. Takođe, poslodavac može da pregleda ocene i komentare koje radnici ostavljaju o njemu, pri čemu može da prijavi komentare za koje smatra da su neprimereni. Poslodavac dobija e-mail obaveštenja o prijavama na postavljenje oglase. Prijavu može da pregleda i da, na osnovu uvida, odluči da li prijava prihvaćena ili odbijena. Ukoliko je prijava prihvaćena, poslodavac zakazuje intervju za određeni datum koji može da se menja, a obaveštenje o odabranom terminu se šalje radniku putem e-mail-a. Kada datum intervjua prođe, poslodavac obeležava da li je radnik dobio posao ili ne. E-mail o rezultatu intervjua se šalje radniku.

### Administrator

Administrator ima uvid u čitav sistem, tj. može da pregleda sve korisnike, postavljene oglase i ocene i komentare. Ukoliko smatra da se neki od entiteta treba ukloniti (npr. spam oglasi), može da ukloni taj entitet, a može i da blokira korisnike koji postavljaju neprimereni sadržaj unutar sistema. Takođe, dobija prijave o neprimerenim komentarima od strane poslodavaca, na osnovu kojih može pregledati komentar i zaključiti da li komentar treba da se ukloni ili ne.

## Arhitektura sistema

### GoLang

* API Gateway
* Mikroservis za neautentifikovane korisnike
* Mikroservis za radnike
* Mikroservis za poslodavce
* Mikroservis za administratore
* Mikroservis za oglase
* Mikroservis za ocene i komentare

### Pharo

* Mikroservis za e-mail

### React

* Front-end veb aplikacija
