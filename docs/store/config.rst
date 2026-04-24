config.go
=========

Purpose
-------

``config.go`` manages Canvas configuration. It stores the Canvas domain, encrypted API token, and polling interval in a separate config file, keeping Canvas credentials out of the main ``data.json`` file.

The token is encrypted locally before it is saved. The app can decrypt it later when Canvas sync needs to run.

Main Data Types
---------------

CanvasConfig
~~~~~~~~~~~~

.. code-block:: go

   type CanvasConfig struct {
       Domain      string
       Token       string
       PollMinutes int
   }

Runtime Canvas configuration after decryption.

Fields:

``Domain``
    The Canvas domain, such as ``school.instructure.com``

``Token``
    The decrypted Canvas API token

``PollMinutes``
    How often the app should poll Canvas. If not set, the app defaults to 5 minutes

storedConfig
~~~~~~~~~~~~

Internal JSON representation of the config file.

Unlike ``CanvasConfig``, this struct stores the token in encrypted form as ``TokenEnc``.

Function Documentation
----------------------

HasCredentials
~~~~~~~~~~~~~~

.. code-block:: go

   func (c CanvasConfig) HasCredentials() bool

Returns true when both the Canvas domain and token are present. The sync code uses this to decide whether Canvas polling should run.

PollInterval
~~~~~~~~~~~~

.. code-block:: go

   func (c CanvasConfig) PollInterval() time.Duration

Returns the polling interval as a ``time.Duration``. If the saved value is invalid or missing, it returns 5 minutes.

LoadCanvasConfig
~~~~~~~~~~~~~~~~

.. code-block:: go

   func LoadCanvasConfig() (CanvasConfig, bool, error)

Loads the Canvas config file and decrypts the saved token.

Returns:

``CanvasConfig``
    The usable runtime configuration

``bool``
    Whether a config file existed

``error``
    Any read, decode, or decrypt error

How it works:

1. Finds the config file path
2. If no file exists, returns an empty config and ``false``
3. Reads and decodes the JSON config
4. Decrypts the token
5. Applies a default poll interval if needed

SaveCanvasConfig
~~~~~~~~~~~~~~~~

.. code-block:: go

   func SaveCanvasConfig(cfg CanvasConfig) error

Encrypts and saves Canvas configuration.

How it works:

- Builds the config file path
- Creates the config directory if needed
- Encrypts the token
- Writes JSON with file permissions ``0600``

configFilePath
~~~~~~~~~~~~~~

.. code-block:: go

   func configFilePath() (string, error)

Builds the cross-platform path to the Canvas config file using ``os.UserConfigDir``. The file is stored under the app folder as ``config.json``.

machineSecret
~~~~~~~~~~~~~

.. code-block:: go

   func machineSecret() ([]byte, error)

Creates a local encryption key from the current username and hostname. The result is hashed with SHA-256 so it can be used as an AES key.

This keeps the token from being stored as plain text, but it is still intended as lightweight local protection rather than a full password manager.

encryptString
~~~~~~~~~~~~~

.. code-block:: go

   func encryptString(plain string) (string, error)

Encrypts a string using AES-GCM and returns it as base64 text.

How it works:

1. Gets the local machine secret
2. Creates an AES cipher
3. Wraps it in GCM mode
4. Generates a random nonce
5. Seals the plaintext and encodes it as base64

decryptString
~~~~~~~~~~~~~

.. code-block:: go

   func decryptString(enc string) (string, error)

Decrypts a base64 encoded AES-GCM string.

If the encrypted value is empty, it returns an empty string. Otherwise it decodes the base64 data, separates the nonce from the ciphertext, and opens the encrypted token.
