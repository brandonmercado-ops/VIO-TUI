client.go
=========

Purpose
-------

``client.go`` contains the Canvas API client.
It is responsible for creating authenticated Canvas requests, fetching active courses, fetching assignments for those courses, handling pagination, and converting Canvas API responses into VIO's local model structs.

Main Data Types
---------------

Client
~~~~~~

.. code-block:: go

   type Client struct {
       baseURL string
       token   string
       http    *http.Client
   }

Stores the Canvas base URL, the user's API token, and an HTTP client. The HTTP client uses a timeout so Canvas requests cannot hang forever.

apiCourse
~~~~~~~~~

Internal struct matching the Canvas course response fields used by VIO. It is not exposed to the rest of the app.

apiAssignment
~~~~~~~~~~~~~

Internal struct matching the Canvas assignment response fields used by VIO. It includes due dates, lock/unlock dates, point values, and submission information.

apiSubmission
~~~~~~~~~~~~~

Internal struct for assignment submission status. VIO uses it to determine whether an assignment has been submitted, is late, or is missing.

Function Documentation
----------------------

NewClient
~~~~~~~~~

.. code-block:: go

   func NewClient(cfg store.CanvasConfig) (*Client, error)

Creates a Canvas client from saved Canvas configuration.

How it works:

- Trims the Canvas domain and API token
- Rejects empty configuration
- Adds ``https://`` if the user only entered a domain
- Removes trailing slashes from the domain
- Validates that the domain is a parseable URL
- Creates an HTTP client with a 20 second timeout

SyncCoursesAndAssignments
~~~~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func SyncCoursesAndAssignments(cfg store.CanvasConfig, current model.AppData) (model.AppData, error)

Fetches current Canvas courses and assignments, then merges them into the existing app data.

How it works:

1. Builds a Canvas client from config
2. Fetches active courses
3. Loops through each course and fetches assignments for that course
4. Replaces the app's assignment list with the fresh Canvas assignment list
5. Merges courses while preserving user-entered course fields like location, meeting times, and notes

FetchActiveCourses
~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func (c *Client) FetchActiveCourses() ([]model.Course, error)

Fetches active Canvas courses for the current user.

How it works:

- Sends a request to ``/api/v1/courses``
- Requests active and available courses
- Includes term data
- Converts each Canvas course into a ``model.Course``
- Uses ``canvas-course-<id>`` as the local course ID format

FetchAssignments
~~~~~~~~~~~~~~~~

.. code-block:: go

   func (c *Client) FetchAssignments(courseCanvasID int) ([]model.Assignment, error)

Fetches assignments for a single Canvas course.

How it works:

- Calls ``/api/v1/courses/<course id>/assignments``
- Includes submission data
- Orders results by due date
- Skips assignments with empty names
- Converts Canvas timestamps to local time
- Maps Canvas assignment fields into ``model.Assignment``

getPaginatedJSON
~~~~~~~~~~~~~~~~

.. code-block:: go

   func getPaginatedJSON[T any](client *Client, path string, query url.Values) ([]T, error)

Generic helper for fetching paginated Canvas API results.

How it works:

- Builds the full URL from the base URL, path, and query values
- Adds the ``Authorization: Bearer`` token header
- Decodes each page of JSON into a slice of the requested type
- Uses the Canvas ``Link`` header to find the next page
- Appends all pages into one returned slice

canvasCourseKey
~~~~~~~~~~~~~~~

.. code-block:: go

   func canvasCourseKey(id int) string

Converts a numeric Canvas course ID into VIO's string ID format.

canvasAssignmentKey
~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func canvasAssignmentKey(id int) string

Converts a numeric Canvas assignment ID into VIO's string ID format.

mergeCourses
~~~~~~~~~~~~

.. code-block:: go

   func mergeCourses(existing []model.Course, fresh []model.Course) []model.Course

Combines newly fetched Canvas course data with existing local course data.

Canvas-owned fields are refreshed, but user-owned fields are preserved:

- ``Location``
- ``MeetingTimes``
- ``Notes``

nextLink
~~~~~~~~

.. code-block:: go

   func nextLink(linkHeader string) string

Extracts the next page URL from a Canvas pagination ``Link`` header. If no next page exists, it returns an empty string.

toLocalTimePtr
~~~~~~~~~~~~~~

.. code-block:: go

   func toLocalTimePtr(t *time.Time) *time.Time

Converts a Canvas timestamp pointer to local time. If the input is nil, it returns nil.
