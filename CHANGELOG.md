# v3.2.1

- Changed `isNaN()` to safer `Integer.isNaN()`.

# v3.2.0

- Added `krypt strength` and `krypt make` global functions.
- Added more efficient and maintainable password leak checking.

# v3.1.4

- Added `krypt license` command to view the `LICENSE`.

# v3.1.3

- Added a Krypt terminal logo (See `krypt.logo`).

# v3.1.2

- Fixed Krypt "feature" which disabled use of capitals in names.

# v3.1.1

- Added missing info commands

# v3.1.0

- Extended `archive` to include directories.

# v3.0.0

- Added `archive` function to store encrypted files.

# v2.7.0

- Replaced smaller dictionary of `words.json` with **Webster Dictionary**.

# v2.6.4

- Database name shown on startup.

# v2.6.3

- Added better name format checking with `RegExp`.

# v2.6.2

- Added more strength groups.

# v2.6.1

- Added strength calculator function `strength <password>`.

# v2.6.0

- Changed strength calculator to popular `zxcvbn`.

# v2.5.2

- Removed unnecessary functions.

# v2.5.1

- Faster database reload for some functions.

# v2.5.0

- Added first version of `CHANGELOG.md` to the repository.

# v2.4.0

- Implemented required measures and updates to bring `PBKDF-2` to action. Key generation and checksum now with `PBKDF-2`.

# v2.3.0

- Replaced `SHA-256` with `PBKDF-2`, using node.js `crypto` implementation: `crypto.pbkdf2Sync()` (wrapped by `PBKDF2_HASH()`).

# v2.2.0

- Fixed cryptographically insecure `Math.random()` with secure `crypto.randomInt()` (wrapped by `crypt.random()`) for password generation.

# v2.1.0

- Faster database loading with removal of unnecessary reloads.

# v2.0.0

- `Wordy Passwords` introduced

# v1.3.0

- `2FA` introduced.

# v1.2.0

- `search` command to search for a password

# v1.1.0

- `list` command to list all passwords

# v1.0.0

- AES Encryption for password storage

- SHA Password checksum

- Password managing

- Password strength checker

- Password leak checking with [have i been pwned](https://haveibeenpwned.com)

- `unsafe` Strong password generator. **[ Uses cryptographically insecure `Math.random()` ]** **[ Fixed in 516f896 ]**

- Password health checking.

​

​