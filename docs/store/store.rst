store.go
========

Purpose
-------

``store.go`` manages the main local JSON data file. It loads app data, creates starter data on first run, saves changes, and normalizes older or incomplete JSON into the current expected format.

This file is the local persistence layer for:

- student info
- courses
- assignments
- tasks

Constants
---------

appFolderName
~~~~~~~~~~~~~

Name of the folder created inside the user's config directory.

``dataFileName``
~~~~~~~~~~~~~~~~

Name of the main app data file: ``data.json``.

Function Documentation
----------------------

LoadAppData
~~~~~~~~~~~

.. code-block:: go

   func LoadAppData() (*model.AppData, error)

Loads the local app data file.

How it works:

1. Builds the data file path
2. Creates the file with starter data if it does not exist
3. Reads the JSON file
4. Decodes it into ``model.AppData``
5. Normalizes the data
6. Saves the repaired data back to disk if anything changed

This allows the app to recover from missing fields, old task statuses, and incomplete JSON.

SaveAppData
~~~~~~~~~~~

.. code-block:: go

   func SaveAppData(data model.AppData) error

Writes app data back to ``data.json``.

Before writing, it normalizes the data to keep the JSON format consistent. This is used after Canvas sync and will also be useful later for manual task editing.

dataFilePath
~~~~~~~~~~~~

.. code-block:: go

   func dataFilePath() (string, error)

Builds the cross-platform path to ``data.json`` using ``os.UserConfigDir``. This avoids hardcoding Linux or macOS-specific paths.

ensureDataFile
~~~~~~~~~~~~~~

.. code-block:: go

   func ensureDataFile(path string) error

Makes sure the app config directory and ``data.json`` file exist.

If the file is missing, it writes starter data from ``defaultAppData``.

defaultAppData
~~~~~~~~~~~~~~

.. code-block:: go

   func defaultAppData() model.AppData

Creates starter data so the UI has something to display on first launch.

The starter data includes:

- default student information
- two sample courses
- two sample assignments
- two sample tasks

The dates are generated relative to ``time.Now`` so they remain useful during testing.

normalizeAppData
~~~~~~~~~~~~~~~~

.. code-block:: go

   func normalizeAppData(data model.AppData) (model.AppData, bool)

Checks app data for missing or invalid fields and repairs them.

Repairs include:

- filling missing student name or school
- filling missing course IDs, names, and codes
- avoiding duplicate course IDs
- replacing missing assignments with defaults
- repairing invalid assignment course IDs
- filling missing assignment due dates
- avoiding duplicate assignment IDs
- filling missing task IDs and titles
- clearing invalid task course IDs
- filling missing task due dates
- creating default task start times
- normalizing task status values
- marking incomplete tasks as overdue after their due time passes

Returns the repaired data and a boolean showing whether anything changed.

normalizeTaskStatus
~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func normalizeTaskStatus(status string) string

Converts many possible task status strings into one of the app's supported statuses.

Examples:

- ``done`` and ``completed`` become ``complete``
- ``todo``, ``pending``, and ``not_started`` become ``in_progress``
- unknown values become ``in_progress``

This keeps user-edited JSON from breaking the UI.
