Dashboard Widgets Module
========================

Source file: ``internal/widgets/widgets.go``

Overview
--------

This module builds and refreshes the main dashboard. The dashboard is made of five panels:

* Calendar summary
* Courses summary
* Task summary
* Assignments due today
* Upcoming assignments

The module also contains helper functions that turn ``model.AppData`` into short strings displayed inside those panels.

Functions
---------

BuildMainWidgets
~~~~~~~~~~~~~~~~

.. code-block:: go

   func BuildMainWidgets(data *model.AppData) ([]tview.Primitive, *tview.Flex)

Builds the full dashboard layout and returns both the selectable widgets and the root layout.

How it works:

The function creates five ``TextView`` widgets, enables dynamic colors, adds borders, and assigns each panel a numbered title. It then fills the widgets with text generated from helper functions.

The layout is built with nested ``tview.Flex`` containers. The header contains the quit/help text and the VIO title, while the body contains the five dashboard panels.

Returns:

``[]tview.Primitive``
    The five dashboard panels. This slice is used by navigation and routing.

``*tview.Flex``
    The full dashboard layout, suitable for passing to ``app.SetRoot``.

buildCalendarSummary
~~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func buildCalendarSummary(data *model.AppData) string

Creates a short list of assignment due dates for the current month.

The function scans all assignments, skips assignments without due dates, and keeps only assignments due in the current month and year. It then sorts the lines and displays up to eight entries.

If no assignments are due this month, it returns a message saying so.

buildCoursesSummary
~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func buildCoursesSummary(data *model.AppData) string

Creates the course summary shown on the dashboard.

Each course is displayed with its course code and name. If no courses are loaded, it returns a simple empty-state message.

buildTasksSummary
~~~~~~~~~~~~~~~~~

.. code-block:: go

   func buildTasksSummary(data *model.AppData) string

Creates a count summary for personal tasks.

The dashboard groups tasks into three states:

``in_progress``
    Tasks that are still active and not overdue.

``overdue``
    Tasks whose due time has passed and are not complete.

``complete``
    Tasks marked complete.

The actual category is calculated with ``dashboardTaskStatus``.

dashboardTaskStatus
~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func dashboardTaskStatus(task model.Task) string

Determines which dashboard category a task belongs to.

If the task status is ``complete``, it returns ``complete``. If the task has a due date before the current time, it returns ``overdue``. Otherwise, it returns ``in_progress``.

This allows the dashboard to show overdue tasks automatically even if the stored JSON has not yet been updated.

buildScheduleSummary
~~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func buildScheduleSummary(data *model.AppData) string

Builds the dashboard panel for assignments due today.

The function checks all assignments, converts due dates to local time, and keeps only those due on the same local day as today. It displays the assignment name and due time.

If nothing is due today, it returns a short empty-state message.

buildAssignmentsSummary
~~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func buildAssignmentsSummary(data *model.AppData) string

Builds the upcoming assignments dashboard panel.

The function filters out assignments without due dates and assignments that are already past due. It sorts the remaining assignments by due date and displays the top three upcoming items.

RefreshMainWidgets
~~~~~~~~~~~~~~~~~~

.. code-block:: go

   func RefreshMainWidgets(widgets []tview.Primitive, data *model.AppData)

Refreshes the text inside the five dashboard panels.

This is used after Canvas sync updates the app data. Instead of rebuilding the whole dashboard, it updates each existing ``TextView`` with fresh summary text.

The function checks that at least five widgets exist and type-asserts each one to ``*tview.TextView`` before calling ``SetText``.

sameLocalDay
~~~~~~~~~~~~

.. code-block:: go

   func sameLocalDay(a, b time.Time) bool

Compares two times using the local timezone and returns true if they fall on the same calendar day.

This is useful for due-date display because Canvas timestamps may include timezone information, and the UI should group assignments according to the user's local day.

Notes
-----

This module only creates dashboard summaries. Full-screen pages such as Calendar, Courses, Tasks, Schedule, and Assignments are handled in their own view packages.
