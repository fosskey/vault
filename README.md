<p align="center">
  <a href="https://fosskey.com">
    <img alt="Fosskey" src="https://user-images.githubusercontent.com/508043/210156279-fce2059d-4715-46a7-94e9-35f1d67ea431.png" width="143" />
  </a>
</p>

# Fosskey Vault

This repository contains a Go package to work with an encrypted vault. Check out [Fosskey CLI][cli-repo] if you're looking for a tool to store your secrets and passwords.

## What is Fosskey?

Fosskey is a [**F**]ree, [**O**]pen-source, [**S**]ecure, and [**S**]elf-custodial keychain.

## How does it work?

Fosskey does not store the master key. Instead, it uses the [Argon2id][argon2] key-derivation function to generate a 256-bit key from the master key and a 128-bit random salt. And then, the derived key and another 192-bit random nonce are used to encrypt the secret payload (plain text) using [XChaCha20-Poly1305][chacha20-poly1305] AEAD (Authenticated Encryption with Associated Data). Finally, the nonce, the cipher text, and the salt are glued together and stored in the `.foss/vault` file under your user's home directory (e.g. in `~/.foss/vault` on macOS and Linux, or `C:\Users\yourname\.foss\vault` on Windows).

<p align="left">
  <a href="https://fosskey.com">
    <img alt="Fosskey" src="https://user-images.githubusercontent.com/508043/210265188-671411b3-433c-4713-8734-1b1b8ee07d76.png" width="830" />
  </a>
</p>

## How secure is it?

While using the recommended parameters specified in [RFC 9106][rfc9106-params], the encryption/decryption method took about 0.8 seconds to process on a quad-core Intel processor with 16 GiB of memory. If a master key is composed of 8 characters of upper-case (A-Z), lower-case (a-z) letters and numbers (0-9), and symbols (32), there will be a total of 94 possible characters. Therefore, at least a total of B=nP(r-1) brute-force attacks is required to guess the correct master key. Here "B" is the permutation of (n, r-1). Thus, with the target hardware configuration (quad-core, 16 GiB memory), it will take about 1.3 million computation years to brute-force the 8-character long master key.

[cli-repo]: https://github.com/fosskey/cli
[chacha20-poly1305]: https://en.wikipedia.org/wiki/ChaCha20-Poly1305
[argon2]: https://en.wikipedia.org/wiki/Argon2
[rfc9106-params]: https://www.rfc-editor.org/rfc/rfc9106.html#name-parameter-choice
