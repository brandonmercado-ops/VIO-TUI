models.go
=========

Purpose
-------

``models.go`` defines the main data structures used throughout VIO. These structs are the shared data format between the local JSON store, the Canvas API sync layer, and the UI views.

The model is intentionally small and compatible with the Instructure Canvas API. Canvas can provide courses and assignments, while the user can add local task data and extra course information.

Struct Documentation
--------------------

Student
~~~~~~~

.. code-block:: go

   type Student struct {
       Name   string `json:"name"`
       School string `json:"school"`
   }

Stores basic user profile information shown in the UI.

Fields:

``Name``
    The student's name.

``School``
    The student's school name.

Course
~~~~~~

.. code-block:: go

   type Course struct {
       ID           string   `json:"id"`
       Name         string   `json:"name"`
       Code         string   `json:"code"`
       Term         string   `json:"term,omitempty"`
       CanvasID     int      `json:"canvas_id,omitempty"`
       Workflow     string   `json:"workflow_state,omitempty"`
       Location     string   `json:"location,omitempty"`
       MeetingTimes []string `json:"meeting_times,omitempty"`
       Notes        string   `json:"notes,omitempty"`
   }

Represents an active course.

Canvas-backed fields:

``ID``
    Local string ID used by VIO. Canvas courses are stored as ``canvas-course-<id>``

``Name``
    Full course name

``Code``
    Short course code, such as ``COMP 123``

``Term``
    Academic term from Canvas, if available

``CanvasID``
    Original numeric Canvas course ID

``Workflow``
    Canvas workflow state, such as available

User-managed fields:

``Location``
    Optional class location entered by the user

``MeetingTimes``
    Optional list of meeting times entered by the user

``Notes``
    Optional course notes entered by the user

Assignment
~~~~~~~~~~

.. code-block:: go

   type Assignment struct {
       ID             string     `json:"id"`
       CourseID       string     `json:"course_id"`
       Name           string     `json:"name"`
       Description    string     `json:"description,omitempty"`
       DueAt          *time.Time `json:"due_at,omitempty"`
       UnlockAt       *time.Time `json:"unlock_at,omitempty"`
       LockAt         *time.Time `json:"lock_at,omitempty"`
       HasSubmitted   bool       `json:"has_submitted"`
       IsMissing      bool       `json:"is_missing"`
       IsLate         bool       `json:"is_late"`
       PointsPossible float64    `json:"points_possible,omitempty"`
       CanvasID       int        `json:"canvas_id,omitempty"`
   }

Represents an assignment from Canvas.

Important fields:

``CourseID``
    Links the assignment to a course in ``AppData.Courses``

``DueAt``
    Full date and time when the assignment is due

``UnlockAt`` and ``LockAt``
    Optional Canvas availability times

``HasSubmitted``
    True when Canvas reports a submission timestamp

``IsMissing`` and ``IsLate``
    Canvas submission status indicators

Task
~~~~

.. code-block:: go

   type Task struct {
       ID          string     `json:"id"`
       Title       string     `json:"title"`
       CourseID    string     `json:"course_id,omitempty"`
       Status      string     `json:"status"`
       StartsAt    *time.Time `json:"starts_at,omitempty"`
       DueAt       *time.Time `json:"due_at,omitempty"`
       Description string     `json:"description,omitempty"`
   }

Represents a user-created task or reminder.

Tasks are not fetched from Canvas. They are local user data.

Common statuses:

- ``in_progress``
- ``overdue``
- ``complete``

``DueAt`` stores both date and time. A task can become overdue automatically after its due time passes.

AppData
~~~~~~~

.. code-block:: go

   type AppData struct {
       Student     Student      `json:"student"`
       Courses     []Course     `json:"courses"`
       Assignments []Assignment `json:"assignments"`
       Tasks       []Task       `json:"tasks"`
   }

The main data container for the app.

Everything shown in the UI is pulled from this structure. It is also the shape written to ``data.json``.
